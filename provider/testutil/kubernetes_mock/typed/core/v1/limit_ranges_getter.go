// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// LimitRangesGetter is an autogenerated mock type for the LimitRangesGetter type
type LimitRangesGetter struct {
	mock.Mock
}

type LimitRangesGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *LimitRangesGetter) EXPECT() *LimitRangesGetter_Expecter {
	return &LimitRangesGetter_Expecter{mock: &_m.Mock}
}

// LimitRanges provides a mock function with given fields: namespace
func (_m *LimitRangesGetter) LimitRanges(namespace string) v1.LimitRangeInterface {
	ret := _m.Called(namespace)

	var r0 v1.LimitRangeInterface
	if rf, ok := ret.Get(0).(func(string) v1.LimitRangeInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.LimitRangeInterface)
		}
	}

	return r0
}

// LimitRangesGetter_LimitRanges_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LimitRanges'
type LimitRangesGetter_LimitRanges_Call struct {
	*mock.Call
}

// LimitRanges is a helper method to define mock.On call
//   - namespace string
func (_e *LimitRangesGetter_Expecter) LimitRanges(namespace interface{}) *LimitRangesGetter_LimitRanges_Call {
	return &LimitRangesGetter_LimitRanges_Call{Call: _e.mock.On("LimitRanges", namespace)}
}

func (_c *LimitRangesGetter_LimitRanges_Call) Run(run func(namespace string)) *LimitRangesGetter_LimitRanges_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *LimitRangesGetter_LimitRanges_Call) Return(_a0 v1.LimitRangeInterface) *LimitRangesGetter_LimitRanges_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LimitRangesGetter_LimitRanges_Call) RunAndReturn(run func(string) v1.LimitRangeInterface) *LimitRangesGetter_LimitRanges_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewLimitRangesGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewLimitRangesGetter creates a new instance of LimitRangesGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLimitRangesGetter(t mockConstructorTestingTNewLimitRangesGetter) *LimitRangesGetter {
	mock := &LimitRangesGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
