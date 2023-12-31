// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	v1beta1 "k8s.io/client-go/kubernetes/typed/discovery/v1beta1"
)

// DiscoveryV1beta1Interface is an autogenerated mock type for the DiscoveryV1beta1Interface type
type DiscoveryV1beta1Interface struct {
	mock.Mock
}

type DiscoveryV1beta1Interface_Expecter struct {
	mock *mock.Mock
}

func (_m *DiscoveryV1beta1Interface) EXPECT() *DiscoveryV1beta1Interface_Expecter {
	return &DiscoveryV1beta1Interface_Expecter{mock: &_m.Mock}
}

// EndpointSlices provides a mock function with given fields: namespace
func (_m *DiscoveryV1beta1Interface) EndpointSlices(namespace string) v1beta1.EndpointSliceInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.EndpointSliceInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.EndpointSliceInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.EndpointSliceInterface)
		}
	}

	return r0
}

// DiscoveryV1beta1Interface_EndpointSlices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EndpointSlices'
type DiscoveryV1beta1Interface_EndpointSlices_Call struct {
	*mock.Call
}

// EndpointSlices is a helper method to define mock.On call
//   - namespace string
func (_e *DiscoveryV1beta1Interface_Expecter) EndpointSlices(namespace interface{}) *DiscoveryV1beta1Interface_EndpointSlices_Call {
	return &DiscoveryV1beta1Interface_EndpointSlices_Call{Call: _e.mock.On("EndpointSlices", namespace)}
}

func (_c *DiscoveryV1beta1Interface_EndpointSlices_Call) Run(run func(namespace string)) *DiscoveryV1beta1Interface_EndpointSlices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *DiscoveryV1beta1Interface_EndpointSlices_Call) Return(_a0 v1beta1.EndpointSliceInterface) *DiscoveryV1beta1Interface_EndpointSlices_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DiscoveryV1beta1Interface_EndpointSlices_Call) RunAndReturn(run func(string) v1beta1.EndpointSliceInterface) *DiscoveryV1beta1Interface_EndpointSlices_Call {
	_c.Call.Return(run)
	return _c
}

// RESTClient provides a mock function with given fields:
func (_m *DiscoveryV1beta1Interface) RESTClient() rest.Interface {
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

// DiscoveryV1beta1Interface_RESTClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RESTClient'
type DiscoveryV1beta1Interface_RESTClient_Call struct {
	*mock.Call
}

// RESTClient is a helper method to define mock.On call
func (_e *DiscoveryV1beta1Interface_Expecter) RESTClient() *DiscoveryV1beta1Interface_RESTClient_Call {
	return &DiscoveryV1beta1Interface_RESTClient_Call{Call: _e.mock.On("RESTClient")}
}

func (_c *DiscoveryV1beta1Interface_RESTClient_Call) Run(run func()) *DiscoveryV1beta1Interface_RESTClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DiscoveryV1beta1Interface_RESTClient_Call) Return(_a0 rest.Interface) *DiscoveryV1beta1Interface_RESTClient_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DiscoveryV1beta1Interface_RESTClient_Call) RunAndReturn(run func() rest.Interface) *DiscoveryV1beta1Interface_RESTClient_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewDiscoveryV1beta1Interface interface {
	mock.TestingT
	Cleanup(func())
}

// NewDiscoveryV1beta1Interface creates a new instance of DiscoveryV1beta1Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDiscoveryV1beta1Interface(t mockConstructorTestingTNewDiscoveryV1beta1Interface) *DiscoveryV1beta1Interface {
	mock := &DiscoveryV1beta1Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
