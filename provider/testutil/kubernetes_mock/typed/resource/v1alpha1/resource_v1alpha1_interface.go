// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	v1alpha1 "k8s.io/client-go/kubernetes/typed/resource/v1alpha1"
)

// ResourceV1alpha1Interface is an autogenerated mock type for the ResourceV1alpha1Interface type
type ResourceV1alpha1Interface struct {
	mock.Mock
}

type ResourceV1alpha1Interface_Expecter struct {
	mock *mock.Mock
}

func (_m *ResourceV1alpha1Interface) EXPECT() *ResourceV1alpha1Interface_Expecter {
	return &ResourceV1alpha1Interface_Expecter{mock: &_m.Mock}
}

// PodSchedulings provides a mock function with given fields: namespace
func (_m *ResourceV1alpha1Interface) PodSchedulings(namespace string) v1alpha1.PodSchedulingInterface {
	ret := _m.Called(namespace)

	var r0 v1alpha1.PodSchedulingInterface
	if rf, ok := ret.Get(0).(func(string) v1alpha1.PodSchedulingInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.PodSchedulingInterface)
		}
	}

	return r0
}

// ResourceV1alpha1Interface_PodSchedulings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PodSchedulings'
type ResourceV1alpha1Interface_PodSchedulings_Call struct {
	*mock.Call
}

// PodSchedulings is a helper method to define mock.On call
//   - namespace string
func (_e *ResourceV1alpha1Interface_Expecter) PodSchedulings(namespace interface{}) *ResourceV1alpha1Interface_PodSchedulings_Call {
	return &ResourceV1alpha1Interface_PodSchedulings_Call{Call: _e.mock.On("PodSchedulings", namespace)}
}

func (_c *ResourceV1alpha1Interface_PodSchedulings_Call) Run(run func(namespace string)) *ResourceV1alpha1Interface_PodSchedulings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ResourceV1alpha1Interface_PodSchedulings_Call) Return(_a0 v1alpha1.PodSchedulingInterface) *ResourceV1alpha1Interface_PodSchedulings_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResourceV1alpha1Interface_PodSchedulings_Call) RunAndReturn(run func(string) v1alpha1.PodSchedulingInterface) *ResourceV1alpha1Interface_PodSchedulings_Call {
	_c.Call.Return(run)
	return _c
}

// RESTClient provides a mock function with given fields:
func (_m *ResourceV1alpha1Interface) RESTClient() rest.Interface {
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

// ResourceV1alpha1Interface_RESTClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RESTClient'
type ResourceV1alpha1Interface_RESTClient_Call struct {
	*mock.Call
}

// RESTClient is a helper method to define mock.On call
func (_e *ResourceV1alpha1Interface_Expecter) RESTClient() *ResourceV1alpha1Interface_RESTClient_Call {
	return &ResourceV1alpha1Interface_RESTClient_Call{Call: _e.mock.On("RESTClient")}
}

func (_c *ResourceV1alpha1Interface_RESTClient_Call) Run(run func()) *ResourceV1alpha1Interface_RESTClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResourceV1alpha1Interface_RESTClient_Call) Return(_a0 rest.Interface) *ResourceV1alpha1Interface_RESTClient_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResourceV1alpha1Interface_RESTClient_Call) RunAndReturn(run func() rest.Interface) *ResourceV1alpha1Interface_RESTClient_Call {
	_c.Call.Return(run)
	return _c
}

// ResourceClaimTemplates provides a mock function with given fields: namespace
func (_m *ResourceV1alpha1Interface) ResourceClaimTemplates(namespace string) v1alpha1.ResourceClaimTemplateInterface {
	ret := _m.Called(namespace)

	var r0 v1alpha1.ResourceClaimTemplateInterface
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ResourceClaimTemplateInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ResourceClaimTemplateInterface)
		}
	}

	return r0
}

// ResourceV1alpha1Interface_ResourceClaimTemplates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResourceClaimTemplates'
type ResourceV1alpha1Interface_ResourceClaimTemplates_Call struct {
	*mock.Call
}

// ResourceClaimTemplates is a helper method to define mock.On call
//   - namespace string
func (_e *ResourceV1alpha1Interface_Expecter) ResourceClaimTemplates(namespace interface{}) *ResourceV1alpha1Interface_ResourceClaimTemplates_Call {
	return &ResourceV1alpha1Interface_ResourceClaimTemplates_Call{Call: _e.mock.On("ResourceClaimTemplates", namespace)}
}

func (_c *ResourceV1alpha1Interface_ResourceClaimTemplates_Call) Run(run func(namespace string)) *ResourceV1alpha1Interface_ResourceClaimTemplates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ResourceV1alpha1Interface_ResourceClaimTemplates_Call) Return(_a0 v1alpha1.ResourceClaimTemplateInterface) *ResourceV1alpha1Interface_ResourceClaimTemplates_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResourceV1alpha1Interface_ResourceClaimTemplates_Call) RunAndReturn(run func(string) v1alpha1.ResourceClaimTemplateInterface) *ResourceV1alpha1Interface_ResourceClaimTemplates_Call {
	_c.Call.Return(run)
	return _c
}

// ResourceClaims provides a mock function with given fields: namespace
func (_m *ResourceV1alpha1Interface) ResourceClaims(namespace string) v1alpha1.ResourceClaimInterface {
	ret := _m.Called(namespace)

	var r0 v1alpha1.ResourceClaimInterface
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ResourceClaimInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ResourceClaimInterface)
		}
	}

	return r0
}

// ResourceV1alpha1Interface_ResourceClaims_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResourceClaims'
type ResourceV1alpha1Interface_ResourceClaims_Call struct {
	*mock.Call
}

// ResourceClaims is a helper method to define mock.On call
//   - namespace string
func (_e *ResourceV1alpha1Interface_Expecter) ResourceClaims(namespace interface{}) *ResourceV1alpha1Interface_ResourceClaims_Call {
	return &ResourceV1alpha1Interface_ResourceClaims_Call{Call: _e.mock.On("ResourceClaims", namespace)}
}

func (_c *ResourceV1alpha1Interface_ResourceClaims_Call) Run(run func(namespace string)) *ResourceV1alpha1Interface_ResourceClaims_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ResourceV1alpha1Interface_ResourceClaims_Call) Return(_a0 v1alpha1.ResourceClaimInterface) *ResourceV1alpha1Interface_ResourceClaims_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResourceV1alpha1Interface_ResourceClaims_Call) RunAndReturn(run func(string) v1alpha1.ResourceClaimInterface) *ResourceV1alpha1Interface_ResourceClaims_Call {
	_c.Call.Return(run)
	return _c
}

// ResourceClasses provides a mock function with given fields:
func (_m *ResourceV1alpha1Interface) ResourceClasses() v1alpha1.ResourceClassInterface {
	ret := _m.Called()

	var r0 v1alpha1.ResourceClassInterface
	if rf, ok := ret.Get(0).(func() v1alpha1.ResourceClassInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ResourceClassInterface)
		}
	}

	return r0
}

// ResourceV1alpha1Interface_ResourceClasses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResourceClasses'
type ResourceV1alpha1Interface_ResourceClasses_Call struct {
	*mock.Call
}

// ResourceClasses is a helper method to define mock.On call
func (_e *ResourceV1alpha1Interface_Expecter) ResourceClasses() *ResourceV1alpha1Interface_ResourceClasses_Call {
	return &ResourceV1alpha1Interface_ResourceClasses_Call{Call: _e.mock.On("ResourceClasses")}
}

func (_c *ResourceV1alpha1Interface_ResourceClasses_Call) Run(run func()) *ResourceV1alpha1Interface_ResourceClasses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResourceV1alpha1Interface_ResourceClasses_Call) Return(_a0 v1alpha1.ResourceClassInterface) *ResourceV1alpha1Interface_ResourceClasses_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResourceV1alpha1Interface_ResourceClasses_Call) RunAndReturn(run func() v1alpha1.ResourceClassInterface) *ResourceV1alpha1Interface_ResourceClasses_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewResourceV1alpha1Interface interface {
	mock.TestingT
	Cleanup(func())
}

// NewResourceV1alpha1Interface creates a new instance of ResourceV1alpha1Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewResourceV1alpha1Interface(t mockConstructorTestingTNewResourceV1alpha1Interface) *ResourceV1alpha1Interface {
	mock := &ResourceV1alpha1Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
