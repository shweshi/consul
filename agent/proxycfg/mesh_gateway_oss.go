// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !consulent
// +build !consulent

package proxycfg

import (
	"context"

	"github.com/hashicorp/go-hclog"
)

func (s *handlerMeshGateway) initializeEntWatches(_ context.Context) error {
	return nil
}

func (s *handlerMeshGateway) handleEntUpdate(_ hclog.Logger, _ context.Context, _ UpdateEvent, _ *ConfigSnapshot) error {
	return nil
}
