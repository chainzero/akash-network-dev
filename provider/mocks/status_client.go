// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	context "context"

	provider "github.com/akash-network/provider"
	mock "github.com/stretchr/testify/mock"
)

// StatusClient is an autogenerated mock type for the StatusClient type
type StatusClient struct {
	mock.Mock
}

type StatusClient_Expecter struct {
	mock *mock.Mock
}

func (_m *StatusClient) EXPECT() *StatusClient_Expecter {
	return &StatusClient_Expecter{mock: &_m.Mock}
}

// Status provides a mock function with given fields: _a0
func (_m *StatusClient) Status(_a0 context.Context) (*provider.Status, error) {
	ret := _m.Called(_a0)

	var r0 *provider.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*provider.Status, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *provider.Status); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provider.Status)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatusClient_Status_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Status'
type StatusClient_Status_Call struct {
	*mock.Call
}

// Status is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *StatusClient_Expecter) Status(_a0 interface{}) *StatusClient_Status_Call {
	return &StatusClient_Status_Call{Call: _e.mock.On("Status", _a0)}
}

func (_c *StatusClient_Status_Call) Run(run func(_a0 context.Context)) *StatusClient_Status_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *StatusClient_Status_Call) Return(_a0 *provider.Status, _a1 error) *StatusClient_Status_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StatusClient_Status_Call) RunAndReturn(run func(context.Context) (*provider.Status, error)) *StatusClient_Status_Call {
	_c.Call.Return(run)
	return _c
}

// NewStatusClient creates a new instance of StatusClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStatusClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *StatusClient {
	mock := &StatusClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}