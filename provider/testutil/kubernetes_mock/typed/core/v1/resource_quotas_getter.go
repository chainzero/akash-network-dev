// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// ResourceQuotasGetter is an autogenerated mock type for the ResourceQuotasGetter type
type ResourceQuotasGetter struct {
	mock.Mock
}

type ResourceQuotasGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *ResourceQuotasGetter) EXPECT() *ResourceQuotasGetter_Expecter {
	return &ResourceQuotasGetter_Expecter{mock: &_m.Mock}
}

// ResourceQuotas provides a mock function with given fields: namespace
func (_m *ResourceQuotasGetter) ResourceQuotas(namespace string) v1.ResourceQuotaInterface {
	ret := _m.Called(namespace)

	var r0 v1.ResourceQuotaInterface
	if rf, ok := ret.Get(0).(func(string) v1.ResourceQuotaInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.ResourceQuotaInterface)
		}
	}

	return r0
}

// ResourceQuotasGetter_ResourceQuotas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResourceQuotas'
type ResourceQuotasGetter_ResourceQuotas_Call struct {
	*mock.Call
}

// ResourceQuotas is a helper method to define mock.On call
//   - namespace string
func (_e *ResourceQuotasGetter_Expecter) ResourceQuotas(namespace interface{}) *ResourceQuotasGetter_ResourceQuotas_Call {
	return &ResourceQuotasGetter_ResourceQuotas_Call{Call: _e.mock.On("ResourceQuotas", namespace)}
}

func (_c *ResourceQuotasGetter_ResourceQuotas_Call) Run(run func(namespace string)) *ResourceQuotasGetter_ResourceQuotas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ResourceQuotasGetter_ResourceQuotas_Call) Return(_a0 v1.ResourceQuotaInterface) *ResourceQuotasGetter_ResourceQuotas_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResourceQuotasGetter_ResourceQuotas_Call) RunAndReturn(run func(string) v1.ResourceQuotaInterface) *ResourceQuotasGetter_ResourceQuotas_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewResourceQuotasGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewResourceQuotasGetter creates a new instance of ResourceQuotasGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewResourceQuotasGetter(t mockConstructorTestingTNewResourceQuotasGetter) *ResourceQuotasGetter {
	mock := &ResourceQuotasGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}