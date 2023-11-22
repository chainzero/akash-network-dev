// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
)

// CronJobsGetter is an autogenerated mock type for the CronJobsGetter type
type CronJobsGetter struct {
	mock.Mock
}

type CronJobsGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *CronJobsGetter) EXPECT() *CronJobsGetter_Expecter {
	return &CronJobsGetter_Expecter{mock: &_m.Mock}
}

// CronJobs provides a mock function with given fields: namespace
func (_m *CronJobsGetter) CronJobs(namespace string) v1beta1.CronJobInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.CronJobInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.CronJobInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.CronJobInterface)
		}
	}

	return r0
}

// CronJobsGetter_CronJobs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CronJobs'
type CronJobsGetter_CronJobs_Call struct {
	*mock.Call
}

// CronJobs is a helper method to define mock.On call
//   - namespace string
func (_e *CronJobsGetter_Expecter) CronJobs(namespace interface{}) *CronJobsGetter_CronJobs_Call {
	return &CronJobsGetter_CronJobs_Call{Call: _e.mock.On("CronJobs", namespace)}
}

func (_c *CronJobsGetter_CronJobs_Call) Run(run func(namespace string)) *CronJobsGetter_CronJobs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *CronJobsGetter_CronJobs_Call) Return(_a0 v1beta1.CronJobInterface) *CronJobsGetter_CronJobs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CronJobsGetter_CronJobs_Call) RunAndReturn(run func(string) v1beta1.CronJobInterface) *CronJobsGetter_CronJobs_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCronJobsGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewCronJobsGetter creates a new instance of CronJobsGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCronJobsGetter(t mockConstructorTestingTNewCronJobsGetter) *CronJobsGetter {
	mock := &CronJobsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}