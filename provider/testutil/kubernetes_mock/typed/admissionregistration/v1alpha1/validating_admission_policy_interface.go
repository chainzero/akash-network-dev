// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	admissionregistrationv1alpha1 "k8s.io/api/admissionregistration/v1alpha1"

	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "k8s.io/client-go/applyconfigurations/admissionregistration/v1alpha1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// ValidatingAdmissionPolicyInterface is an autogenerated mock type for the ValidatingAdmissionPolicyInterface type
type ValidatingAdmissionPolicyInterface struct {
	mock.Mock
}

type ValidatingAdmissionPolicyInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ValidatingAdmissionPolicyInterface) EXPECT() *ValidatingAdmissionPolicyInterface_Expecter {
	return &ValidatingAdmissionPolicyInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, validatingAdmissionPolicy, opts
func (_m *ValidatingAdmissionPolicyInterface) Apply(ctx context.Context, validatingAdmissionPolicy *v1alpha1.ValidatingAdmissionPolicyApplyConfiguration, opts v1.ApplyOptions) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicy, error) {
	ret := _m.Called(ctx, validatingAdmissionPolicy, opts)

	var r0 *admissionregistrationv1alpha1.ValidatingAdmissionPolicy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.ValidatingAdmissionPolicyApplyConfiguration, v1.ApplyOptions) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicy, error)); ok {
		return rf(ctx, validatingAdmissionPolicy, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.ValidatingAdmissionPolicyApplyConfiguration, v1.ApplyOptions) *admissionregistrationv1alpha1.ValidatingAdmissionPolicy); ok {
		r0 = rf(ctx, validatingAdmissionPolicy, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1alpha1.ValidatingAdmissionPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1alpha1.ValidatingAdmissionPolicyApplyConfiguration, v1.ApplyOptions) error); ok {
		r1 = rf(ctx, validatingAdmissionPolicy, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatingAdmissionPolicyInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type ValidatingAdmissionPolicyInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - validatingAdmissionPolicy *v1alpha1.ValidatingAdmissionPolicyApplyConfiguration
//   - opts v1.ApplyOptions
func (_e *ValidatingAdmissionPolicyInterface_Expecter) Apply(ctx interface{}, validatingAdmissionPolicy interface{}, opts interface{}) *ValidatingAdmissionPolicyInterface_Apply_Call {
	return &ValidatingAdmissionPolicyInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, validatingAdmissionPolicy, opts)}
}

func (_c *ValidatingAdmissionPolicyInterface_Apply_Call) Run(run func(ctx context.Context, validatingAdmissionPolicy *v1alpha1.ValidatingAdmissionPolicyApplyConfiguration, opts v1.ApplyOptions)) *ValidatingAdmissionPolicyInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1alpha1.ValidatingAdmissionPolicyApplyConfiguration), args[2].(v1.ApplyOptions))
	})
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_Apply_Call) Return(result *admissionregistrationv1alpha1.ValidatingAdmissionPolicy, err error) *ValidatingAdmissionPolicyInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1alpha1.ValidatingAdmissionPolicyApplyConfiguration, v1.ApplyOptions) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicy, error)) *ValidatingAdmissionPolicyInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, validatingAdmissionPolicy, opts
func (_m *ValidatingAdmissionPolicyInterface) Create(ctx context.Context, validatingAdmissionPolicy *admissionregistrationv1alpha1.ValidatingAdmissionPolicy, opts v1.CreateOptions) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicy, error) {
	ret := _m.Called(ctx, validatingAdmissionPolicy, opts)

	var r0 *admissionregistrationv1alpha1.ValidatingAdmissionPolicy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *admissionregistrationv1alpha1.ValidatingAdmissionPolicy, v1.CreateOptions) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicy, error)); ok {
		return rf(ctx, validatingAdmissionPolicy, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *admissionregistrationv1alpha1.ValidatingAdmissionPolicy, v1.CreateOptions) *admissionregistrationv1alpha1.ValidatingAdmissionPolicy); ok {
		r0 = rf(ctx, validatingAdmissionPolicy, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1alpha1.ValidatingAdmissionPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *admissionregistrationv1alpha1.ValidatingAdmissionPolicy, v1.CreateOptions) error); ok {
		r1 = rf(ctx, validatingAdmissionPolicy, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatingAdmissionPolicyInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ValidatingAdmissionPolicyInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - validatingAdmissionPolicy *admissionregistrationv1alpha1.ValidatingAdmissionPolicy
//   - opts v1.CreateOptions
func (_e *ValidatingAdmissionPolicyInterface_Expecter) Create(ctx interface{}, validatingAdmissionPolicy interface{}, opts interface{}) *ValidatingAdmissionPolicyInterface_Create_Call {
	return &ValidatingAdmissionPolicyInterface_Create_Call{Call: _e.mock.On("Create", ctx, validatingAdmissionPolicy, opts)}
}

func (_c *ValidatingAdmissionPolicyInterface_Create_Call) Run(run func(ctx context.Context, validatingAdmissionPolicy *admissionregistrationv1alpha1.ValidatingAdmissionPolicy, opts v1.CreateOptions)) *ValidatingAdmissionPolicyInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admissionregistrationv1alpha1.ValidatingAdmissionPolicy), args[2].(v1.CreateOptions))
	})
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_Create_Call) Return(_a0 *admissionregistrationv1alpha1.ValidatingAdmissionPolicy, _a1 error) *ValidatingAdmissionPolicyInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_Create_Call) RunAndReturn(run func(context.Context, *admissionregistrationv1alpha1.ValidatingAdmissionPolicy, v1.CreateOptions) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicy, error)) *ValidatingAdmissionPolicyInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *ValidatingAdmissionPolicyInterface) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidatingAdmissionPolicyInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ValidatingAdmissionPolicyInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.DeleteOptions
func (_e *ValidatingAdmissionPolicyInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *ValidatingAdmissionPolicyInterface_Delete_Call {
	return &ValidatingAdmissionPolicyInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *ValidatingAdmissionPolicyInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts v1.DeleteOptions)) *ValidatingAdmissionPolicyInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.DeleteOptions))
	})
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_Delete_Call) Return(_a0 error) *ValidatingAdmissionPolicyInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_Delete_Call) RunAndReturn(run func(context.Context, string, v1.DeleteOptions) error) *ValidatingAdmissionPolicyInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *ValidatingAdmissionPolicyInterface) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.DeleteOptions, v1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidatingAdmissionPolicyInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type ValidatingAdmissionPolicyInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.DeleteOptions
//   - listOpts v1.ListOptions
func (_e *ValidatingAdmissionPolicyInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *ValidatingAdmissionPolicyInterface_DeleteCollection_Call {
	return &ValidatingAdmissionPolicyInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *ValidatingAdmissionPolicyInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions)) *ValidatingAdmissionPolicyInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.DeleteOptions), args[2].(v1.ListOptions))
	})
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_DeleteCollection_Call) Return(_a0 error) *ValidatingAdmissionPolicyInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, v1.DeleteOptions, v1.ListOptions) error) *ValidatingAdmissionPolicyInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *ValidatingAdmissionPolicyInterface) Get(ctx context.Context, name string, opts v1.GetOptions) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicy, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *admissionregistrationv1alpha1.ValidatingAdmissionPolicy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicy, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) *admissionregistrationv1alpha1.ValidatingAdmissionPolicy); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1alpha1.ValidatingAdmissionPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, v1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatingAdmissionPolicyInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ValidatingAdmissionPolicyInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.GetOptions
func (_e *ValidatingAdmissionPolicyInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *ValidatingAdmissionPolicyInterface_Get_Call {
	return &ValidatingAdmissionPolicyInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *ValidatingAdmissionPolicyInterface_Get_Call) Run(run func(ctx context.Context, name string, opts v1.GetOptions)) *ValidatingAdmissionPolicyInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.GetOptions))
	})
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_Get_Call) Return(_a0 *admissionregistrationv1alpha1.ValidatingAdmissionPolicy, _a1 error) *ValidatingAdmissionPolicyInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_Get_Call) RunAndReturn(run func(context.Context, string, v1.GetOptions) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicy, error)) *ValidatingAdmissionPolicyInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *ValidatingAdmissionPolicyInterface) List(ctx context.Context, opts v1.ListOptions) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicyList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *admissionregistrationv1alpha1.ValidatingAdmissionPolicyList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicyList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) *admissionregistrationv1alpha1.ValidatingAdmissionPolicyList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1alpha1.ValidatingAdmissionPolicyList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatingAdmissionPolicyInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ValidatingAdmissionPolicyInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *ValidatingAdmissionPolicyInterface_Expecter) List(ctx interface{}, opts interface{}) *ValidatingAdmissionPolicyInterface_List_Call {
	return &ValidatingAdmissionPolicyInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *ValidatingAdmissionPolicyInterface_List_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *ValidatingAdmissionPolicyInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_List_Call) Return(_a0 *admissionregistrationv1alpha1.ValidatingAdmissionPolicyList, _a1 error) *ValidatingAdmissionPolicyInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_List_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicyList, error)) *ValidatingAdmissionPolicyInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *ValidatingAdmissionPolicyInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicy, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *admissionregistrationv1alpha1.ValidatingAdmissionPolicy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicy, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) *admissionregistrationv1alpha1.ValidatingAdmissionPolicy); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1alpha1.ValidatingAdmissionPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatingAdmissionPolicyInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type ValidatingAdmissionPolicyInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts v1.PatchOptions
//   - subresources ...string
func (_e *ValidatingAdmissionPolicyInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *ValidatingAdmissionPolicyInterface_Patch_Call {
	return &ValidatingAdmissionPolicyInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *ValidatingAdmissionPolicyInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string)) *ValidatingAdmissionPolicyInterface_Patch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-5)
		for i, a := range args[5:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(types.PatchType), args[3].([]byte), args[4].(v1.PatchOptions), variadicArgs...)
	})
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_Patch_Call) Return(result *admissionregistrationv1alpha1.ValidatingAdmissionPolicy, err error) *ValidatingAdmissionPolicyInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicy, error)) *ValidatingAdmissionPolicyInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, validatingAdmissionPolicy, opts
func (_m *ValidatingAdmissionPolicyInterface) Update(ctx context.Context, validatingAdmissionPolicy *admissionregistrationv1alpha1.ValidatingAdmissionPolicy, opts v1.UpdateOptions) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicy, error) {
	ret := _m.Called(ctx, validatingAdmissionPolicy, opts)

	var r0 *admissionregistrationv1alpha1.ValidatingAdmissionPolicy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *admissionregistrationv1alpha1.ValidatingAdmissionPolicy, v1.UpdateOptions) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicy, error)); ok {
		return rf(ctx, validatingAdmissionPolicy, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *admissionregistrationv1alpha1.ValidatingAdmissionPolicy, v1.UpdateOptions) *admissionregistrationv1alpha1.ValidatingAdmissionPolicy); ok {
		r0 = rf(ctx, validatingAdmissionPolicy, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1alpha1.ValidatingAdmissionPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *admissionregistrationv1alpha1.ValidatingAdmissionPolicy, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, validatingAdmissionPolicy, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatingAdmissionPolicyInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ValidatingAdmissionPolicyInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - validatingAdmissionPolicy *admissionregistrationv1alpha1.ValidatingAdmissionPolicy
//   - opts v1.UpdateOptions
func (_e *ValidatingAdmissionPolicyInterface_Expecter) Update(ctx interface{}, validatingAdmissionPolicy interface{}, opts interface{}) *ValidatingAdmissionPolicyInterface_Update_Call {
	return &ValidatingAdmissionPolicyInterface_Update_Call{Call: _e.mock.On("Update", ctx, validatingAdmissionPolicy, opts)}
}

func (_c *ValidatingAdmissionPolicyInterface_Update_Call) Run(run func(ctx context.Context, validatingAdmissionPolicy *admissionregistrationv1alpha1.ValidatingAdmissionPolicy, opts v1.UpdateOptions)) *ValidatingAdmissionPolicyInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admissionregistrationv1alpha1.ValidatingAdmissionPolicy), args[2].(v1.UpdateOptions))
	})
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_Update_Call) Return(_a0 *admissionregistrationv1alpha1.ValidatingAdmissionPolicy, _a1 error) *ValidatingAdmissionPolicyInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_Update_Call) RunAndReturn(run func(context.Context, *admissionregistrationv1alpha1.ValidatingAdmissionPolicy, v1.UpdateOptions) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicy, error)) *ValidatingAdmissionPolicyInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *ValidatingAdmissionPolicyInterface) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) (watch.Interface, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatingAdmissionPolicyInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type ValidatingAdmissionPolicyInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *ValidatingAdmissionPolicyInterface_Expecter) Watch(ctx interface{}, opts interface{}) *ValidatingAdmissionPolicyInterface_Watch_Call {
	return &ValidatingAdmissionPolicyInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *ValidatingAdmissionPolicyInterface_Watch_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *ValidatingAdmissionPolicyInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *ValidatingAdmissionPolicyInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ValidatingAdmissionPolicyInterface_Watch_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (watch.Interface, error)) *ValidatingAdmissionPolicyInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewValidatingAdmissionPolicyInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewValidatingAdmissionPolicyInterface creates a new instance of ValidatingAdmissionPolicyInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewValidatingAdmissionPolicyInterface(t mockConstructorTestingTNewValidatingAdmissionPolicyInterface) *ValidatingAdmissionPolicyInterface {
	mock := &ValidatingAdmissionPolicyInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}