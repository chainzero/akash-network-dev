// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import mock "github.com/stretchr/testify/mock"

// SubjectAccessReviewExpansion is an autogenerated mock type for the SubjectAccessReviewExpansion type
type SubjectAccessReviewExpansion struct {
	mock.Mock
}

type SubjectAccessReviewExpansion_Expecter struct {
	mock *mock.Mock
}

func (_m *SubjectAccessReviewExpansion) EXPECT() *SubjectAccessReviewExpansion_Expecter {
	return &SubjectAccessReviewExpansion_Expecter{mock: &_m.Mock}
}

type mockConstructorTestingTNewSubjectAccessReviewExpansion interface {
	mock.TestingT
	Cleanup(func())
}

// NewSubjectAccessReviewExpansion creates a new instance of SubjectAccessReviewExpansion. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSubjectAccessReviewExpansion(t mockConstructorTestingTNewSubjectAccessReviewExpansion) *SubjectAccessReviewExpansion {
	mock := &SubjectAccessReviewExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}