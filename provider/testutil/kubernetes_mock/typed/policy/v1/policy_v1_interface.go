// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	v1 "k8s.io/client-go/kubernetes/typed/policy/v1"
)

// PolicyV1Interface is an autogenerated mock type for the PolicyV1Interface type
type PolicyV1Interface struct {
	mock.Mock
}

type PolicyV1Interface_Expecter struct {
	mock *mock.Mock
}

func (_m *PolicyV1Interface) EXPECT() *PolicyV1Interface_Expecter {
	return &PolicyV1Interface_Expecter{mock: &_m.Mock}
}

// Evictions provides a mock function with given fields: namespace
func (_m *PolicyV1Interface) Evictions(namespace string) v1.EvictionInterface {
	ret := _m.Called(namespace)

	var r0 v1.EvictionInterface
	if rf, ok := ret.Get(0).(func(string) v1.EvictionInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.EvictionInterface)
		}
	}

	return r0
}

// PolicyV1Interface_Evictions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Evictions'
type PolicyV1Interface_Evictions_Call struct {
	*mock.Call
}

// Evictions is a helper method to define mock.On call
//   - namespace string
func (_e *PolicyV1Interface_Expecter) Evictions(namespace interface{}) *PolicyV1Interface_Evictions_Call {
	return &PolicyV1Interface_Evictions_Call{Call: _e.mock.On("Evictions", namespace)}
}

func (_c *PolicyV1Interface_Evictions_Call) Run(run func(namespace string)) *PolicyV1Interface_Evictions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *PolicyV1Interface_Evictions_Call) Return(_a0 v1.EvictionInterface) *PolicyV1Interface_Evictions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PolicyV1Interface_Evictions_Call) RunAndReturn(run func(string) v1.EvictionInterface) *PolicyV1Interface_Evictions_Call {
	_c.Call.Return(run)
	return _c
}

// PodDisruptionBudgets provides a mock function with given fields: namespace
func (_m *PolicyV1Interface) PodDisruptionBudgets(namespace string) v1.PodDisruptionBudgetInterface {
	ret := _m.Called(namespace)

	var r0 v1.PodDisruptionBudgetInterface
	if rf, ok := ret.Get(0).(func(string) v1.PodDisruptionBudgetInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.PodDisruptionBudgetInterface)
		}
	}

	return r0
}

// PolicyV1Interface_PodDisruptionBudgets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PodDisruptionBudgets'
type PolicyV1Interface_PodDisruptionBudgets_Call struct {
	*mock.Call
}

// PodDisruptionBudgets is a helper method to define mock.On call
//   - namespace string
func (_e *PolicyV1Interface_Expecter) PodDisruptionBudgets(namespace interface{}) *PolicyV1Interface_PodDisruptionBudgets_Call {
	return &PolicyV1Interface_PodDisruptionBudgets_Call{Call: _e.mock.On("PodDisruptionBudgets", namespace)}
}

func (_c *PolicyV1Interface_PodDisruptionBudgets_Call) Run(run func(namespace string)) *PolicyV1Interface_PodDisruptionBudgets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *PolicyV1Interface_PodDisruptionBudgets_Call) Return(_a0 v1.PodDisruptionBudgetInterface) *PolicyV1Interface_PodDisruptionBudgets_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PolicyV1Interface_PodDisruptionBudgets_Call) RunAndReturn(run func(string) v1.PodDisruptionBudgetInterface) *PolicyV1Interface_PodDisruptionBudgets_Call {
	_c.Call.Return(run)
	return _c
}

// RESTClient provides a mock function with given fields:
func (_m *PolicyV1Interface) RESTClient() rest.Interface {
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

// PolicyV1Interface_RESTClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RESTClient'
type PolicyV1Interface_RESTClient_Call struct {
	*mock.Call
}

// RESTClient is a helper method to define mock.On call
func (_e *PolicyV1Interface_Expecter) RESTClient() *PolicyV1Interface_RESTClient_Call {
	return &PolicyV1Interface_RESTClient_Call{Call: _e.mock.On("RESTClient")}
}

func (_c *PolicyV1Interface_RESTClient_Call) Run(run func()) *PolicyV1Interface_RESTClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PolicyV1Interface_RESTClient_Call) Return(_a0 rest.Interface) *PolicyV1Interface_RESTClient_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PolicyV1Interface_RESTClient_Call) RunAndReturn(run func() rest.Interface) *PolicyV1Interface_RESTClient_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewPolicyV1Interface interface {
	mock.TestingT
	Cleanup(func())
}

// NewPolicyV1Interface creates a new instance of PolicyV1Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPolicyV1Interface(t mockConstructorTestingTNewPolicyV1Interface) *PolicyV1Interface {
	mock := &PolicyV1Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
