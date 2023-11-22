// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
)

// MutatingWebhookConfigurationsGetter is an autogenerated mock type for the MutatingWebhookConfigurationsGetter type
type MutatingWebhookConfigurationsGetter struct {
	mock.Mock
}

type MutatingWebhookConfigurationsGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *MutatingWebhookConfigurationsGetter) EXPECT() *MutatingWebhookConfigurationsGetter_Expecter {
	return &MutatingWebhookConfigurationsGetter_Expecter{mock: &_m.Mock}
}

// MutatingWebhookConfigurations provides a mock function with given fields:
func (_m *MutatingWebhookConfigurationsGetter) MutatingWebhookConfigurations() v1.MutatingWebhookConfigurationInterface {
	ret := _m.Called()

	var r0 v1.MutatingWebhookConfigurationInterface
	if rf, ok := ret.Get(0).(func() v1.MutatingWebhookConfigurationInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.MutatingWebhookConfigurationInterface)
		}
	}

	return r0
}

// MutatingWebhookConfigurationsGetter_MutatingWebhookConfigurations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MutatingWebhookConfigurations'
type MutatingWebhookConfigurationsGetter_MutatingWebhookConfigurations_Call struct {
	*mock.Call
}

// MutatingWebhookConfigurations is a helper method to define mock.On call
func (_e *MutatingWebhookConfigurationsGetter_Expecter) MutatingWebhookConfigurations() *MutatingWebhookConfigurationsGetter_MutatingWebhookConfigurations_Call {
	return &MutatingWebhookConfigurationsGetter_MutatingWebhookConfigurations_Call{Call: _e.mock.On("MutatingWebhookConfigurations")}
}

func (_c *MutatingWebhookConfigurationsGetter_MutatingWebhookConfigurations_Call) Run(run func()) *MutatingWebhookConfigurationsGetter_MutatingWebhookConfigurations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MutatingWebhookConfigurationsGetter_MutatingWebhookConfigurations_Call) Return(_a0 v1.MutatingWebhookConfigurationInterface) *MutatingWebhookConfigurationsGetter_MutatingWebhookConfigurations_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutatingWebhookConfigurationsGetter_MutatingWebhookConfigurations_Call) RunAndReturn(run func() v1.MutatingWebhookConfigurationInterface) *MutatingWebhookConfigurationsGetter_MutatingWebhookConfigurations_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMutatingWebhookConfigurationsGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewMutatingWebhookConfigurationsGetter creates a new instance of MutatingWebhookConfigurationsGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMutatingWebhookConfigurationsGetter(t mockConstructorTestingTNewMutatingWebhookConfigurationsGetter) *MutatingWebhookConfigurationsGetter {
	mock := &MutatingWebhookConfigurationsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}