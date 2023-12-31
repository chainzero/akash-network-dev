// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/api/core/v1"
)

// NodeExpansion is an autogenerated mock type for the NodeExpansion type
type NodeExpansion struct {
	mock.Mock
}

type NodeExpansion_Expecter struct {
	mock *mock.Mock
}

func (_m *NodeExpansion) EXPECT() *NodeExpansion_Expecter {
	return &NodeExpansion_Expecter{mock: &_m.Mock}
}

// PatchStatus provides a mock function with given fields: ctx, nodeName, data
func (_m *NodeExpansion) PatchStatus(ctx context.Context, nodeName string, data []byte) (*v1.Node, error) {
	ret := _m.Called(ctx, nodeName, data)

	var r0 *v1.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) (*v1.Node, error)); ok {
		return rf(ctx, nodeName, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) *v1.Node); ok {
		r0 = rf(ctx, nodeName, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []byte) error); ok {
		r1 = rf(ctx, nodeName, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeExpansion_PatchStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchStatus'
type NodeExpansion_PatchStatus_Call struct {
	*mock.Call
}

// PatchStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeName string
//   - data []byte
func (_e *NodeExpansion_Expecter) PatchStatus(ctx interface{}, nodeName interface{}, data interface{}) *NodeExpansion_PatchStatus_Call {
	return &NodeExpansion_PatchStatus_Call{Call: _e.mock.On("PatchStatus", ctx, nodeName, data)}
}

func (_c *NodeExpansion_PatchStatus_Call) Run(run func(ctx context.Context, nodeName string, data []byte)) *NodeExpansion_PatchStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]byte))
	})
	return _c
}

func (_c *NodeExpansion_PatchStatus_Call) Return(_a0 *v1.Node, _a1 error) *NodeExpansion_PatchStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NodeExpansion_PatchStatus_Call) RunAndReturn(run func(context.Context, string, []byte) (*v1.Node, error)) *NodeExpansion_PatchStatus_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewNodeExpansion interface {
	mock.TestingT
	Cleanup(func())
}

// NewNodeExpansion creates a new instance of NodeExpansion. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNodeExpansion(t mockConstructorTestingTNewNodeExpansion) *NodeExpansion {
	mock := &NodeExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
