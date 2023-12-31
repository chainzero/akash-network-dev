// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
)

// ClusterRoleBindingsGetter is an autogenerated mock type for the ClusterRoleBindingsGetter type
type ClusterRoleBindingsGetter struct {
	mock.Mock
}

type ClusterRoleBindingsGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *ClusterRoleBindingsGetter) EXPECT() *ClusterRoleBindingsGetter_Expecter {
	return &ClusterRoleBindingsGetter_Expecter{mock: &_m.Mock}
}

// ClusterRoleBindings provides a mock function with given fields:
func (_m *ClusterRoleBindingsGetter) ClusterRoleBindings() v1beta1.ClusterRoleBindingInterface {
	ret := _m.Called()

	var r0 v1beta1.ClusterRoleBindingInterface
	if rf, ok := ret.Get(0).(func() v1beta1.ClusterRoleBindingInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.ClusterRoleBindingInterface)
		}
	}

	return r0
}

// ClusterRoleBindingsGetter_ClusterRoleBindings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClusterRoleBindings'
type ClusterRoleBindingsGetter_ClusterRoleBindings_Call struct {
	*mock.Call
}

// ClusterRoleBindings is a helper method to define mock.On call
func (_e *ClusterRoleBindingsGetter_Expecter) ClusterRoleBindings() *ClusterRoleBindingsGetter_ClusterRoleBindings_Call {
	return &ClusterRoleBindingsGetter_ClusterRoleBindings_Call{Call: _e.mock.On("ClusterRoleBindings")}
}

func (_c *ClusterRoleBindingsGetter_ClusterRoleBindings_Call) Run(run func()) *ClusterRoleBindingsGetter_ClusterRoleBindings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ClusterRoleBindingsGetter_ClusterRoleBindings_Call) Return(_a0 v1beta1.ClusterRoleBindingInterface) *ClusterRoleBindingsGetter_ClusterRoleBindings_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClusterRoleBindingsGetter_ClusterRoleBindings_Call) RunAndReturn(run func() v1beta1.ClusterRoleBindingInterface) *ClusterRoleBindingsGetter_ClusterRoleBindings_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewClusterRoleBindingsGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewClusterRoleBindingsGetter creates a new instance of ClusterRoleBindingsGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClusterRoleBindingsGetter(t mockConstructorTestingTNewClusterRoleBindingsGetter) *ClusterRoleBindingsGetter {
	mock := &ClusterRoleBindingsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
