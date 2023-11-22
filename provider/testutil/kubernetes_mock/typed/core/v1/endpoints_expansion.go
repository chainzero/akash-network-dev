// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import mock "github.com/stretchr/testify/mock"

// EndpointsExpansion is an autogenerated mock type for the EndpointsExpansion type
type EndpointsExpansion struct {
	mock.Mock
}

type EndpointsExpansion_Expecter struct {
	mock *mock.Mock
}

func (_m *EndpointsExpansion) EXPECT() *EndpointsExpansion_Expecter {
	return &EndpointsExpansion_Expecter{mock: &_m.Mock}
}

type mockConstructorTestingTNewEndpointsExpansion interface {
	mock.TestingT
	Cleanup(func())
}

// NewEndpointsExpansion creates a new instance of EndpointsExpansion. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEndpointsExpansion(t mockConstructorTestingTNewEndpointsExpansion) *EndpointsExpansion {
	mock := &EndpointsExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
