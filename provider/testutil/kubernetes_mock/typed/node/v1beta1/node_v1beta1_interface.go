// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	v1beta1 "k8s.io/client-go/kubernetes/typed/node/v1beta1"
)

// NodeV1beta1Interface is an autogenerated mock type for the NodeV1beta1Interface type
type NodeV1beta1Interface struct {
	mock.Mock
}

type NodeV1beta1Interface_Expecter struct {
	mock *mock.Mock
}

func (_m *NodeV1beta1Interface) EXPECT() *NodeV1beta1Interface_Expecter {
	return &NodeV1beta1Interface_Expecter{mock: &_m.Mock}
}

// RESTClient provides a mock function with given fields:
func (_m *NodeV1beta1Interface) RESTClient() rest.Interface {
	ret := _m.Called()

	var r0 rest.Interface
	if rf, ok := ret.Get(0).(func() rest.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.Interface)
		}
	}

	return r0
}

// NodeV1beta1Interface_RESTClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RESTClient'
type NodeV1beta1Interface_RESTClient_Call struct {
	*mock.Call
}

// RESTClient is a helper method to define mock.On call
func (_e *NodeV1beta1Interface_Expecter) RESTClient() *NodeV1beta1Interface_RESTClient_Call {
	return &NodeV1beta1Interface_RESTClient_Call{Call: _e.mock.On("RESTClient")}
}

func (_c *NodeV1beta1Interface_RESTClient_Call) Run(run func()) *NodeV1beta1Interface_RESTClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *NodeV1beta1Interface_RESTClient_Call) Return(_a0 rest.Interface) *NodeV1beta1Interface_RESTClient_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NodeV1beta1Interface_RESTClient_Call) RunAndReturn(run func() rest.Interface) *NodeV1beta1Interface_RESTClient_Call {
	_c.Call.Return(run)
	return _c
}

// RuntimeClasses provides a mock function with given fields:
func (_m *NodeV1beta1Interface) RuntimeClasses() v1beta1.RuntimeClassInterface {
	ret := _m.Called()

	var r0 v1beta1.RuntimeClassInterface
	if rf, ok := ret.Get(0).(func() v1beta1.RuntimeClassInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.RuntimeClassInterface)
		}
	}

	return r0
}

// NodeV1beta1Interface_RuntimeClasses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RuntimeClasses'
type NodeV1beta1Interface_RuntimeClasses_Call struct {
	*mock.Call
}

// RuntimeClasses is a helper method to define mock.On call
func (_e *NodeV1beta1Interface_Expecter) RuntimeClasses() *NodeV1beta1Interface_RuntimeClasses_Call {
	return &NodeV1beta1Interface_RuntimeClasses_Call{Call: _e.mock.On("RuntimeClasses")}
}

func (_c *NodeV1beta1Interface_RuntimeClasses_Call) Run(run func()) *NodeV1beta1Interface_RuntimeClasses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *NodeV1beta1Interface_RuntimeClasses_Call) Return(_a0 v1beta1.RuntimeClassInterface) *NodeV1beta1Interface_RuntimeClasses_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NodeV1beta1Interface_RuntimeClasses_Call) RunAndReturn(run func() v1beta1.RuntimeClassInterface) *NodeV1beta1Interface_RuntimeClasses_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewNodeV1beta1Interface interface {
	mock.TestingT
	Cleanup(func())
}

// NewNodeV1beta1Interface creates a new instance of NodeV1beta1Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNodeV1beta1Interface(t mockConstructorTestingTNewNodeV1beta1Interface) *NodeV1beta1Interface {
	mock := &NodeV1beta1Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
