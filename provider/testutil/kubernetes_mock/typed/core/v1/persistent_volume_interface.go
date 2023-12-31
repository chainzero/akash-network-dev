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

// PersistentVolumeInterface is an autogenerated mock type for the PersistentVolumeInterface type
type PersistentVolumeInterface struct {
	mock.Mock
}

type PersistentVolumeInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *PersistentVolumeInterface) EXPECT() *PersistentVolumeInterface_Expecter {
	return &PersistentVolumeInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, persistentVolume, opts
func (_m *PersistentVolumeInterface) Apply(ctx context.Context, persistentVolume *v1.PersistentVolumeApplyConfiguration, opts metav1.ApplyOptions) (*corev1.PersistentVolume, error) {
	ret := _m.Called(ctx, persistentVolume, opts)

	var r0 *corev1.PersistentVolume
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.PersistentVolumeApplyConfiguration, metav1.ApplyOptions) (*corev1.PersistentVolume, error)); ok {
		return rf(ctx, persistentVolume, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.PersistentVolumeApplyConfiguration, metav1.ApplyOptions) *corev1.PersistentVolume); ok {
		r0 = rf(ctx, persistentVolume, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PersistentVolume)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.PersistentVolumeApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, persistentVolume, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersistentVolumeInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type PersistentVolumeInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - persistentVolume *v1.PersistentVolumeApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *PersistentVolumeInterface_Expecter) Apply(ctx interface{}, persistentVolume interface{}, opts interface{}) *PersistentVolumeInterface_Apply_Call {
	return &PersistentVolumeInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, persistentVolume, opts)}
}

func (_c *PersistentVolumeInterface_Apply_Call) Run(run func(ctx context.Context, persistentVolume *v1.PersistentVolumeApplyConfiguration, opts metav1.ApplyOptions)) *PersistentVolumeInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.PersistentVolumeApplyConfiguration), args[2].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *PersistentVolumeInterface_Apply_Call) Return(result *corev1.PersistentVolume, err error) *PersistentVolumeInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *PersistentVolumeInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1.PersistentVolumeApplyConfiguration, metav1.ApplyOptions) (*corev1.PersistentVolume, error)) *PersistentVolumeInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// ApplyStatus provides a mock function with given fields: ctx, persistentVolume, opts
func (_m *PersistentVolumeInterface) ApplyStatus(ctx context.Context, persistentVolume *v1.PersistentVolumeApplyConfiguration, opts metav1.ApplyOptions) (*corev1.PersistentVolume, error) {
	ret := _m.Called(ctx, persistentVolume, opts)

	var r0 *corev1.PersistentVolume
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.PersistentVolumeApplyConfiguration, metav1.ApplyOptions) (*corev1.PersistentVolume, error)); ok {
		return rf(ctx, persistentVolume, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.PersistentVolumeApplyConfiguration, metav1.ApplyOptions) *corev1.PersistentVolume); ok {
		r0 = rf(ctx, persistentVolume, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PersistentVolume)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.PersistentVolumeApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, persistentVolume, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersistentVolumeInterface_ApplyStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyStatus'
type PersistentVolumeInterface_ApplyStatus_Call struct {
	*mock.Call
}

// ApplyStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - persistentVolume *v1.PersistentVolumeApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *PersistentVolumeInterface_Expecter) ApplyStatus(ctx interface{}, persistentVolume interface{}, opts interface{}) *PersistentVolumeInterface_ApplyStatus_Call {
	return &PersistentVolumeInterface_ApplyStatus_Call{Call: _e.mock.On("ApplyStatus", ctx, persistentVolume, opts)}
}

func (_c *PersistentVolumeInterface_ApplyStatus_Call) Run(run func(ctx context.Context, persistentVolume *v1.PersistentVolumeApplyConfiguration, opts metav1.ApplyOptions)) *PersistentVolumeInterface_ApplyStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.PersistentVolumeApplyConfiguration), args[2].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *PersistentVolumeInterface_ApplyStatus_Call) Return(result *corev1.PersistentVolume, err error) *PersistentVolumeInterface_ApplyStatus_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *PersistentVolumeInterface_ApplyStatus_Call) RunAndReturn(run func(context.Context, *v1.PersistentVolumeApplyConfiguration, metav1.ApplyOptions) (*corev1.PersistentVolume, error)) *PersistentVolumeInterface_ApplyStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, persistentVolume, opts
func (_m *PersistentVolumeInterface) Create(ctx context.Context, persistentVolume *corev1.PersistentVolume, opts metav1.CreateOptions) (*corev1.PersistentVolume, error) {
	ret := _m.Called(ctx, persistentVolume, opts)

	var r0 *corev1.PersistentVolume
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.PersistentVolume, metav1.CreateOptions) (*corev1.PersistentVolume, error)); ok {
		return rf(ctx, persistentVolume, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.PersistentVolume, metav1.CreateOptions) *corev1.PersistentVolume); ok {
		r0 = rf(ctx, persistentVolume, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PersistentVolume)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *corev1.PersistentVolume, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, persistentVolume, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersistentVolumeInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type PersistentVolumeInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - persistentVolume *corev1.PersistentVolume
//   - opts metav1.CreateOptions
func (_e *PersistentVolumeInterface_Expecter) Create(ctx interface{}, persistentVolume interface{}, opts interface{}) *PersistentVolumeInterface_Create_Call {
	return &PersistentVolumeInterface_Create_Call{Call: _e.mock.On("Create", ctx, persistentVolume, opts)}
}

func (_c *PersistentVolumeInterface_Create_Call) Run(run func(ctx context.Context, persistentVolume *corev1.PersistentVolume, opts metav1.CreateOptions)) *PersistentVolumeInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.PersistentVolume), args[2].(metav1.CreateOptions))
	})
	return _c
}

func (_c *PersistentVolumeInterface_Create_Call) Return(_a0 *corev1.PersistentVolume, _a1 error) *PersistentVolumeInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PersistentVolumeInterface_Create_Call) RunAndReturn(run func(context.Context, *corev1.PersistentVolume, metav1.CreateOptions) (*corev1.PersistentVolume, error)) *PersistentVolumeInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *PersistentVolumeInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PersistentVolumeInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type PersistentVolumeInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.DeleteOptions
func (_e *PersistentVolumeInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *PersistentVolumeInterface_Delete_Call {
	return &PersistentVolumeInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *PersistentVolumeInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts metav1.DeleteOptions)) *PersistentVolumeInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.DeleteOptions))
	})
	return _c
}

func (_c *PersistentVolumeInterface_Delete_Call) Return(_a0 error) *PersistentVolumeInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PersistentVolumeInterface_Delete_Call) RunAndReturn(run func(context.Context, string, metav1.DeleteOptions) error) *PersistentVolumeInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *PersistentVolumeInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PersistentVolumeInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type PersistentVolumeInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.DeleteOptions
//   - listOpts metav1.ListOptions
func (_e *PersistentVolumeInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *PersistentVolumeInterface_DeleteCollection_Call {
	return &PersistentVolumeInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *PersistentVolumeInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions)) *PersistentVolumeInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.DeleteOptions), args[2].(metav1.ListOptions))
	})
	return _c
}

func (_c *PersistentVolumeInterface_DeleteCollection_Call) Return(_a0 error) *PersistentVolumeInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PersistentVolumeInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error) *PersistentVolumeInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *PersistentVolumeInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.PersistentVolume, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *corev1.PersistentVolume
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) (*corev1.PersistentVolume, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *corev1.PersistentVolume); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PersistentVolume)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersistentVolumeInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type PersistentVolumeInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.GetOptions
func (_e *PersistentVolumeInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *PersistentVolumeInterface_Get_Call {
	return &PersistentVolumeInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *PersistentVolumeInterface_Get_Call) Run(run func(ctx context.Context, name string, opts metav1.GetOptions)) *PersistentVolumeInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.GetOptions))
	})
	return _c
}

func (_c *PersistentVolumeInterface_Get_Call) Return(_a0 *corev1.PersistentVolume, _a1 error) *PersistentVolumeInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PersistentVolumeInterface_Get_Call) RunAndReturn(run func(context.Context, string, metav1.GetOptions) (*corev1.PersistentVolume, error)) *PersistentVolumeInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *PersistentVolumeInterface) List(ctx context.Context, opts metav1.ListOptions) (*corev1.PersistentVolumeList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *corev1.PersistentVolumeList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) (*corev1.PersistentVolumeList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *corev1.PersistentVolumeList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PersistentVolumeList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersistentVolumeInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type PersistentVolumeInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *PersistentVolumeInterface_Expecter) List(ctx interface{}, opts interface{}) *PersistentVolumeInterface_List_Call {
	return &PersistentVolumeInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *PersistentVolumeInterface_List_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *PersistentVolumeInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *PersistentVolumeInterface_List_Call) Return(_a0 *corev1.PersistentVolumeList, _a1 error) *PersistentVolumeInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PersistentVolumeInterface_List_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (*corev1.PersistentVolumeList, error)) *PersistentVolumeInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *PersistentVolumeInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1.PersistentVolume, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *corev1.PersistentVolume
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*corev1.PersistentVolume, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *corev1.PersistentVolume); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PersistentVolume)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersistentVolumeInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type PersistentVolumeInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts metav1.PatchOptions
//   - subresources ...string
func (_e *PersistentVolumeInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *PersistentVolumeInterface_Patch_Call {
	return &PersistentVolumeInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *PersistentVolumeInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string)) *PersistentVolumeInterface_Patch_Call {
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

func (_c *PersistentVolumeInterface_Patch_Call) Return(result *corev1.PersistentVolume, err error) *PersistentVolumeInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *PersistentVolumeInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*corev1.PersistentVolume, error)) *PersistentVolumeInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, persistentVolume, opts
func (_m *PersistentVolumeInterface) Update(ctx context.Context, persistentVolume *corev1.PersistentVolume, opts metav1.UpdateOptions) (*corev1.PersistentVolume, error) {
	ret := _m.Called(ctx, persistentVolume, opts)

	var r0 *corev1.PersistentVolume
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.PersistentVolume, metav1.UpdateOptions) (*corev1.PersistentVolume, error)); ok {
		return rf(ctx, persistentVolume, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.PersistentVolume, metav1.UpdateOptions) *corev1.PersistentVolume); ok {
		r0 = rf(ctx, persistentVolume, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PersistentVolume)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *corev1.PersistentVolume, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, persistentVolume, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersistentVolumeInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type PersistentVolumeInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - persistentVolume *corev1.PersistentVolume
//   - opts metav1.UpdateOptions
func (_e *PersistentVolumeInterface_Expecter) Update(ctx interface{}, persistentVolume interface{}, opts interface{}) *PersistentVolumeInterface_Update_Call {
	return &PersistentVolumeInterface_Update_Call{Call: _e.mock.On("Update", ctx, persistentVolume, opts)}
}

func (_c *PersistentVolumeInterface_Update_Call) Run(run func(ctx context.Context, persistentVolume *corev1.PersistentVolume, opts metav1.UpdateOptions)) *PersistentVolumeInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.PersistentVolume), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *PersistentVolumeInterface_Update_Call) Return(_a0 *corev1.PersistentVolume, _a1 error) *PersistentVolumeInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PersistentVolumeInterface_Update_Call) RunAndReturn(run func(context.Context, *corev1.PersistentVolume, metav1.UpdateOptions) (*corev1.PersistentVolume, error)) *PersistentVolumeInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStatus provides a mock function with given fields: ctx, persistentVolume, opts
func (_m *PersistentVolumeInterface) UpdateStatus(ctx context.Context, persistentVolume *corev1.PersistentVolume, opts metav1.UpdateOptions) (*corev1.PersistentVolume, error) {
	ret := _m.Called(ctx, persistentVolume, opts)

	var r0 *corev1.PersistentVolume
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.PersistentVolume, metav1.UpdateOptions) (*corev1.PersistentVolume, error)); ok {
		return rf(ctx, persistentVolume, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.PersistentVolume, metav1.UpdateOptions) *corev1.PersistentVolume); ok {
		r0 = rf(ctx, persistentVolume, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PersistentVolume)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *corev1.PersistentVolume, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, persistentVolume, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersistentVolumeInterface_UpdateStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStatus'
type PersistentVolumeInterface_UpdateStatus_Call struct {
	*mock.Call
}

// UpdateStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - persistentVolume *corev1.PersistentVolume
//   - opts metav1.UpdateOptions
func (_e *PersistentVolumeInterface_Expecter) UpdateStatus(ctx interface{}, persistentVolume interface{}, opts interface{}) *PersistentVolumeInterface_UpdateStatus_Call {
	return &PersistentVolumeInterface_UpdateStatus_Call{Call: _e.mock.On("UpdateStatus", ctx, persistentVolume, opts)}
}

func (_c *PersistentVolumeInterface_UpdateStatus_Call) Run(run func(ctx context.Context, persistentVolume *corev1.PersistentVolume, opts metav1.UpdateOptions)) *PersistentVolumeInterface_UpdateStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.PersistentVolume), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *PersistentVolumeInterface_UpdateStatus_Call) Return(_a0 *corev1.PersistentVolume, _a1 error) *PersistentVolumeInterface_UpdateStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PersistentVolumeInterface_UpdateStatus_Call) RunAndReturn(run func(context.Context, *corev1.PersistentVolume, metav1.UpdateOptions) (*corev1.PersistentVolume, error)) *PersistentVolumeInterface_UpdateStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *PersistentVolumeInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
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

// PersistentVolumeInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type PersistentVolumeInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *PersistentVolumeInterface_Expecter) Watch(ctx interface{}, opts interface{}) *PersistentVolumeInterface_Watch_Call {
	return &PersistentVolumeInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *PersistentVolumeInterface_Watch_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *PersistentVolumeInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *PersistentVolumeInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *PersistentVolumeInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PersistentVolumeInterface_Watch_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (watch.Interface, error)) *PersistentVolumeInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewPersistentVolumeInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewPersistentVolumeInterface creates a new instance of PersistentVolumeInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPersistentVolumeInterface(t mockConstructorTestingTNewPersistentVolumeInterface) *PersistentVolumeInterface {
	mock := &PersistentVolumeInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
