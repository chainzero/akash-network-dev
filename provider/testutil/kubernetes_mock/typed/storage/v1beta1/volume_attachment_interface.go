// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	storagev1beta1 "k8s.io/api/storage/v1beta1"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta1 "k8s.io/client-go/applyconfigurations/storage/v1beta1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// VolumeAttachmentInterface is an autogenerated mock type for the VolumeAttachmentInterface type
type VolumeAttachmentInterface struct {
	mock.Mock
}

type VolumeAttachmentInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *VolumeAttachmentInterface) EXPECT() *VolumeAttachmentInterface_Expecter {
	return &VolumeAttachmentInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, volumeAttachment, opts
func (_m *VolumeAttachmentInterface) Apply(ctx context.Context, volumeAttachment *v1beta1.VolumeAttachmentApplyConfiguration, opts v1.ApplyOptions) (*storagev1beta1.VolumeAttachment, error) {
	ret := _m.Called(ctx, volumeAttachment, opts)

	var r0 *storagev1beta1.VolumeAttachment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.VolumeAttachmentApplyConfiguration, v1.ApplyOptions) (*storagev1beta1.VolumeAttachment, error)); ok {
		return rf(ctx, volumeAttachment, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.VolumeAttachmentApplyConfiguration, v1.ApplyOptions) *storagev1beta1.VolumeAttachment); ok {
		r0 = rf(ctx, volumeAttachment, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1beta1.VolumeAttachment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.VolumeAttachmentApplyConfiguration, v1.ApplyOptions) error); ok {
		r1 = rf(ctx, volumeAttachment, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VolumeAttachmentInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type VolumeAttachmentInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - volumeAttachment *v1beta1.VolumeAttachmentApplyConfiguration
//   - opts v1.ApplyOptions
func (_e *VolumeAttachmentInterface_Expecter) Apply(ctx interface{}, volumeAttachment interface{}, opts interface{}) *VolumeAttachmentInterface_Apply_Call {
	return &VolumeAttachmentInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, volumeAttachment, opts)}
}

func (_c *VolumeAttachmentInterface_Apply_Call) Run(run func(ctx context.Context, volumeAttachment *v1beta1.VolumeAttachmentApplyConfiguration, opts v1.ApplyOptions)) *VolumeAttachmentInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1beta1.VolumeAttachmentApplyConfiguration), args[2].(v1.ApplyOptions))
	})
	return _c
}

func (_c *VolumeAttachmentInterface_Apply_Call) Return(result *storagev1beta1.VolumeAttachment, err error) *VolumeAttachmentInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *VolumeAttachmentInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1beta1.VolumeAttachmentApplyConfiguration, v1.ApplyOptions) (*storagev1beta1.VolumeAttachment, error)) *VolumeAttachmentInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// ApplyStatus provides a mock function with given fields: ctx, volumeAttachment, opts
func (_m *VolumeAttachmentInterface) ApplyStatus(ctx context.Context, volumeAttachment *v1beta1.VolumeAttachmentApplyConfiguration, opts v1.ApplyOptions) (*storagev1beta1.VolumeAttachment, error) {
	ret := _m.Called(ctx, volumeAttachment, opts)

	var r0 *storagev1beta1.VolumeAttachment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.VolumeAttachmentApplyConfiguration, v1.ApplyOptions) (*storagev1beta1.VolumeAttachment, error)); ok {
		return rf(ctx, volumeAttachment, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.VolumeAttachmentApplyConfiguration, v1.ApplyOptions) *storagev1beta1.VolumeAttachment); ok {
		r0 = rf(ctx, volumeAttachment, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1beta1.VolumeAttachment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.VolumeAttachmentApplyConfiguration, v1.ApplyOptions) error); ok {
		r1 = rf(ctx, volumeAttachment, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VolumeAttachmentInterface_ApplyStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyStatus'
type VolumeAttachmentInterface_ApplyStatus_Call struct {
	*mock.Call
}

// ApplyStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - volumeAttachment *v1beta1.VolumeAttachmentApplyConfiguration
//   - opts v1.ApplyOptions
func (_e *VolumeAttachmentInterface_Expecter) ApplyStatus(ctx interface{}, volumeAttachment interface{}, opts interface{}) *VolumeAttachmentInterface_ApplyStatus_Call {
	return &VolumeAttachmentInterface_ApplyStatus_Call{Call: _e.mock.On("ApplyStatus", ctx, volumeAttachment, opts)}
}

func (_c *VolumeAttachmentInterface_ApplyStatus_Call) Run(run func(ctx context.Context, volumeAttachment *v1beta1.VolumeAttachmentApplyConfiguration, opts v1.ApplyOptions)) *VolumeAttachmentInterface_ApplyStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1beta1.VolumeAttachmentApplyConfiguration), args[2].(v1.ApplyOptions))
	})
	return _c
}

func (_c *VolumeAttachmentInterface_ApplyStatus_Call) Return(result *storagev1beta1.VolumeAttachment, err error) *VolumeAttachmentInterface_ApplyStatus_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *VolumeAttachmentInterface_ApplyStatus_Call) RunAndReturn(run func(context.Context, *v1beta1.VolumeAttachmentApplyConfiguration, v1.ApplyOptions) (*storagev1beta1.VolumeAttachment, error)) *VolumeAttachmentInterface_ApplyStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, volumeAttachment, opts
func (_m *VolumeAttachmentInterface) Create(ctx context.Context, volumeAttachment *storagev1beta1.VolumeAttachment, opts v1.CreateOptions) (*storagev1beta1.VolumeAttachment, error) {
	ret := _m.Called(ctx, volumeAttachment, opts)

	var r0 *storagev1beta1.VolumeAttachment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storagev1beta1.VolumeAttachment, v1.CreateOptions) (*storagev1beta1.VolumeAttachment, error)); ok {
		return rf(ctx, volumeAttachment, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storagev1beta1.VolumeAttachment, v1.CreateOptions) *storagev1beta1.VolumeAttachment); ok {
		r0 = rf(ctx, volumeAttachment, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1beta1.VolumeAttachment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storagev1beta1.VolumeAttachment, v1.CreateOptions) error); ok {
		r1 = rf(ctx, volumeAttachment, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VolumeAttachmentInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type VolumeAttachmentInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - volumeAttachment *storagev1beta1.VolumeAttachment
//   - opts v1.CreateOptions
func (_e *VolumeAttachmentInterface_Expecter) Create(ctx interface{}, volumeAttachment interface{}, opts interface{}) *VolumeAttachmentInterface_Create_Call {
	return &VolumeAttachmentInterface_Create_Call{Call: _e.mock.On("Create", ctx, volumeAttachment, opts)}
}

func (_c *VolumeAttachmentInterface_Create_Call) Run(run func(ctx context.Context, volumeAttachment *storagev1beta1.VolumeAttachment, opts v1.CreateOptions)) *VolumeAttachmentInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*storagev1beta1.VolumeAttachment), args[2].(v1.CreateOptions))
	})
	return _c
}

func (_c *VolumeAttachmentInterface_Create_Call) Return(_a0 *storagev1beta1.VolumeAttachment, _a1 error) *VolumeAttachmentInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VolumeAttachmentInterface_Create_Call) RunAndReturn(run func(context.Context, *storagev1beta1.VolumeAttachment, v1.CreateOptions) (*storagev1beta1.VolumeAttachment, error)) *VolumeAttachmentInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *VolumeAttachmentInterface) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VolumeAttachmentInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type VolumeAttachmentInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.DeleteOptions
func (_e *VolumeAttachmentInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *VolumeAttachmentInterface_Delete_Call {
	return &VolumeAttachmentInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *VolumeAttachmentInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts v1.DeleteOptions)) *VolumeAttachmentInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.DeleteOptions))
	})
	return _c
}

func (_c *VolumeAttachmentInterface_Delete_Call) Return(_a0 error) *VolumeAttachmentInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *VolumeAttachmentInterface_Delete_Call) RunAndReturn(run func(context.Context, string, v1.DeleteOptions) error) *VolumeAttachmentInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *VolumeAttachmentInterface) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.DeleteOptions, v1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VolumeAttachmentInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type VolumeAttachmentInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.DeleteOptions
//   - listOpts v1.ListOptions
func (_e *VolumeAttachmentInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *VolumeAttachmentInterface_DeleteCollection_Call {
	return &VolumeAttachmentInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *VolumeAttachmentInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions)) *VolumeAttachmentInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.DeleteOptions), args[2].(v1.ListOptions))
	})
	return _c
}

func (_c *VolumeAttachmentInterface_DeleteCollection_Call) Return(_a0 error) *VolumeAttachmentInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *VolumeAttachmentInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, v1.DeleteOptions, v1.ListOptions) error) *VolumeAttachmentInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *VolumeAttachmentInterface) Get(ctx context.Context, name string, opts v1.GetOptions) (*storagev1beta1.VolumeAttachment, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *storagev1beta1.VolumeAttachment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) (*storagev1beta1.VolumeAttachment, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) *storagev1beta1.VolumeAttachment); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1beta1.VolumeAttachment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, v1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VolumeAttachmentInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type VolumeAttachmentInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.GetOptions
func (_e *VolumeAttachmentInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *VolumeAttachmentInterface_Get_Call {
	return &VolumeAttachmentInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *VolumeAttachmentInterface_Get_Call) Run(run func(ctx context.Context, name string, opts v1.GetOptions)) *VolumeAttachmentInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.GetOptions))
	})
	return _c
}

func (_c *VolumeAttachmentInterface_Get_Call) Return(_a0 *storagev1beta1.VolumeAttachment, _a1 error) *VolumeAttachmentInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VolumeAttachmentInterface_Get_Call) RunAndReturn(run func(context.Context, string, v1.GetOptions) (*storagev1beta1.VolumeAttachment, error)) *VolumeAttachmentInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *VolumeAttachmentInterface) List(ctx context.Context, opts v1.ListOptions) (*storagev1beta1.VolumeAttachmentList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *storagev1beta1.VolumeAttachmentList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) (*storagev1beta1.VolumeAttachmentList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) *storagev1beta1.VolumeAttachmentList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1beta1.VolumeAttachmentList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VolumeAttachmentInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type VolumeAttachmentInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *VolumeAttachmentInterface_Expecter) List(ctx interface{}, opts interface{}) *VolumeAttachmentInterface_List_Call {
	return &VolumeAttachmentInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *VolumeAttachmentInterface_List_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *VolumeAttachmentInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *VolumeAttachmentInterface_List_Call) Return(_a0 *storagev1beta1.VolumeAttachmentList, _a1 error) *VolumeAttachmentInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VolumeAttachmentInterface_List_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (*storagev1beta1.VolumeAttachmentList, error)) *VolumeAttachmentInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *VolumeAttachmentInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (*storagev1beta1.VolumeAttachment, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storagev1beta1.VolumeAttachment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*storagev1beta1.VolumeAttachment, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) *storagev1beta1.VolumeAttachment); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1beta1.VolumeAttachment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VolumeAttachmentInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type VolumeAttachmentInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts v1.PatchOptions
//   - subresources ...string
func (_e *VolumeAttachmentInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *VolumeAttachmentInterface_Patch_Call {
	return &VolumeAttachmentInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *VolumeAttachmentInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string)) *VolumeAttachmentInterface_Patch_Call {
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

func (_c *VolumeAttachmentInterface_Patch_Call) Return(result *storagev1beta1.VolumeAttachment, err error) *VolumeAttachmentInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *VolumeAttachmentInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*storagev1beta1.VolumeAttachment, error)) *VolumeAttachmentInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, volumeAttachment, opts
func (_m *VolumeAttachmentInterface) Update(ctx context.Context, volumeAttachment *storagev1beta1.VolumeAttachment, opts v1.UpdateOptions) (*storagev1beta1.VolumeAttachment, error) {
	ret := _m.Called(ctx, volumeAttachment, opts)

	var r0 *storagev1beta1.VolumeAttachment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storagev1beta1.VolumeAttachment, v1.UpdateOptions) (*storagev1beta1.VolumeAttachment, error)); ok {
		return rf(ctx, volumeAttachment, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storagev1beta1.VolumeAttachment, v1.UpdateOptions) *storagev1beta1.VolumeAttachment); ok {
		r0 = rf(ctx, volumeAttachment, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1beta1.VolumeAttachment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storagev1beta1.VolumeAttachment, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, volumeAttachment, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VolumeAttachmentInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type VolumeAttachmentInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - volumeAttachment *storagev1beta1.VolumeAttachment
//   - opts v1.UpdateOptions
func (_e *VolumeAttachmentInterface_Expecter) Update(ctx interface{}, volumeAttachment interface{}, opts interface{}) *VolumeAttachmentInterface_Update_Call {
	return &VolumeAttachmentInterface_Update_Call{Call: _e.mock.On("Update", ctx, volumeAttachment, opts)}
}

func (_c *VolumeAttachmentInterface_Update_Call) Run(run func(ctx context.Context, volumeAttachment *storagev1beta1.VolumeAttachment, opts v1.UpdateOptions)) *VolumeAttachmentInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*storagev1beta1.VolumeAttachment), args[2].(v1.UpdateOptions))
	})
	return _c
}

func (_c *VolumeAttachmentInterface_Update_Call) Return(_a0 *storagev1beta1.VolumeAttachment, _a1 error) *VolumeAttachmentInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VolumeAttachmentInterface_Update_Call) RunAndReturn(run func(context.Context, *storagev1beta1.VolumeAttachment, v1.UpdateOptions) (*storagev1beta1.VolumeAttachment, error)) *VolumeAttachmentInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStatus provides a mock function with given fields: ctx, volumeAttachment, opts
func (_m *VolumeAttachmentInterface) UpdateStatus(ctx context.Context, volumeAttachment *storagev1beta1.VolumeAttachment, opts v1.UpdateOptions) (*storagev1beta1.VolumeAttachment, error) {
	ret := _m.Called(ctx, volumeAttachment, opts)

	var r0 *storagev1beta1.VolumeAttachment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storagev1beta1.VolumeAttachment, v1.UpdateOptions) (*storagev1beta1.VolumeAttachment, error)); ok {
		return rf(ctx, volumeAttachment, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storagev1beta1.VolumeAttachment, v1.UpdateOptions) *storagev1beta1.VolumeAttachment); ok {
		r0 = rf(ctx, volumeAttachment, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1beta1.VolumeAttachment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storagev1beta1.VolumeAttachment, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, volumeAttachment, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VolumeAttachmentInterface_UpdateStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStatus'
type VolumeAttachmentInterface_UpdateStatus_Call struct {
	*mock.Call
}

// UpdateStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - volumeAttachment *storagev1beta1.VolumeAttachment
//   - opts v1.UpdateOptions
func (_e *VolumeAttachmentInterface_Expecter) UpdateStatus(ctx interface{}, volumeAttachment interface{}, opts interface{}) *VolumeAttachmentInterface_UpdateStatus_Call {
	return &VolumeAttachmentInterface_UpdateStatus_Call{Call: _e.mock.On("UpdateStatus", ctx, volumeAttachment, opts)}
}

func (_c *VolumeAttachmentInterface_UpdateStatus_Call) Run(run func(ctx context.Context, volumeAttachment *storagev1beta1.VolumeAttachment, opts v1.UpdateOptions)) *VolumeAttachmentInterface_UpdateStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*storagev1beta1.VolumeAttachment), args[2].(v1.UpdateOptions))
	})
	return _c
}

func (_c *VolumeAttachmentInterface_UpdateStatus_Call) Return(_a0 *storagev1beta1.VolumeAttachment, _a1 error) *VolumeAttachmentInterface_UpdateStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VolumeAttachmentInterface_UpdateStatus_Call) RunAndReturn(run func(context.Context, *storagev1beta1.VolumeAttachment, v1.UpdateOptions) (*storagev1beta1.VolumeAttachment, error)) *VolumeAttachmentInterface_UpdateStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *VolumeAttachmentInterface) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
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

// VolumeAttachmentInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type VolumeAttachmentInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *VolumeAttachmentInterface_Expecter) Watch(ctx interface{}, opts interface{}) *VolumeAttachmentInterface_Watch_Call {
	return &VolumeAttachmentInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *VolumeAttachmentInterface_Watch_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *VolumeAttachmentInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *VolumeAttachmentInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *VolumeAttachmentInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VolumeAttachmentInterface_Watch_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (watch.Interface, error)) *VolumeAttachmentInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewVolumeAttachmentInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewVolumeAttachmentInterface creates a new instance of VolumeAttachmentInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewVolumeAttachmentInterface(t mockConstructorTestingTNewVolumeAttachmentInterface) *VolumeAttachmentInterface {
	mock := &VolumeAttachmentInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
