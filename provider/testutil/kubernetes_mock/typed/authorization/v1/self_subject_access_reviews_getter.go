// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
)

// SelfSubjectAccessReviewsGetter is an autogenerated mock type for the SelfSubjectAccessReviewsGetter type
type SelfSubjectAccessReviewsGetter struct {
	mock.Mock
}

type SelfSubjectAccessReviewsGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *SelfSubjectAccessReviewsGetter) EXPECT() *SelfSubjectAccessReviewsGetter_Expecter {
	return &SelfSubjectAccessReviewsGetter_Expecter{mock: &_m.Mock}
}

// SelfSubjectAccessReviews provides a mock function with given fields:
func (_m *SelfSubjectAccessReviewsGetter) SelfSubjectAccessReviews() v1.SelfSubjectAccessReviewInterface {
	ret := _m.Called()

	var r0 v1.SelfSubjectAccessReviewInterface
	if rf, ok := ret.Get(0).(func() v1.SelfSubjectAccessReviewInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.SelfSubjectAccessReviewInterface)
		}
	}

	return r0
}

// SelfSubjectAccessReviewsGetter_SelfSubjectAccessReviews_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelfSubjectAccessReviews'
type SelfSubjectAccessReviewsGetter_SelfSubjectAccessReviews_Call struct {
	*mock.Call
}

// SelfSubjectAccessReviews is a helper method to define mock.On call
func (_e *SelfSubjectAccessReviewsGetter_Expecter) SelfSubjectAccessReviews() *SelfSubjectAccessReviewsGetter_SelfSubjectAccessReviews_Call {
	return &SelfSubjectAccessReviewsGetter_SelfSubjectAccessReviews_Call{Call: _e.mock.On("SelfSubjectAccessReviews")}
}

func (_c *SelfSubjectAccessReviewsGetter_SelfSubjectAccessReviews_Call) Run(run func()) *SelfSubjectAccessReviewsGetter_SelfSubjectAccessReviews_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SelfSubjectAccessReviewsGetter_SelfSubjectAccessReviews_Call) Return(_a0 v1.SelfSubjectAccessReviewInterface) *SelfSubjectAccessReviewsGetter_SelfSubjectAccessReviews_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SelfSubjectAccessReviewsGetter_SelfSubjectAccessReviews_Call) RunAndReturn(run func() v1.SelfSubjectAccessReviewInterface) *SelfSubjectAccessReviewsGetter_SelfSubjectAccessReviews_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewSelfSubjectAccessReviewsGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewSelfSubjectAccessReviewsGetter creates a new instance of SelfSubjectAccessReviewsGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSelfSubjectAccessReviewsGetter(t mockConstructorTestingTNewSelfSubjectAccessReviewsGetter) *SelfSubjectAccessReviewsGetter {
	mock := &SelfSubjectAccessReviewsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
