// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/client-go/kubernetes/typed/discovery/v1"
)

// EndpointSlicesGetter is an autogenerated mock type for the EndpointSlicesGetter type
type EndpointSlicesGetter struct {
	mock.Mock
}

type EndpointSlicesGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *EndpointSlicesGetter) EXPECT() *EndpointSlicesGetter_Expecter {
	return &EndpointSlicesGetter_Expecter{mock: &_m.Mock}
}

// EndpointSlices provides a mock function with given fields: namespace
func (_m *EndpointSlicesGetter) EndpointSlices(namespace string) v1.EndpointSliceInterface {
	ret := _m.Called(namespace)

	var r0 v1.EndpointSliceInterface
	if rf, ok := ret.Get(0).(func(string) v1.EndpointSliceInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.EndpointSliceInterface)
		}
	}

	return r0
}

// EndpointSlicesGetter_EndpointSlices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EndpointSlices'
type EndpointSlicesGetter_EndpointSlices_Call struct {
	*mock.Call
}

// EndpointSlices is a helper method to define mock.On call
//   - namespace string
func (_e *EndpointSlicesGetter_Expecter) EndpointSlices(namespace interface{}) *EndpointSlicesGetter_EndpointSlices_Call {
	return &EndpointSlicesGetter_EndpointSlices_Call{Call: _e.mock.On("EndpointSlices", namespace)}
}

func (_c *EndpointSlicesGetter_EndpointSlices_Call) Run(run func(namespace string)) *EndpointSlicesGetter_EndpointSlices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *EndpointSlicesGetter_EndpointSlices_Call) Return(_a0 v1.EndpointSliceInterface) *EndpointSlicesGetter_EndpointSlices_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EndpointSlicesGetter_EndpointSlices_Call) RunAndReturn(run func(string) v1.EndpointSliceInterface) *EndpointSlicesGetter_EndpointSlices_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewEndpointSlicesGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewEndpointSlicesGetter creates a new instance of EndpointSlicesGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEndpointSlicesGetter(t mockConstructorTestingTNewEndpointSlicesGetter) *EndpointSlicesGetter {
	mock := &EndpointSlicesGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
