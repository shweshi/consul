// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package proxycfgglue

import (
	"context"

	"github.com/hashicorp/go-memdb"

	"github.com/hashicorp/consul/agent/cache"
	cachetype "github.com/hashicorp/consul/agent/cache-types"
	"github.com/hashicorp/consul/agent/configentry"
	"github.com/hashicorp/consul/agent/consul/watch"
	"github.com/hashicorp/consul/agent/proxycfg"
	"github.com/hashicorp/consul/agent/structs"
)

// CacheResolvedServiceConfig satisfies the proxycfg.ResolvedServiceConfig
// interface by sourcing data from the agent cache.
func CacheResolvedServiceConfig(c *cache.Cache) proxycfg.ResolvedServiceConfig {
	return &cacheProxyDataSource[*structs.ServiceConfigRequest]{c, cachetype.ResolvedServiceConfigName}
}

// ServerResolvedServiceConfig satisfies the proxycfg.ResolvedServiceConfig
// interface by sourcing data from a blocking query against the server's state
// store.
func ServerResolvedServiceConfig(deps ServerDataSourceDeps, remoteSource proxycfg.ResolvedServiceConfig) proxycfg.ResolvedServiceConfig {
	return &serverResolvedServiceConfig{deps, remoteSource}
}

type serverResolvedServiceConfig struct {
	deps         ServerDataSourceDeps
	remoteSource proxycfg.ResolvedServiceConfig
}

func (s *serverResolvedServiceConfig) Notify(ctx context.Context, req *structs.ServiceConfigRequest, correlationID string, ch chan<- proxycfg.UpdateEvent) error {
	if req.Datacenter != s.deps.Datacenter {
		return s.remoteSource.Notify(ctx, req, correlationID, ch)
	}

	return watch.ServerLocalNotify(ctx, correlationID, s.deps.GetStore,
		func(ws memdb.WatchSet, store Store) (uint64, *structs.ServiceConfigResponse, error) {
			authz, err := s.deps.ACLResolver.ResolveTokenAndDefaultMeta(req.Token, &req.EnterpriseMeta, nil)
			if err != nil {
				return 0, nil, err
			}

			if err := authz.ToAllowAuthorizer().ServiceReadAllowed(req.Name, nil); err != nil {
				return 0, nil, err
			}

			idx, entries, err := store.ReadResolvedServiceConfigEntries(ws, req.Name, &req.EnterpriseMeta, req.GetLocalUpstreamIDs(), req.Mode)
			if err != nil {
				return 0, nil, err
			}

			reply, err := configentry.ComputeResolvedServiceConfig(req, entries, s.deps.Logger)
			if err != nil {
				return 0, nil, err
			}
			reply.Index = idx

			return idx, reply, nil
		},
		dispatchBlockingQueryUpdate[*structs.ServiceConfigResponse](ch),
	)
}
