// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	v1beta1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta1"
)

// FlowcontrolV1beta1Interface is an autogenerated mock type for the FlowcontrolV1beta1Interface type
type FlowcontrolV1beta1Interface struct {
	mock.Mock
}

type FlowcontrolV1beta1Interface_Expecter struct {
	mock *mock.Mock
}

func (_m *FlowcontrolV1beta1Interface) EXPECT() *FlowcontrolV1beta1Interface_Expecter {
	return &FlowcontrolV1beta1Interface_Expecter{mock: &_m.Mock}
}

// FlowSchemas provides a mock function with given fields:
func (_m *FlowcontrolV1beta1Interface) FlowSchemas() v1beta1.FlowSchemaInterface {
	ret := _m.Called()

	var r0 v1beta1.FlowSchemaInterface
	if rf, ok := ret.Get(0).(func() v1beta1.FlowSchemaInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.FlowSchemaInterface)
		}
	}

	return r0
}

// FlowcontrolV1beta1Interface_FlowSchemas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FlowSchemas'
type FlowcontrolV1beta1Interface_FlowSchemas_Call struct {
	*mock.Call
}

// FlowSchemas is a helper method to define mock.On call
func (_e *FlowcontrolV1beta1Interface_Expecter) FlowSchemas() *FlowcontrolV1beta1Interface_FlowSchemas_Call {
	return &FlowcontrolV1beta1Interface_FlowSchemas_Call{Call: _e.mock.On("FlowSchemas")}
}

func (_c *FlowcontrolV1beta1Interface_FlowSchemas_Call) Run(run func()) *FlowcontrolV1beta1Interface_FlowSchemas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FlowcontrolV1beta1Interface_FlowSchemas_Call) Return(_a0 v1beta1.FlowSchemaInterface) *FlowcontrolV1beta1Interface_FlowSchemas_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FlowcontrolV1beta1Interface_FlowSchemas_Call) RunAndReturn(run func() v1beta1.FlowSchemaInterface) *FlowcontrolV1beta1Interface_FlowSchemas_Call {
	_c.Call.Return(run)
	return _c
}

// PriorityLevelConfigurations provides a mock function with given fields:
func (_m *FlowcontrolV1beta1Interface) PriorityLevelConfigurations() v1beta1.PriorityLevelConfigurationInterface {
	ret := _m.Called()

	var r0 v1beta1.PriorityLevelConfigurationInterface
	if rf, ok := ret.Get(0).(func() v1beta1.PriorityLevelConfigurationInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.PriorityLevelConfigurationInterface)
		}
	}

	return r0
}

// FlowcontrolV1beta1Interface_PriorityLevelConfigurations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PriorityLevelConfigurations'
type FlowcontrolV1beta1Interface_PriorityLevelConfigurations_Call struct {
	*mock.Call
}

// PriorityLevelConfigurations is a helper method to define mock.On call
func (_e *FlowcontrolV1beta1Interface_Expecter) PriorityLevelConfigurations() *FlowcontrolV1beta1Interface_PriorityLevelConfigurations_Call {
	return &FlowcontrolV1beta1Interface_PriorityLevelConfigurations_Call{Call: _e.mock.On("PriorityLevelConfigurations")}
}

func (_c *FlowcontrolV1beta1Interface_PriorityLevelConfigurations_Call) Run(run func()) *FlowcontrolV1beta1Interface_PriorityLevelConfigurations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FlowcontrolV1beta1Interface_PriorityLevelConfigurations_Call) Return(_a0 v1beta1.PriorityLevelConfigurationInterface) *FlowcontrolV1beta1Interface_PriorityLevelConfigurations_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FlowcontrolV1beta1Interface_PriorityLevelConfigurations_Call) RunAndReturn(run func() v1beta1.PriorityLevelConfigurationInterface) *FlowcontrolV1beta1Interface_PriorityLevelConfigurations_Call {
	_c.Call.Return(run)
	return _c
}

// RESTClient provides a mock function with given fields:
func (_m *FlowcontrolV1beta1Interface) RESTClient() rest.Interface {
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

// FlowcontrolV1beta1Interface_RESTClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RESTClient'
type FlowcontrolV1beta1Interface_RESTClient_Call struct {
	*mock.Call
}

// RESTClient is a helper method to define mock.On call
func (_e *FlowcontrolV1beta1Interface_Expecter) RESTClient() *FlowcontrolV1beta1Interface_RESTClient_Call {
	return &FlowcontrolV1beta1Interface_RESTClient_Call{Call: _e.mock.On("RESTClient")}
}

func (_c *FlowcontrolV1beta1Interface_RESTClient_Call) Run(run func()) *FlowcontrolV1beta1Interface_RESTClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FlowcontrolV1beta1Interface_RESTClient_Call) Return(_a0 rest.Interface) *FlowcontrolV1beta1Interface_RESTClient_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FlowcontrolV1beta1Interface_RESTClient_Call) RunAndReturn(run func() rest.Interface) *FlowcontrolV1beta1Interface_RESTClient_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewFlowcontrolV1beta1Interface interface {
	mock.TestingT
	Cleanup(func())
}

// NewFlowcontrolV1beta1Interface creates a new instance of FlowcontrolV1beta1Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFlowcontrolV1beta1Interface(t mockConstructorTestingTNewFlowcontrolV1beta1Interface) *FlowcontrolV1beta1Interface {
	mock := &FlowcontrolV1beta1Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
