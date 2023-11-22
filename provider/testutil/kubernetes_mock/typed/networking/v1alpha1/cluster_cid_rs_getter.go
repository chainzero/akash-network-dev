// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1alpha1 "k8s.io/client-go/kubernetes/typed/networking/v1alpha1"
)

// ClusterCIDRsGetter is an autogenerated mock type for the ClusterCIDRsGetter type
type ClusterCIDRsGetter struct {
	mock.Mock
}

type ClusterCIDRsGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *ClusterCIDRsGetter) EXPECT() *ClusterCIDRsGetter_Expecter {
	return &ClusterCIDRsGetter_Expecter{mock: &_m.Mock}
}

// ClusterCIDRs provides a mock function with given fields:
func (_m *ClusterCIDRsGetter) ClusterCIDRs() v1alpha1.ClusterCIDRInterface {
	ret := _m.Called()

	var r0 v1alpha1.ClusterCIDRInterface
	if rf, ok := ret.Get(0).(func() v1alpha1.ClusterCIDRInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ClusterCIDRInterface)
		}
	}

	return r0
}

// ClusterCIDRsGetter_ClusterCIDRs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClusterCIDRs'
type ClusterCIDRsGetter_ClusterCIDRs_Call struct {
	*mock.Call
}

// ClusterCIDRs is a helper method to define mock.On call
func (_e *ClusterCIDRsGetter_Expecter) ClusterCIDRs() *ClusterCIDRsGetter_ClusterCIDRs_Call {
	return &ClusterCIDRsGetter_ClusterCIDRs_Call{Call: _e.mock.On("ClusterCIDRs")}
}

func (_c *ClusterCIDRsGetter_ClusterCIDRs_Call) Run(run func()) *ClusterCIDRsGetter_ClusterCIDRs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ClusterCIDRsGetter_ClusterCIDRs_Call) Return(_a0 v1alpha1.ClusterCIDRInterface) *ClusterCIDRsGetter_ClusterCIDRs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClusterCIDRsGetter_ClusterCIDRs_Call) RunAndReturn(run func() v1alpha1.ClusterCIDRInterface) *ClusterCIDRsGetter_ClusterCIDRs_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewClusterCIDRsGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewClusterCIDRsGetter creates a new instance of ClusterCIDRsGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClusterCIDRsGetter(t mockConstructorTestingTNewClusterCIDRsGetter) *ClusterCIDRsGetter {
	mock := &ClusterCIDRsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
