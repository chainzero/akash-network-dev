// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import mock "github.com/stretchr/testify/mock"

// LeaseExpansion is an autogenerated mock type for the LeaseExpansion type
type LeaseExpansion struct {
	mock.Mock
}

type LeaseExpansion_Expecter struct {
	mock *mock.Mock
}

func (_m *LeaseExpansion) EXPECT() *LeaseExpansion_Expecter {
	return &LeaseExpansion_Expecter{mock: &_m.Mock}
}

type mockConstructorTestingTNewLeaseExpansion interface {
	mock.TestingT
	Cleanup(func())
}

// NewLeaseExpansion creates a new instance of LeaseExpansion. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLeaseExpansion(t mockConstructorTestingTNewLeaseExpansion) *LeaseExpansion {
	mock := &LeaseExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
