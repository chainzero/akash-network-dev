// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	v1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
)

// AppsV1beta1Interface is an autogenerated mock type for the AppsV1beta1Interface type
type AppsV1beta1Interface struct {
	mock.Mock
}

type AppsV1beta1Interface_Expecter struct {
	mock *mock.Mock
}

func (_m *AppsV1beta1Interface) EXPECT() *AppsV1beta1Interface_Expecter {
	return &AppsV1beta1Interface_Expecter{mock: &_m.Mock}
}

// ControllerRevisions provides a mock function with given fields: namespace
func (_m *AppsV1beta1Interface) ControllerRevisions(namespace string) v1beta1.ControllerRevisionInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.ControllerRevisionInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.ControllerRevisionInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.ControllerRevisionInterface)
		}
	}

	return r0
}

// AppsV1beta1Interface_ControllerRevisions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ControllerRevisions'
type AppsV1beta1Interface_ControllerRevisions_Call struct {
	*mock.Call
}

// ControllerRevisions is a helper method to define mock.On call
//   - namespace string
func (_e *AppsV1beta1Interface_Expecter) ControllerRevisions(namespace interface{}) *AppsV1beta1Interface_ControllerRevisions_Call {
	return &AppsV1beta1Interface_ControllerRevisions_Call{Call: _e.mock.On("ControllerRevisions", namespace)}
}

func (_c *AppsV1beta1Interface_ControllerRevisions_Call) Run(run func(namespace string)) *AppsV1beta1Interface_ControllerRevisions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AppsV1beta1Interface_ControllerRevisions_Call) Return(_a0 v1beta1.ControllerRevisionInterface) *AppsV1beta1Interface_ControllerRevisions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppsV1beta1Interface_ControllerRevisions_Call) RunAndReturn(run func(string) v1beta1.ControllerRevisionInterface) *AppsV1beta1Interface_ControllerRevisions_Call {
	_c.Call.Return(run)
	return _c
}

// Deployments provides a mock function with given fields: namespace
func (_m *AppsV1beta1Interface) Deployments(namespace string) v1beta1.DeploymentInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.DeploymentInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.DeploymentInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.DeploymentInterface)
		}
	}

	return r0
}

// AppsV1beta1Interface_Deployments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Deployments'
type AppsV1beta1Interface_Deployments_Call struct {
	*mock.Call
}

// Deployments is a helper method to define mock.On call
//   - namespace string
func (_e *AppsV1beta1Interface_Expecter) Deployments(namespace interface{}) *AppsV1beta1Interface_Deployments_Call {
	return &AppsV1beta1Interface_Deployments_Call{Call: _e.mock.On("Deployments", namespace)}
}

func (_c *AppsV1beta1Interface_Deployments_Call) Run(run func(namespace string)) *AppsV1beta1Interface_Deployments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AppsV1beta1Interface_Deployments_Call) Return(_a0 v1beta1.DeploymentInterface) *AppsV1beta1Interface_Deployments_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppsV1beta1Interface_Deployments_Call) RunAndReturn(run func(string) v1beta1.DeploymentInterface) *AppsV1beta1Interface_Deployments_Call {
	_c.Call.Return(run)
	return _c
}

// RESTClient provides a mock function with given fields:
func (_m *AppsV1beta1Interface) RESTClient() rest.Interface {
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

// AppsV1beta1Interface_RESTClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RESTClient'
type AppsV1beta1Interface_RESTClient_Call struct {
	*mock.Call
}

// RESTClient is a helper method to define mock.On call
func (_e *AppsV1beta1Interface_Expecter) RESTClient() *AppsV1beta1Interface_RESTClient_Call {
	return &AppsV1beta1Interface_RESTClient_Call{Call: _e.mock.On("RESTClient")}
}

func (_c *AppsV1beta1Interface_RESTClient_Call) Run(run func()) *AppsV1beta1Interface_RESTClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AppsV1beta1Interface_RESTClient_Call) Return(_a0 rest.Interface) *AppsV1beta1Interface_RESTClient_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppsV1beta1Interface_RESTClient_Call) RunAndReturn(run func() rest.Interface) *AppsV1beta1Interface_RESTClient_Call {
	_c.Call.Return(run)
	return _c
}

// StatefulSets provides a mock function with given fields: namespace
func (_m *AppsV1beta1Interface) StatefulSets(namespace string) v1beta1.StatefulSetInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.StatefulSetInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.StatefulSetInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.StatefulSetInterface)
		}
	}

	return r0
}

// AppsV1beta1Interface_StatefulSets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StatefulSets'
type AppsV1beta1Interface_StatefulSets_Call struct {
	*mock.Call
}

// StatefulSets is a helper method to define mock.On call
//   - namespace string
func (_e *AppsV1beta1Interface_Expecter) StatefulSets(namespace interface{}) *AppsV1beta1Interface_StatefulSets_Call {
	return &AppsV1beta1Interface_StatefulSets_Call{Call: _e.mock.On("StatefulSets", namespace)}
}

func (_c *AppsV1beta1Interface_StatefulSets_Call) Run(run func(namespace string)) *AppsV1beta1Interface_StatefulSets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AppsV1beta1Interface_StatefulSets_Call) Return(_a0 v1beta1.StatefulSetInterface) *AppsV1beta1Interface_StatefulSets_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppsV1beta1Interface_StatefulSets_Call) RunAndReturn(run func(string) v1beta1.StatefulSetInterface) *AppsV1beta1Interface_StatefulSets_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewAppsV1beta1Interface interface {
	mock.TestingT
	Cleanup(func())
}

// NewAppsV1beta1Interface creates a new instance of AppsV1beta1Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAppsV1beta1Interface(t mockConstructorTestingTNewAppsV1beta1Interface) *AppsV1beta1Interface {
	mock := &AppsV1beta1Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
