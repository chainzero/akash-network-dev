// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
)

// ValidatingWebhookConfigurationsGetter is an autogenerated mock type for the ValidatingWebhookConfigurationsGetter type
type ValidatingWebhookConfigurationsGetter struct {
	mock.Mock
}

type ValidatingWebhookConfigurationsGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *ValidatingWebhookConfigurationsGetter) EXPECT() *ValidatingWebhookConfigurationsGetter_Expecter {
	return &ValidatingWebhookConfigurationsGetter_Expecter{mock: &_m.Mock}
}

// ValidatingWebhookConfigurations provides a mock function with given fields:
func (_m *ValidatingWebhookConfigurationsGetter) ValidatingWebhookConfigurations() v1beta1.ValidatingWebhookConfigurationInterface {
	ret := _m.Called()

	var r0 v1beta1.ValidatingWebhookConfigurationInterface
	if rf, ok := ret.Get(0).(func() v1beta1.ValidatingWebhookConfigurationInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.ValidatingWebhookConfigurationInterface)
		}
	}

	return r0
}

// ValidatingWebhookConfigurationsGetter_ValidatingWebhookConfigurations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidatingWebhookConfigurations'
type ValidatingWebhookConfigurationsGetter_ValidatingWebhookConfigurations_Call struct {
	*mock.Call
}

// ValidatingWebhookConfigurations is a helper method to define mock.On call
func (_e *ValidatingWebhookConfigurationsGetter_Expecter) ValidatingWebhookConfigurations() *ValidatingWebhookConfigurationsGetter_ValidatingWebhookConfigurations_Call {
	return &ValidatingWebhookConfigurationsGetter_ValidatingWebhookConfigurations_Call{Call: _e.mock.On("ValidatingWebhookConfigurations")}
}

func (_c *ValidatingWebhookConfigurationsGetter_ValidatingWebhookConfigurations_Call) Run(run func()) *ValidatingWebhookConfigurationsGetter_ValidatingWebhookConfigurations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ValidatingWebhookConfigurationsGetter_ValidatingWebhookConfigurations_Call) Return(_a0 v1beta1.ValidatingWebhookConfigurationInterface) *ValidatingWebhookConfigurationsGetter_ValidatingWebhookConfigurations_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ValidatingWebhookConfigurationsGetter_ValidatingWebhookConfigurations_Call) RunAndReturn(run func() v1beta1.ValidatingWebhookConfigurationInterface) *ValidatingWebhookConfigurationsGetter_ValidatingWebhookConfigurations_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewValidatingWebhookConfigurationsGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewValidatingWebhookConfigurationsGetter creates a new instance of ValidatingWebhookConfigurationsGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewValidatingWebhookConfigurationsGetter(t mockConstructorTestingTNewValidatingWebhookConfigurationsGetter) *ValidatingWebhookConfigurationsGetter {
	mock := &ValidatingWebhookConfigurationsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}