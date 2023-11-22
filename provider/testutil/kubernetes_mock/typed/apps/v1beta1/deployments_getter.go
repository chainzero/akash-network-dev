// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
)

// DeploymentsGetter is an autogenerated mock type for the DeploymentsGetter type
type DeploymentsGetter struct {
	mock.Mock
}

type DeploymentsGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *DeploymentsGetter) EXPECT() *DeploymentsGetter_Expecter {
	return &DeploymentsGetter_Expecter{mock: &_m.Mock}
}

// Deployments provides a mock function with given fields: namespace
func (_m *DeploymentsGetter) Deployments(namespace string) v1beta1.DeploymentInterface {
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

// DeploymentsGetter_Deployments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Deployments'
type DeploymentsGetter_Deployments_Call struct {
	*mock.Call
}

// Deployments is a helper method to define mock.On call
//   - namespace string
func (_e *DeploymentsGetter_Expecter) Deployments(namespace interface{}) *DeploymentsGetter_Deployments_Call {
	return &DeploymentsGetter_Deployments_Call{Call: _e.mock.On("Deployments", namespace)}
}

func (_c *DeploymentsGetter_Deployments_Call) Run(run func(namespace string)) *DeploymentsGetter_Deployments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *DeploymentsGetter_Deployments_Call) Return(_a0 v1beta1.DeploymentInterface) *DeploymentsGetter_Deployments_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeploymentsGetter_Deployments_Call) RunAndReturn(run func(string) v1beta1.DeploymentInterface) *DeploymentsGetter_Deployments_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewDeploymentsGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewDeploymentsGetter creates a new instance of DeploymentsGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDeploymentsGetter(t mockConstructorTestingTNewDeploymentsGetter) *DeploymentsGetter {
	mock := &DeploymentsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
