// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
)

// TokenReviewsGetter is an autogenerated mock type for the TokenReviewsGetter type
type TokenReviewsGetter struct {
	mock.Mock
}

type TokenReviewsGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *TokenReviewsGetter) EXPECT() *TokenReviewsGetter_Expecter {
	return &TokenReviewsGetter_Expecter{mock: &_m.Mock}
}

// TokenReviews provides a mock function with given fields:
func (_m *TokenReviewsGetter) TokenReviews() v1beta1.TokenReviewInterface {
	ret := _m.Called()

	var r0 v1beta1.TokenReviewInterface
	if rf, ok := ret.Get(0).(func() v1beta1.TokenReviewInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.TokenReviewInterface)
		}
	}

	return r0
}

// TokenReviewsGetter_TokenReviews_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TokenReviews'
type TokenReviewsGetter_TokenReviews_Call struct {
	*mock.Call
}

// TokenReviews is a helper method to define mock.On call
func (_e *TokenReviewsGetter_Expecter) TokenReviews() *TokenReviewsGetter_TokenReviews_Call {
	return &TokenReviewsGetter_TokenReviews_Call{Call: _e.mock.On("TokenReviews")}
}

func (_c *TokenReviewsGetter_TokenReviews_Call) Run(run func()) *TokenReviewsGetter_TokenReviews_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TokenReviewsGetter_TokenReviews_Call) Return(_a0 v1beta1.TokenReviewInterface) *TokenReviewsGetter_TokenReviews_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TokenReviewsGetter_TokenReviews_Call) RunAndReturn(run func() v1beta1.TokenReviewInterface) *TokenReviewsGetter_TokenReviews_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewTokenReviewsGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewTokenReviewsGetter creates a new instance of TokenReviewsGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTokenReviewsGetter(t mockConstructorTestingTNewTokenReviewsGetter) *TokenReviewsGetter {
	mock := &TokenReviewsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}