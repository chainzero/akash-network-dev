// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	corev1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/core/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// NamespaceInterface is an autogenerated mock type for the NamespaceInterface type
type NamespaceInterface struct {
	mock.Mock
}

type NamespaceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *NamespaceInterface) EXPECT() *NamespaceInterface_Expecter {
	return &NamespaceInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, namespace, opts
func (_m *NamespaceInterface) Apply(ctx context.Context, namespace *v1.NamespaceApplyConfiguration, opts metav1.ApplyOptions) (*corev1.Namespace, error) {
	ret := _m.Called(ctx, namespace, opts)

	var r0 *corev1.Namespace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.NamespaceApplyConfiguration, metav1.ApplyOptions) (*corev1.Namespace, error)); ok {
		return rf(ctx, namespace, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.NamespaceApplyConfiguration, metav1.ApplyOptions) *corev1.Namespace); ok {
		r0 = rf(ctx, namespace, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.NamespaceApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, namespace, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type NamespaceInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - namespace *v1.NamespaceApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *NamespaceInterface_Expecter) Apply(ctx interface{}, namespace interface{}, opts interface{}) *NamespaceInterface_Apply_Call {
	return &NamespaceInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, namespace, opts)}
}

func (_c *NamespaceInterface_Apply_Call) Run(run func(ctx context.Context, namespace *v1.NamespaceApplyConfiguration, opts metav1.ApplyOptions)) *NamespaceInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.NamespaceApplyConfiguration), args[2].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *NamespaceInterface_Apply_Call) Return(result *corev1.Namespace, err error) *NamespaceInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *NamespaceInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1.NamespaceApplyConfiguration, metav1.ApplyOptions) (*corev1.Namespace, error)) *NamespaceInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// ApplyStatus provides a mock function with given fields: ctx, namespace, opts
func (_m *NamespaceInterface) ApplyStatus(ctx context.Context, namespace *v1.NamespaceApplyConfiguration, opts metav1.ApplyOptions) (*corev1.Namespace, error) {
	ret := _m.Called(ctx, namespace, opts)

	var r0 *corev1.Namespace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.NamespaceApplyConfiguration, metav1.ApplyOptions) (*corev1.Namespace, error)); ok {
		return rf(ctx, namespace, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.NamespaceApplyConfiguration, metav1.ApplyOptions) *corev1.Namespace); ok {
		r0 = rf(ctx, namespace, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.NamespaceApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, namespace, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceInterface_ApplyStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyStatus'
type NamespaceInterface_ApplyStatus_Call struct {
	*mock.Call
}

// ApplyStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - namespace *v1.NamespaceApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *NamespaceInterface_Expecter) ApplyStatus(ctx interface{}, namespace interface{}, opts interface{}) *NamespaceInterface_ApplyStatus_Call {
	return &NamespaceInterface_ApplyStatus_Call{Call: _e.mock.On("ApplyStatus", ctx, namespace, opts)}
}

func (_c *NamespaceInterface_ApplyStatus_Call) Run(run func(ctx context.Context, namespace *v1.NamespaceApplyConfiguration, opts metav1.ApplyOptions)) *NamespaceInterface_ApplyStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.NamespaceApplyConfiguration), args[2].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *NamespaceInterface_ApplyStatus_Call) Return(result *corev1.Namespace, err error) *NamespaceInterface_ApplyStatus_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *NamespaceInterface_ApplyStatus_Call) RunAndReturn(run func(context.Context, *v1.NamespaceApplyConfiguration, metav1.ApplyOptions) (*corev1.Namespace, error)) *NamespaceInterface_ApplyStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, namespace, opts
func (_m *NamespaceInterface) Create(ctx context.Context, namespace *corev1.Namespace, opts metav1.CreateOptions) (*corev1.Namespace, error) {
	ret := _m.Called(ctx, namespace, opts)

	var r0 *corev1.Namespace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Namespace, metav1.CreateOptions) (*corev1.Namespace, error)); ok {
		return rf(ctx, namespace, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Namespace, metav1.CreateOptions) *corev1.Namespace); ok {
		r0 = rf(ctx, namespace, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *corev1.Namespace, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, namespace, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type NamespaceInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - namespace *corev1.Namespace
//   - opts metav1.CreateOptions
func (_e *NamespaceInterface_Expecter) Create(ctx interface{}, namespace interface{}, opts interface{}) *NamespaceInterface_Create_Call {
	return &NamespaceInterface_Create_Call{Call: _e.mock.On("Create", ctx, namespace, opts)}
}

func (_c *NamespaceInterface_Create_Call) Run(run func(ctx context.Context, namespace *corev1.Namespace, opts metav1.CreateOptions)) *NamespaceInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.Namespace), args[2].(metav1.CreateOptions))
	})
	return _c
}

func (_c *NamespaceInterface_Create_Call) Return(_a0 *corev1.Namespace, _a1 error) *NamespaceInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NamespaceInterface_Create_Call) RunAndReturn(run func(context.Context, *corev1.Namespace, metav1.CreateOptions) (*corev1.Namespace, error)) *NamespaceInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *NamespaceInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NamespaceInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type NamespaceInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.DeleteOptions
func (_e *NamespaceInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *NamespaceInterface_Delete_Call {
	return &NamespaceInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *NamespaceInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts metav1.DeleteOptions)) *NamespaceInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.DeleteOptions))
	})
	return _c
}

func (_c *NamespaceInterface_Delete_Call) Return(_a0 error) *NamespaceInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NamespaceInterface_Delete_Call) RunAndReturn(run func(context.Context, string, metav1.DeleteOptions) error) *NamespaceInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Finalize provides a mock function with given fields: ctx, item, opts
func (_m *NamespaceInterface) Finalize(ctx context.Context, item *corev1.Namespace, opts metav1.UpdateOptions) (*corev1.Namespace, error) {
	ret := _m.Called(ctx, item, opts)

	var r0 *corev1.Namespace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Namespace, metav1.UpdateOptions) (*corev1.Namespace, error)); ok {
		return rf(ctx, item, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Namespace, metav1.UpdateOptions) *corev1.Namespace); ok {
		r0 = rf(ctx, item, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *corev1.Namespace, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, item, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceInterface_Finalize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Finalize'
type NamespaceInterface_Finalize_Call struct {
	*mock.Call
}

// Finalize is a helper method to define mock.On call
//   - ctx context.Context
//   - item *corev1.Namespace
//   - opts metav1.UpdateOptions
func (_e *NamespaceInterface_Expecter) Finalize(ctx interface{}, item interface{}, opts interface{}) *NamespaceInterface_Finalize_Call {
	return &NamespaceInterface_Finalize_Call{Call: _e.mock.On("Finalize", ctx, item, opts)}
}

func (_c *NamespaceInterface_Finalize_Call) Run(run func(ctx context.Context, item *corev1.Namespace, opts metav1.UpdateOptions)) *NamespaceInterface_Finalize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.Namespace), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *NamespaceInterface_Finalize_Call) Return(_a0 *corev1.Namespace, _a1 error) *NamespaceInterface_Finalize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NamespaceInterface_Finalize_Call) RunAndReturn(run func(context.Context, *corev1.Namespace, metav1.UpdateOptions) (*corev1.Namespace, error)) *NamespaceInterface_Finalize_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *NamespaceInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.Namespace, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *corev1.Namespace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) (*corev1.Namespace, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *corev1.Namespace); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type NamespaceInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.GetOptions
func (_e *NamespaceInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *NamespaceInterface_Get_Call {
	return &NamespaceInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *NamespaceInterface_Get_Call) Run(run func(ctx context.Context, name string, opts metav1.GetOptions)) *NamespaceInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.GetOptions))
	})
	return _c
}

func (_c *NamespaceInterface_Get_Call) Return(_a0 *corev1.Namespace, _a1 error) *NamespaceInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NamespaceInterface_Get_Call) RunAndReturn(run func(context.Context, string, metav1.GetOptions) (*corev1.Namespace, error)) *NamespaceInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *NamespaceInterface) List(ctx context.Context, opts metav1.ListOptions) (*corev1.NamespaceList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *corev1.NamespaceList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) (*corev1.NamespaceList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *corev1.NamespaceList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.NamespaceList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type NamespaceInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *NamespaceInterface_Expecter) List(ctx interface{}, opts interface{}) *NamespaceInterface_List_Call {
	return &NamespaceInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *NamespaceInterface_List_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *NamespaceInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *NamespaceInterface_List_Call) Return(_a0 *corev1.NamespaceList, _a1 error) *NamespaceInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NamespaceInterface_List_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (*corev1.NamespaceList, error)) *NamespaceInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *NamespaceInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1.Namespace, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *corev1.Namespace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*corev1.Namespace, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *corev1.Namespace); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type NamespaceInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts metav1.PatchOptions
//   - subresources ...string
func (_e *NamespaceInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *NamespaceInterface_Patch_Call {
	return &NamespaceInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *NamespaceInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string)) *NamespaceInterface_Patch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-5)
		for i, a := range args[5:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(types.PatchType), args[3].([]byte), args[4].(metav1.PatchOptions), variadicArgs...)
	})
	return _c
}

func (_c *NamespaceInterface_Patch_Call) Return(result *corev1.Namespace, err error) *NamespaceInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *NamespaceInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*corev1.Namespace, error)) *NamespaceInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, namespace, opts
func (_m *NamespaceInterface) Update(ctx context.Context, namespace *corev1.Namespace, opts metav1.UpdateOptions) (*corev1.Namespace, error) {
	ret := _m.Called(ctx, namespace, opts)

	var r0 *corev1.Namespace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Namespace, metav1.UpdateOptions) (*corev1.Namespace, error)); ok {
		return rf(ctx, namespace, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Namespace, metav1.UpdateOptions) *corev1.Namespace); ok {
		r0 = rf(ctx, namespace, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *corev1.Namespace, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, namespace, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type NamespaceInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - namespace *corev1.Namespace
//   - opts metav1.UpdateOptions
func (_e *NamespaceInterface_Expecter) Update(ctx interface{}, namespace interface{}, opts interface{}) *NamespaceInterface_Update_Call {
	return &NamespaceInterface_Update_Call{Call: _e.mock.On("Update", ctx, namespace, opts)}
}

func (_c *NamespaceInterface_Update_Call) Run(run func(ctx context.Context, namespace *corev1.Namespace, opts metav1.UpdateOptions)) *NamespaceInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.Namespace), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *NamespaceInterface_Update_Call) Return(_a0 *corev1.Namespace, _a1 error) *NamespaceInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NamespaceInterface_Update_Call) RunAndReturn(run func(context.Context, *corev1.Namespace, metav1.UpdateOptions) (*corev1.Namespace, error)) *NamespaceInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStatus provides a mock function with given fields: ctx, namespace, opts
func (_m *NamespaceInterface) UpdateStatus(ctx context.Context, namespace *corev1.Namespace, opts metav1.UpdateOptions) (*corev1.Namespace, error) {
	ret := _m.Called(ctx, namespace, opts)

	var r0 *corev1.Namespace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Namespace, metav1.UpdateOptions) (*corev1.Namespace, error)); ok {
		return rf(ctx, namespace, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Namespace, metav1.UpdateOptions) *corev1.Namespace); ok {
		r0 = rf(ctx, namespace, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *corev1.Namespace, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, namespace, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceInterface_UpdateStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStatus'
type NamespaceInterface_UpdateStatus_Call struct {
	*mock.Call
}

// UpdateStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - namespace *corev1.Namespace
//   - opts metav1.UpdateOptions
func (_e *NamespaceInterface_Expecter) UpdateStatus(ctx interface{}, namespace interface{}, opts interface{}) *NamespaceInterface_UpdateStatus_Call {
	return &NamespaceInterface_UpdateStatus_Call{Call: _e.mock.On("UpdateStatus", ctx, namespace, opts)}
}

func (_c *NamespaceInterface_UpdateStatus_Call) Run(run func(ctx context.Context, namespace *corev1.Namespace, opts metav1.UpdateOptions)) *NamespaceInterface_UpdateStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.Namespace), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *NamespaceInterface_UpdateStatus_Call) Return(_a0 *corev1.Namespace, _a1 error) *NamespaceInterface_UpdateStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NamespaceInterface_UpdateStatus_Call) RunAndReturn(run func(context.Context, *corev1.Namespace, metav1.UpdateOptions) (*corev1.Namespace, error)) *NamespaceInterface_UpdateStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *NamespaceInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) (watch.Interface, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type NamespaceInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *NamespaceInterface_Expecter) Watch(ctx interface{}, opts interface{}) *NamespaceInterface_Watch_Call {
	return &NamespaceInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *NamespaceInterface_Watch_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *NamespaceInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *NamespaceInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *NamespaceInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NamespaceInterface_Watch_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (watch.Interface, error)) *NamespaceInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewNamespaceInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewNamespaceInterface creates a new instance of NamespaceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNamespaceInterface(t mockConstructorTestingTNewNamespaceInterface) *NamespaceInterface {
	mock := &NamespaceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
