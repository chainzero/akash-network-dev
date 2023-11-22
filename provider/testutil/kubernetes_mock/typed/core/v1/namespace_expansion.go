// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/api/core/v1"
)

// NamespaceExpansion is an autogenerated mock type for the NamespaceExpansion type
type NamespaceExpansion struct {
	mock.Mock
}

type NamespaceExpansion_Expecter struct {
	mock *mock.Mock
}

func (_m *NamespaceExpansion) EXPECT() *NamespaceExpansion_Expecter {
	return &NamespaceExpansion_Expecter{mock: &_m.Mock}
}

// Finalize provides a mock function with given fields: ctx, item, opts
func (_m *NamespaceExpansion) Finalize(ctx context.Context, item *v1.Namespace, opts metav1.UpdateOptions) (*v1.Namespace, error) {
	ret := _m.Called(ctx, item, opts)

	var r0 *v1.Namespace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Namespace, metav1.UpdateOptions) (*v1.Namespace, error)); ok {
		return rf(ctx, item, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Namespace, metav1.UpdateOptions) *v1.Namespace); ok {
		r0 = rf(ctx, item, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.Namespace, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, item, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceExpansion_Finalize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Finalize'
type NamespaceExpansion_Finalize_Call struct {
	*mock.Call
}

// Finalize is a helper method to define mock.On call
//   - ctx context.Context
//   - item *v1.Namespace
//   - opts metav1.UpdateOptions
func (_e *NamespaceExpansion_Expecter) Finalize(ctx interface{}, item interface{}, opts interface{}) *NamespaceExpansion_Finalize_Call {
	return &NamespaceExpansion_Finalize_Call{Call: _e.mock.On("Finalize", ctx, item, opts)}
}

func (_c *NamespaceExpansion_Finalize_Call) Run(run func(ctx context.Context, item *v1.Namespace, opts metav1.UpdateOptions)) *NamespaceExpansion_Finalize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.Namespace), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *NamespaceExpansion_Finalize_Call) Return(_a0 *v1.Namespace, _a1 error) *NamespaceExpansion_Finalize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NamespaceExpansion_Finalize_Call) RunAndReturn(run func(context.Context, *v1.Namespace, metav1.UpdateOptions) (*v1.Namespace, error)) *NamespaceExpansion_Finalize_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewNamespaceExpansion interface {
	mock.TestingT
	Cleanup(func())
}

// NewNamespaceExpansion creates a new instance of NamespaceExpansion. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNamespaceExpansion(t mockConstructorTestingTNewNamespaceExpansion) *NamespaceExpansion {
	mock := &NamespaceExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}