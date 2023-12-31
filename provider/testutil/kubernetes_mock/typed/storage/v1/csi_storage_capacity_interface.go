// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	storagev1 "k8s.io/api/storage/v1"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/storage/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// CSIStorageCapacityInterface is an autogenerated mock type for the CSIStorageCapacityInterface type
type CSIStorageCapacityInterface struct {
	mock.Mock
}

type CSIStorageCapacityInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *CSIStorageCapacityInterface) EXPECT() *CSIStorageCapacityInterface_Expecter {
	return &CSIStorageCapacityInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, cSIStorageCapacity, opts
func (_m *CSIStorageCapacityInterface) Apply(ctx context.Context, cSIStorageCapacity *v1.CSIStorageCapacityApplyConfiguration, opts metav1.ApplyOptions) (*storagev1.CSIStorageCapacity, error) {
	ret := _m.Called(ctx, cSIStorageCapacity, opts)

	var r0 *storagev1.CSIStorageCapacity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CSIStorageCapacityApplyConfiguration, metav1.ApplyOptions) (*storagev1.CSIStorageCapacity, error)); ok {
		return rf(ctx, cSIStorageCapacity, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CSIStorageCapacityApplyConfiguration, metav1.ApplyOptions) *storagev1.CSIStorageCapacity); ok {
		r0 = rf(ctx, cSIStorageCapacity, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1.CSIStorageCapacity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.CSIStorageCapacityApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, cSIStorageCapacity, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSIStorageCapacityInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type CSIStorageCapacityInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - cSIStorageCapacity *v1.CSIStorageCapacityApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *CSIStorageCapacityInterface_Expecter) Apply(ctx interface{}, cSIStorageCapacity interface{}, opts interface{}) *CSIStorageCapacityInterface_Apply_Call {
	return &CSIStorageCapacityInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, cSIStorageCapacity, opts)}
}

func (_c *CSIStorageCapacityInterface_Apply_Call) Run(run func(ctx context.Context, cSIStorageCapacity *v1.CSIStorageCapacityApplyConfiguration, opts metav1.ApplyOptions)) *CSIStorageCapacityInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.CSIStorageCapacityApplyConfiguration), args[2].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *CSIStorageCapacityInterface_Apply_Call) Return(result *storagev1.CSIStorageCapacity, err error) *CSIStorageCapacityInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *CSIStorageCapacityInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1.CSIStorageCapacityApplyConfiguration, metav1.ApplyOptions) (*storagev1.CSIStorageCapacity, error)) *CSIStorageCapacityInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, cSIStorageCapacity, opts
func (_m *CSIStorageCapacityInterface) Create(ctx context.Context, cSIStorageCapacity *storagev1.CSIStorageCapacity, opts metav1.CreateOptions) (*storagev1.CSIStorageCapacity, error) {
	ret := _m.Called(ctx, cSIStorageCapacity, opts)

	var r0 *storagev1.CSIStorageCapacity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storagev1.CSIStorageCapacity, metav1.CreateOptions) (*storagev1.CSIStorageCapacity, error)); ok {
		return rf(ctx, cSIStorageCapacity, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storagev1.CSIStorageCapacity, metav1.CreateOptions) *storagev1.CSIStorageCapacity); ok {
		r0 = rf(ctx, cSIStorageCapacity, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1.CSIStorageCapacity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storagev1.CSIStorageCapacity, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, cSIStorageCapacity, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSIStorageCapacityInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type CSIStorageCapacityInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - cSIStorageCapacity *storagev1.CSIStorageCapacity
//   - opts metav1.CreateOptions
func (_e *CSIStorageCapacityInterface_Expecter) Create(ctx interface{}, cSIStorageCapacity interface{}, opts interface{}) *CSIStorageCapacityInterface_Create_Call {
	return &CSIStorageCapacityInterface_Create_Call{Call: _e.mock.On("Create", ctx, cSIStorageCapacity, opts)}
}

func (_c *CSIStorageCapacityInterface_Create_Call) Run(run func(ctx context.Context, cSIStorageCapacity *storagev1.CSIStorageCapacity, opts metav1.CreateOptions)) *CSIStorageCapacityInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*storagev1.CSIStorageCapacity), args[2].(metav1.CreateOptions))
	})
	return _c
}

func (_c *CSIStorageCapacityInterface_Create_Call) Return(_a0 *storagev1.CSIStorageCapacity, _a1 error) *CSIStorageCapacityInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CSIStorageCapacityInterface_Create_Call) RunAndReturn(run func(context.Context, *storagev1.CSIStorageCapacity, metav1.CreateOptions) (*storagev1.CSIStorageCapacity, error)) *CSIStorageCapacityInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *CSIStorageCapacityInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CSIStorageCapacityInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type CSIStorageCapacityInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.DeleteOptions
func (_e *CSIStorageCapacityInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *CSIStorageCapacityInterface_Delete_Call {
	return &CSIStorageCapacityInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *CSIStorageCapacityInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts metav1.DeleteOptions)) *CSIStorageCapacityInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.DeleteOptions))
	})
	return _c
}

func (_c *CSIStorageCapacityInterface_Delete_Call) Return(_a0 error) *CSIStorageCapacityInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CSIStorageCapacityInterface_Delete_Call) RunAndReturn(run func(context.Context, string, metav1.DeleteOptions) error) *CSIStorageCapacityInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *CSIStorageCapacityInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CSIStorageCapacityInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type CSIStorageCapacityInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.DeleteOptions
//   - listOpts metav1.ListOptions
func (_e *CSIStorageCapacityInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *CSIStorageCapacityInterface_DeleteCollection_Call {
	return &CSIStorageCapacityInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *CSIStorageCapacityInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions)) *CSIStorageCapacityInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.DeleteOptions), args[2].(metav1.ListOptions))
	})
	return _c
}

func (_c *CSIStorageCapacityInterface_DeleteCollection_Call) Return(_a0 error) *CSIStorageCapacityInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CSIStorageCapacityInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error) *CSIStorageCapacityInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *CSIStorageCapacityInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*storagev1.CSIStorageCapacity, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *storagev1.CSIStorageCapacity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) (*storagev1.CSIStorageCapacity, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *storagev1.CSIStorageCapacity); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1.CSIStorageCapacity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSIStorageCapacityInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type CSIStorageCapacityInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.GetOptions
func (_e *CSIStorageCapacityInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *CSIStorageCapacityInterface_Get_Call {
	return &CSIStorageCapacityInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *CSIStorageCapacityInterface_Get_Call) Run(run func(ctx context.Context, name string, opts metav1.GetOptions)) *CSIStorageCapacityInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.GetOptions))
	})
	return _c
}

func (_c *CSIStorageCapacityInterface_Get_Call) Return(_a0 *storagev1.CSIStorageCapacity, _a1 error) *CSIStorageCapacityInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CSIStorageCapacityInterface_Get_Call) RunAndReturn(run func(context.Context, string, metav1.GetOptions) (*storagev1.CSIStorageCapacity, error)) *CSIStorageCapacityInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *CSIStorageCapacityInterface) List(ctx context.Context, opts metav1.ListOptions) (*storagev1.CSIStorageCapacityList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *storagev1.CSIStorageCapacityList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) (*storagev1.CSIStorageCapacityList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *storagev1.CSIStorageCapacityList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1.CSIStorageCapacityList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSIStorageCapacityInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type CSIStorageCapacityInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *CSIStorageCapacityInterface_Expecter) List(ctx interface{}, opts interface{}) *CSIStorageCapacityInterface_List_Call {
	return &CSIStorageCapacityInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *CSIStorageCapacityInterface_List_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *CSIStorageCapacityInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *CSIStorageCapacityInterface_List_Call) Return(_a0 *storagev1.CSIStorageCapacityList, _a1 error) *CSIStorageCapacityInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CSIStorageCapacityInterface_List_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (*storagev1.CSIStorageCapacityList, error)) *CSIStorageCapacityInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *CSIStorageCapacityInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*storagev1.CSIStorageCapacity, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storagev1.CSIStorageCapacity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*storagev1.CSIStorageCapacity, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *storagev1.CSIStorageCapacity); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1.CSIStorageCapacity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSIStorageCapacityInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type CSIStorageCapacityInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts metav1.PatchOptions
//   - subresources ...string
func (_e *CSIStorageCapacityInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *CSIStorageCapacityInterface_Patch_Call {
	return &CSIStorageCapacityInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *CSIStorageCapacityInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string)) *CSIStorageCapacityInterface_Patch_Call {
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

func (_c *CSIStorageCapacityInterface_Patch_Call) Return(result *storagev1.CSIStorageCapacity, err error) *CSIStorageCapacityInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *CSIStorageCapacityInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*storagev1.CSIStorageCapacity, error)) *CSIStorageCapacityInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, cSIStorageCapacity, opts
func (_m *CSIStorageCapacityInterface) Update(ctx context.Context, cSIStorageCapacity *storagev1.CSIStorageCapacity, opts metav1.UpdateOptions) (*storagev1.CSIStorageCapacity, error) {
	ret := _m.Called(ctx, cSIStorageCapacity, opts)

	var r0 *storagev1.CSIStorageCapacity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storagev1.CSIStorageCapacity, metav1.UpdateOptions) (*storagev1.CSIStorageCapacity, error)); ok {
		return rf(ctx, cSIStorageCapacity, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storagev1.CSIStorageCapacity, metav1.UpdateOptions) *storagev1.CSIStorageCapacity); ok {
		r0 = rf(ctx, cSIStorageCapacity, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1.CSIStorageCapacity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storagev1.CSIStorageCapacity, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, cSIStorageCapacity, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSIStorageCapacityInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type CSIStorageCapacityInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - cSIStorageCapacity *storagev1.CSIStorageCapacity
//   - opts metav1.UpdateOptions
func (_e *CSIStorageCapacityInterface_Expecter) Update(ctx interface{}, cSIStorageCapacity interface{}, opts interface{}) *CSIStorageCapacityInterface_Update_Call {
	return &CSIStorageCapacityInterface_Update_Call{Call: _e.mock.On("Update", ctx, cSIStorageCapacity, opts)}
}

func (_c *CSIStorageCapacityInterface_Update_Call) Run(run func(ctx context.Context, cSIStorageCapacity *storagev1.CSIStorageCapacity, opts metav1.UpdateOptions)) *CSIStorageCapacityInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*storagev1.CSIStorageCapacity), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *CSIStorageCapacityInterface_Update_Call) Return(_a0 *storagev1.CSIStorageCapacity, _a1 error) *CSIStorageCapacityInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CSIStorageCapacityInterface_Update_Call) RunAndReturn(run func(context.Context, *storagev1.CSIStorageCapacity, metav1.UpdateOptions) (*storagev1.CSIStorageCapacity, error)) *CSIStorageCapacityInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *CSIStorageCapacityInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
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

// CSIStorageCapacityInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type CSIStorageCapacityInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *CSIStorageCapacityInterface_Expecter) Watch(ctx interface{}, opts interface{}) *CSIStorageCapacityInterface_Watch_Call {
	return &CSIStorageCapacityInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *CSIStorageCapacityInterface_Watch_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *CSIStorageCapacityInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *CSIStorageCapacityInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *CSIStorageCapacityInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CSIStorageCapacityInterface_Watch_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (watch.Interface, error)) *CSIStorageCapacityInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCSIStorageCapacityInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewCSIStorageCapacityInterface creates a new instance of CSIStorageCapacityInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCSIStorageCapacityInterface(t mockConstructorTestingTNewCSIStorageCapacityInterface) *CSIStorageCapacityInterface {
	mock := &CSIStorageCapacityInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
