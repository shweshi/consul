// Code generated by mockery v1.0.0. DO NOT EDIT.

package autopilotevents

import (
	acl "github.com/hashicorp/consul/acl"
	mock "github.com/stretchr/testify/mock"

	structs "github.com/hashicorp/consul/agent/structs"

	types "github.com/hashicorp/consul/types"
)

// MockStateStore is an autogenerated mock type for the StateStore type
type MockStateStore struct {
	mock.Mock
}

// GetNodeID provides a mock function with given fields: _a0, _a1
func (_m *MockStateStore) GetNodeID(_a0 types.NodeID, _a1 *acl.EnterpriseMeta) (uint64, *structs.Node, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(types.NodeID, *acl.EnterpriseMeta) uint64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 *structs.Node
	if rf, ok := ret.Get(1).(func(types.NodeID, *acl.EnterpriseMeta) *structs.Node); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*structs.Node)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(types.NodeID, *acl.EnterpriseMeta) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
