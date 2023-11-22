// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	rbacv1alpha1 "k8s.io/api/rbac/v1alpha1"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "k8s.io/client-go/applyconfigurations/rbac/v1alpha1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// RoleInterface is an autogenerated mock type for the RoleInterface type
type RoleInterface struct {
	mock.Mock
}

type RoleInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *RoleInterface) EXPECT() *RoleInterface_Expecter {
	return &RoleInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, role, opts
func (_m *RoleInterface) Apply(ctx context.Context, role *v1alpha1.RoleApplyConfiguration, opts v1.ApplyOptions) (*rbacv1alpha1.Role, error) {
	ret := _m.Called(ctx, role, opts)

	var r0 *rbacv1alpha1.Role
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.RoleApplyConfiguration, v1.ApplyOptions) (*rbacv1alpha1.Role, error)); ok {
		return rf(ctx, role, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.RoleApplyConfiguration, v1.ApplyOptions) *rbacv1alpha1.Role); ok {
		r0 = rf(ctx, role, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1alpha1.Role)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1alpha1.RoleApplyConfiguration, v1.ApplyOptions) error); ok {
		r1 = rf(ctx, role, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoleInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type RoleInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - role *v1alpha1.RoleApplyConfiguration
//   - opts v1.ApplyOptions
func (_e *RoleInterface_Expecter) Apply(ctx interface{}, role interface{}, opts interface{}) *RoleInterface_Apply_Call {
	return &RoleInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, role, opts)}
}

func (_c *RoleInterface_Apply_Call) Run(run func(ctx context.Context, role *v1alpha1.RoleApplyConfiguration, opts v1.ApplyOptions)) *RoleInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1alpha1.RoleApplyConfiguration), args[2].(v1.ApplyOptions))
	})
	return _c
}

func (_c *RoleInterface_Apply_Call) Return(result *rbacv1alpha1.Role, err error) *RoleInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *RoleInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1alpha1.RoleApplyConfiguration, v1.ApplyOptions) (*rbacv1alpha1.Role, error)) *RoleInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, role, opts
func (_m *RoleInterface) Create(ctx context.Context, role *rbacv1alpha1.Role, opts v1.CreateOptions) (*rbacv1alpha1.Role, error) {
	ret := _m.Called(ctx, role, opts)

	var r0 *rbacv1alpha1.Role
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rbacv1alpha1.Role, v1.CreateOptions) (*rbacv1alpha1.Role, error)); ok {
		return rf(ctx, role, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rbacv1alpha1.Role, v1.CreateOptions) *rbacv1alpha1.Role); ok {
		r0 = rf(ctx, role, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1alpha1.Role)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rbacv1alpha1.Role, v1.CreateOptions) error); ok {
		r1 = rf(ctx, role, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoleInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type RoleInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - role *rbacv1alpha1.Role
//   - opts v1.CreateOptions
func (_e *RoleInterface_Expecter) Create(ctx interface{}, role interface{}, opts interface{}) *RoleInterface_Create_Call {
	return &RoleInterface_Create_Call{Call: _e.mock.On("Create", ctx, role, opts)}
}

func (_c *RoleInterface_Create_Call) Run(run func(ctx context.Context, role *rbacv1alpha1.Role, opts v1.CreateOptions)) *RoleInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*rbacv1alpha1.Role), args[2].(v1.CreateOptions))
	})
	return _c
}

func (_c *RoleInterface_Create_Call) Return(_a0 *rbacv1alpha1.Role, _a1 error) *RoleInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RoleInterface_Create_Call) RunAndReturn(run func(context.Context, *rbacv1alpha1.Role, v1.CreateOptions) (*rbacv1alpha1.Role, error)) *RoleInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *RoleInterface) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RoleInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type RoleInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.DeleteOptions
func (_e *RoleInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *RoleInterface_Delete_Call {
	return &RoleInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *RoleInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts v1.DeleteOptions)) *RoleInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.DeleteOptions))
	})
	return _c
}

func (_c *RoleInterface_Delete_Call) Return(_a0 error) *RoleInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RoleInterface_Delete_Call) RunAndReturn(run func(context.Context, string, v1.DeleteOptions) error) *RoleInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *RoleInterface) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.DeleteOptions, v1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RoleInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type RoleInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.DeleteOptions
//   - listOpts v1.ListOptions
func (_e *RoleInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *RoleInterface_DeleteCollection_Call {
	return &RoleInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *RoleInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions)) *RoleInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.DeleteOptions), args[2].(v1.ListOptions))
	})
	return _c
}

func (_c *RoleInterface_DeleteCollection_Call) Return(_a0 error) *RoleInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RoleInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, v1.DeleteOptions, v1.ListOptions) error) *RoleInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *RoleInterface) Get(ctx context.Context, name string, opts v1.GetOptions) (*rbacv1alpha1.Role, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *rbacv1alpha1.Role
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) (*rbacv1alpha1.Role, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) *rbacv1alpha1.Role); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1alpha1.Role)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, v1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoleInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type RoleInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.GetOptions
func (_e *RoleInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *RoleInterface_Get_Call {
	return &RoleInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *RoleInterface_Get_Call) Run(run func(ctx context.Context, name string, opts v1.GetOptions)) *RoleInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.GetOptions))
	})
	return _c
}

func (_c *RoleInterface_Get_Call) Return(_a0 *rbacv1alpha1.Role, _a1 error) *RoleInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RoleInterface_Get_Call) RunAndReturn(run func(context.Context, string, v1.GetOptions) (*rbacv1alpha1.Role, error)) *RoleInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *RoleInterface) List(ctx context.Context, opts v1.ListOptions) (*rbacv1alpha1.RoleList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *rbacv1alpha1.RoleList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) (*rbacv1alpha1.RoleList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) *rbacv1alpha1.RoleList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1alpha1.RoleList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoleInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type RoleInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *RoleInterface_Expecter) List(ctx interface{}, opts interface{}) *RoleInterface_List_Call {
	return &RoleInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *RoleInterface_List_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *RoleInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *RoleInterface_List_Call) Return(_a0 *rbacv1alpha1.RoleList, _a1 error) *RoleInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RoleInterface_List_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (*rbacv1alpha1.RoleList, error)) *RoleInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *RoleInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (*rbacv1alpha1.Role, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *rbacv1alpha1.Role
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*rbacv1alpha1.Role, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) *rbacv1alpha1.Role); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1alpha1.Role)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoleInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type RoleInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts v1.PatchOptions
//   - subresources ...string
func (_e *RoleInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *RoleInterface_Patch_Call {
	return &RoleInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *RoleInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string)) *RoleInterface_Patch_Call {
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

func (_c *RoleInterface_Patch_Call) Return(result *rbacv1alpha1.Role, err error) *RoleInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *RoleInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*rbacv1alpha1.Role, error)) *RoleInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, role, opts
func (_m *RoleInterface) Update(ctx context.Context, role *rbacv1alpha1.Role, opts v1.UpdateOptions) (*rbacv1alpha1.Role, error) {
	ret := _m.Called(ctx, role, opts)

	var r0 *rbacv1alpha1.Role
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rbacv1alpha1.Role, v1.UpdateOptions) (*rbacv1alpha1.Role, error)); ok {
		return rf(ctx, role, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rbacv1alpha1.Role, v1.UpdateOptions) *rbacv1alpha1.Role); ok {
		r0 = rf(ctx, role, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1alpha1.Role)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rbacv1alpha1.Role, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, role, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoleInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type RoleInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - role *rbacv1alpha1.Role
//   - opts v1.UpdateOptions
func (_e *RoleInterface_Expecter) Update(ctx interface{}, role interface{}, opts interface{}) *RoleInterface_Update_Call {
	return &RoleInterface_Update_Call{Call: _e.mock.On("Update", ctx, role, opts)}
}

func (_c *RoleInterface_Update_Call) Run(run func(ctx context.Context, role *rbacv1alpha1.Role, opts v1.UpdateOptions)) *RoleInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*rbacv1alpha1.Role), args[2].(v1.UpdateOptions))
	})
	return _c
}

func (_c *RoleInterface_Update_Call) Return(_a0 *rbacv1alpha1.Role, _a1 error) *RoleInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RoleInterface_Update_Call) RunAndReturn(run func(context.Context, *rbacv1alpha1.Role, v1.UpdateOptions) (*rbacv1alpha1.Role, error)) *RoleInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *RoleInterface) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
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

// RoleInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type RoleInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *RoleInterface_Expecter) Watch(ctx interface{}, opts interface{}) *RoleInterface_Watch_Call {
	return &RoleInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *RoleInterface_Watch_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *RoleInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *RoleInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *RoleInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RoleInterface_Watch_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (watch.Interface, error)) *RoleInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewRoleInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewRoleInterface creates a new instance of RoleInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoleInterface(t mockConstructorTestingTNewRoleInterface) *RoleInterface {
	mock := &RoleInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
