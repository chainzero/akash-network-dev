// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import mock "github.com/stretchr/testify/mock"

// ClusterCIDRExpansion is an autogenerated mock type for the ClusterCIDRExpansion type
type ClusterCIDRExpansion struct {
	mock.Mock
}

type ClusterCIDRExpansion_Expecter struct {
	mock *mock.Mock
}

func (_m *ClusterCIDRExpansion) EXPECT() *ClusterCIDRExpansion_Expecter {
	return &ClusterCIDRExpansion_Expecter{mock: &_m.Mock}
}

type mockConstructorTestingTNewClusterCIDRExpansion interface {
	mock.TestingT
	Cleanup(func())
}

// NewClusterCIDRExpansion creates a new instance of ClusterCIDRExpansion. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClusterCIDRExpansion(t mockConstructorTestingTNewClusterCIDRExpansion) *ClusterCIDRExpansion {
	mock := &ClusterCIDRExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
