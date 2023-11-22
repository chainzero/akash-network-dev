// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"
)

// ServiceExpansion is an autogenerated mock type for the ServiceExpansion type
type ServiceExpansion struct {
	mock.Mock
}

type ServiceExpansion_Expecter struct {
	mock *mock.Mock
}

func (_m *ServiceExpansion) EXPECT() *ServiceExpansion_Expecter {
	return &ServiceExpansion_Expecter{mock: &_m.Mock}
}

// ProxyGet provides a mock function with given fields: scheme, name, port, path, params
func (_m *ServiceExpansion) ProxyGet(scheme string, name string, port string, path string, params map[string]string) rest.ResponseWrapper {
	ret := _m.Called(scheme, name, port, path, params)

	var r0 rest.ResponseWrapper
	if rf, ok := ret.Get(0).(func(string, string, string, string, map[string]string) rest.ResponseWrapper); ok {
		r0 = rf(scheme, name, port, path, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.ResponseWrapper)
		}
	}

	return r0
}

// ServiceExpansion_ProxyGet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProxyGet'
type ServiceExpansion_ProxyGet_Call struct {
	*mock.Call
}

// ProxyGet is a helper method to define mock.On call
//   - scheme string
//   - name string
//   - port string
//   - path string
//   - params map[string]string
func (_e *ServiceExpansion_Expecter) ProxyGet(scheme interface{}, name interface{}, port interface{}, path interface{}, params interface{}) *ServiceExpansion_ProxyGet_Call {
	return &ServiceExpansion_ProxyGet_Call{Call: _e.mock.On("ProxyGet", scheme, name, port, path, params)}
}

func (_c *ServiceExpansion_ProxyGet_Call) Run(run func(scheme string, name string, port string, path string, params map[string]string)) *ServiceExpansion_ProxyGet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(string), args[4].(map[string]string))
	})
	return _c
}

func (_c *ServiceExpansion_ProxyGet_Call) Return(_a0 rest.ResponseWrapper) *ServiceExpansion_ProxyGet_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServiceExpansion_ProxyGet_Call) RunAndReturn(run func(string, string, string, string, map[string]string) rest.ResponseWrapper) *ServiceExpansion_ProxyGet_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewServiceExpansion interface {
	mock.TestingT
	Cleanup(func())
}

// NewServiceExpansion creates a new instance of ServiceExpansion. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServiceExpansion(t mockConstructorTestingTNewServiceExpansion) *ServiceExpansion {
	mock := &ServiceExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}