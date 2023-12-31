// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	resourcev1alpha1 "k8s.io/api/resource/v1alpha1"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "k8s.io/client-go/applyconfigurations/resource/v1alpha1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// ResourceClassInterface is an autogenerated mock type for the ResourceClassInterface type
type ResourceClassInterface struct {
	mock.Mock
}

type ResourceClassInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ResourceClassInterface) EXPECT() *ResourceClassInterface_Expecter {
	return &ResourceClassInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, resourceClass, opts
func (_m *ResourceClassInterface) Apply(ctx context.Context, resourceClass *v1alpha1.ResourceClassApplyConfiguration, opts v1.ApplyOptions) (*resourcev1alpha1.ResourceClass, error) {
	ret := _m.Called(ctx, resourceClass, opts)

	var r0 *resourcev1alpha1.ResourceClass
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.ResourceClassApplyConfiguration, v1.ApplyOptions) (*resourcev1alpha1.ResourceClass, error)); ok {
		return rf(ctx, resourceClass, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.ResourceClassApplyConfiguration, v1.ApplyOptions) *resourcev1alpha1.ResourceClass); ok {
		r0 = rf(ctx, resourceClass, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcev1alpha1.ResourceClass)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1alpha1.ResourceClassApplyConfiguration, v1.ApplyOptions) error); ok {
		r1 = rf(ctx, resourceClass, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceClassInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type ResourceClassInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - resourceClass *v1alpha1.ResourceClassApplyConfiguration
//   - opts v1.ApplyOptions
func (_e *ResourceClassInterface_Expecter) Apply(ctx interface{}, resourceClass interface{}, opts interface{}) *ResourceClassInterface_Apply_Call {
	return &ResourceClassInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, resourceClass, opts)}
}

func (_c *ResourceClassInterface_Apply_Call) Run(run func(ctx context.Context, resourceClass *v1alpha1.ResourceClassApplyConfiguration, opts v1.ApplyOptions)) *ResourceClassInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1alpha1.ResourceClassApplyConfiguration), args[2].(v1.ApplyOptions))
	})
	return _c
}

func (_c *ResourceClassInterface_Apply_Call) Return(result *resourcev1alpha1.ResourceClass, err error) *ResourceClassInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *ResourceClassInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1alpha1.ResourceClassApplyConfiguration, v1.ApplyOptions) (*resourcev1alpha1.ResourceClass, error)) *ResourceClassInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, resourceClass, opts
func (_m *ResourceClassInterface) Create(ctx context.Context, resourceClass *resourcev1alpha1.ResourceClass, opts v1.CreateOptions) (*resourcev1alpha1.ResourceClass, error) {
	ret := _m.Called(ctx, resourceClass, opts)

	var r0 *resourcev1alpha1.ResourceClass
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcev1alpha1.ResourceClass, v1.CreateOptions) (*resourcev1alpha1.ResourceClass, error)); ok {
		return rf(ctx, resourceClass, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcev1alpha1.ResourceClass, v1.CreateOptions) *resourcev1alpha1.ResourceClass); ok {
		r0 = rf(ctx, resourceClass, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcev1alpha1.ResourceClass)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcev1alpha1.ResourceClass, v1.CreateOptions) error); ok {
		r1 = rf(ctx, resourceClass, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceClassInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ResourceClassInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - resourceClass *resourcev1alpha1.ResourceClass
//   - opts v1.CreateOptions
func (_e *ResourceClassInterface_Expecter) Create(ctx interface{}, resourceClass interface{}, opts interface{}) *ResourceClassInterface_Create_Call {
	return &ResourceClassInterface_Create_Call{Call: _e.mock.On("Create", ctx, resourceClass, opts)}
}

func (_c *ResourceClassInterface_Create_Call) Run(run func(ctx context.Context, resourceClass *resourcev1alpha1.ResourceClass, opts v1.CreateOptions)) *ResourceClassInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*resourcev1alpha1.ResourceClass), args[2].(v1.CreateOptions))
	})
	return _c
}

func (_c *ResourceClassInterface_Create_Call) Return(_a0 *resourcev1alpha1.ResourceClass, _a1 error) *ResourceClassInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceClassInterface_Create_Call) RunAndReturn(run func(context.Context, *resourcev1alpha1.ResourceClass, v1.CreateOptions) (*resourcev1alpha1.ResourceClass, error)) *ResourceClassInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *ResourceClassInterface) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResourceClassInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ResourceClassInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.DeleteOptions
func (_e *ResourceClassInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *ResourceClassInterface_Delete_Call {
	return &ResourceClassInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *ResourceClassInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts v1.DeleteOptions)) *ResourceClassInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.DeleteOptions))
	})
	return _c
}

func (_c *ResourceClassInterface_Delete_Call) Return(_a0 error) *ResourceClassInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResourceClassInterface_Delete_Call) RunAndReturn(run func(context.Context, string, v1.DeleteOptions) error) *ResourceClassInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *ResourceClassInterface) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.DeleteOptions, v1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResourceClassInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type ResourceClassInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.DeleteOptions
//   - listOpts v1.ListOptions
func (_e *ResourceClassInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *ResourceClassInterface_DeleteCollection_Call {
	return &ResourceClassInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *ResourceClassInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions)) *ResourceClassInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.DeleteOptions), args[2].(v1.ListOptions))
	})
	return _c
}

func (_c *ResourceClassInterface_DeleteCollection_Call) Return(_a0 error) *ResourceClassInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResourceClassInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, v1.DeleteOptions, v1.ListOptions) error) *ResourceClassInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *ResourceClassInterface) Get(ctx context.Context, name string, opts v1.GetOptions) (*resourcev1alpha1.ResourceClass, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *resourcev1alpha1.ResourceClass
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) (*resourcev1alpha1.ResourceClass, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) *resourcev1alpha1.ResourceClass); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcev1alpha1.ResourceClass)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, v1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceClassInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ResourceClassInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.GetOptions
func (_e *ResourceClassInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *ResourceClassInterface_Get_Call {
	return &ResourceClassInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *ResourceClassInterface_Get_Call) Run(run func(ctx context.Context, name string, opts v1.GetOptions)) *ResourceClassInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.GetOptions))
	})
	return _c
}

func (_c *ResourceClassInterface_Get_Call) Return(_a0 *resourcev1alpha1.ResourceClass, _a1 error) *ResourceClassInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceClassInterface_Get_Call) RunAndReturn(run func(context.Context, string, v1.GetOptions) (*resourcev1alpha1.ResourceClass, error)) *ResourceClassInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *ResourceClassInterface) List(ctx context.Context, opts v1.ListOptions) (*resourcev1alpha1.ResourceClassList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *resourcev1alpha1.ResourceClassList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) (*resourcev1alpha1.ResourceClassList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) *resourcev1alpha1.ResourceClassList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcev1alpha1.ResourceClassList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceClassInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ResourceClassInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *ResourceClassInterface_Expecter) List(ctx interface{}, opts interface{}) *ResourceClassInterface_List_Call {
	return &ResourceClassInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *ResourceClassInterface_List_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *ResourceClassInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *ResourceClassInterface_List_Call) Return(_a0 *resourcev1alpha1.ResourceClassList, _a1 error) *ResourceClassInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceClassInterface_List_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (*resourcev1alpha1.ResourceClassList, error)) *ResourceClassInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *ResourceClassInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (*resourcev1alpha1.ResourceClass, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *resourcev1alpha1.ResourceClass
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*resourcev1alpha1.ResourceClass, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) *resourcev1alpha1.ResourceClass); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcev1alpha1.ResourceClass)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceClassInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type ResourceClassInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts v1.PatchOptions
//   - subresources ...string
func (_e *ResourceClassInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *ResourceClassInterface_Patch_Call {
	return &ResourceClassInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *ResourceClassInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string)) *ResourceClassInterface_Patch_Call {
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

func (_c *ResourceClassInterface_Patch_Call) Return(result *resourcev1alpha1.ResourceClass, err error) *ResourceClassInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *ResourceClassInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*resourcev1alpha1.ResourceClass, error)) *ResourceClassInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, resourceClass, opts
func (_m *ResourceClassInterface) Update(ctx context.Context, resourceClass *resourcev1alpha1.ResourceClass, opts v1.UpdateOptions) (*resourcev1alpha1.ResourceClass, error) {
	ret := _m.Called(ctx, resourceClass, opts)

	var r0 *resourcev1alpha1.ResourceClass
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcev1alpha1.ResourceClass, v1.UpdateOptions) (*resourcev1alpha1.ResourceClass, error)); ok {
		return rf(ctx, resourceClass, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcev1alpha1.ResourceClass, v1.UpdateOptions) *resourcev1alpha1.ResourceClass); ok {
		r0 = rf(ctx, resourceClass, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcev1alpha1.ResourceClass)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcev1alpha1.ResourceClass, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, resourceClass, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceClassInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ResourceClassInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - resourceClass *resourcev1alpha1.ResourceClass
//   - opts v1.UpdateOptions
func (_e *ResourceClassInterface_Expecter) Update(ctx interface{}, resourceClass interface{}, opts interface{}) *ResourceClassInterface_Update_Call {
	return &ResourceClassInterface_Update_Call{Call: _e.mock.On("Update", ctx, resourceClass, opts)}
}

func (_c *ResourceClassInterface_Update_Call) Run(run func(ctx context.Context, resourceClass *resourcev1alpha1.ResourceClass, opts v1.UpdateOptions)) *ResourceClassInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*resourcev1alpha1.ResourceClass), args[2].(v1.UpdateOptions))
	})
	return _c
}

func (_c *ResourceClassInterface_Update_Call) Return(_a0 *resourcev1alpha1.ResourceClass, _a1 error) *ResourceClassInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceClassInterface_Update_Call) RunAndReturn(run func(context.Context, *resourcev1alpha1.ResourceClass, v1.UpdateOptions) (*resourcev1alpha1.ResourceClass, error)) *ResourceClassInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *ResourceClassInterface) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
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

// ResourceClassInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type ResourceClassInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *ResourceClassInterface_Expecter) Watch(ctx interface{}, opts interface{}) *ResourceClassInterface_Watch_Call {
	return &ResourceClassInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *ResourceClassInterface_Watch_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *ResourceClassInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *ResourceClassInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *ResourceClassInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceClassInterface_Watch_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (watch.Interface, error)) *ResourceClassInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewResourceClassInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewResourceClassInterface creates a new instance of ResourceClassInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewResourceClassInterface(t mockConstructorTestingTNewResourceClassInterface) *ResourceClassInterface {
	mock := &ResourceClassInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
