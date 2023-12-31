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

// LimitRangeInterface is an autogenerated mock type for the LimitRangeInterface type
type LimitRangeInterface struct {
	mock.Mock
}

type LimitRangeInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *LimitRangeInterface) EXPECT() *LimitRangeInterface_Expecter {
	return &LimitRangeInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, limitRange, opts
func (_m *LimitRangeInterface) Apply(ctx context.Context, limitRange *v1.LimitRangeApplyConfiguration, opts metav1.ApplyOptions) (*corev1.LimitRange, error) {
	ret := _m.Called(ctx, limitRange, opts)

	var r0 *corev1.LimitRange
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.LimitRangeApplyConfiguration, metav1.ApplyOptions) (*corev1.LimitRange, error)); ok {
		return rf(ctx, limitRange, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.LimitRangeApplyConfiguration, metav1.ApplyOptions) *corev1.LimitRange); ok {
		r0 = rf(ctx, limitRange, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.LimitRange)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.LimitRangeApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, limitRange, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LimitRangeInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type LimitRangeInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - limitRange *v1.LimitRangeApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *LimitRangeInterface_Expecter) Apply(ctx interface{}, limitRange interface{}, opts interface{}) *LimitRangeInterface_Apply_Call {
	return &LimitRangeInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, limitRange, opts)}
}

func (_c *LimitRangeInterface_Apply_Call) Run(run func(ctx context.Context, limitRange *v1.LimitRangeApplyConfiguration, opts metav1.ApplyOptions)) *LimitRangeInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.LimitRangeApplyConfiguration), args[2].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *LimitRangeInterface_Apply_Call) Return(result *corev1.LimitRange, err error) *LimitRangeInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *LimitRangeInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1.LimitRangeApplyConfiguration, metav1.ApplyOptions) (*corev1.LimitRange, error)) *LimitRangeInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, limitRange, opts
func (_m *LimitRangeInterface) Create(ctx context.Context, limitRange *corev1.LimitRange, opts metav1.CreateOptions) (*corev1.LimitRange, error) {
	ret := _m.Called(ctx, limitRange, opts)

	var r0 *corev1.LimitRange
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.LimitRange, metav1.CreateOptions) (*corev1.LimitRange, error)); ok {
		return rf(ctx, limitRange, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.LimitRange, metav1.CreateOptions) *corev1.LimitRange); ok {
		r0 = rf(ctx, limitRange, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.LimitRange)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *corev1.LimitRange, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, limitRange, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LimitRangeInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type LimitRangeInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - limitRange *corev1.LimitRange
//   - opts metav1.CreateOptions
func (_e *LimitRangeInterface_Expecter) Create(ctx interface{}, limitRange interface{}, opts interface{}) *LimitRangeInterface_Create_Call {
	return &LimitRangeInterface_Create_Call{Call: _e.mock.On("Create", ctx, limitRange, opts)}
}

func (_c *LimitRangeInterface_Create_Call) Run(run func(ctx context.Context, limitRange *corev1.LimitRange, opts metav1.CreateOptions)) *LimitRangeInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.LimitRange), args[2].(metav1.CreateOptions))
	})
	return _c
}

func (_c *LimitRangeInterface_Create_Call) Return(_a0 *corev1.LimitRange, _a1 error) *LimitRangeInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LimitRangeInterface_Create_Call) RunAndReturn(run func(context.Context, *corev1.LimitRange, metav1.CreateOptions) (*corev1.LimitRange, error)) *LimitRangeInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *LimitRangeInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LimitRangeInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type LimitRangeInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.DeleteOptions
func (_e *LimitRangeInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *LimitRangeInterface_Delete_Call {
	return &LimitRangeInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *LimitRangeInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts metav1.DeleteOptions)) *LimitRangeInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.DeleteOptions))
	})
	return _c
}

func (_c *LimitRangeInterface_Delete_Call) Return(_a0 error) *LimitRangeInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LimitRangeInterface_Delete_Call) RunAndReturn(run func(context.Context, string, metav1.DeleteOptions) error) *LimitRangeInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *LimitRangeInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LimitRangeInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type LimitRangeInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.DeleteOptions
//   - listOpts metav1.ListOptions
func (_e *LimitRangeInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *LimitRangeInterface_DeleteCollection_Call {
	return &LimitRangeInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *LimitRangeInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions)) *LimitRangeInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.DeleteOptions), args[2].(metav1.ListOptions))
	})
	return _c
}

func (_c *LimitRangeInterface_DeleteCollection_Call) Return(_a0 error) *LimitRangeInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LimitRangeInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error) *LimitRangeInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *LimitRangeInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.LimitRange, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *corev1.LimitRange
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) (*corev1.LimitRange, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *corev1.LimitRange); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.LimitRange)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LimitRangeInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type LimitRangeInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.GetOptions
func (_e *LimitRangeInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *LimitRangeInterface_Get_Call {
	return &LimitRangeInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *LimitRangeInterface_Get_Call) Run(run func(ctx context.Context, name string, opts metav1.GetOptions)) *LimitRangeInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.GetOptions))
	})
	return _c
}

func (_c *LimitRangeInterface_Get_Call) Return(_a0 *corev1.LimitRange, _a1 error) *LimitRangeInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LimitRangeInterface_Get_Call) RunAndReturn(run func(context.Context, string, metav1.GetOptions) (*corev1.LimitRange, error)) *LimitRangeInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *LimitRangeInterface) List(ctx context.Context, opts metav1.ListOptions) (*corev1.LimitRangeList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *corev1.LimitRangeList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) (*corev1.LimitRangeList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *corev1.LimitRangeList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.LimitRangeList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LimitRangeInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type LimitRangeInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *LimitRangeInterface_Expecter) List(ctx interface{}, opts interface{}) *LimitRangeInterface_List_Call {
	return &LimitRangeInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *LimitRangeInterface_List_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *LimitRangeInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *LimitRangeInterface_List_Call) Return(_a0 *corev1.LimitRangeList, _a1 error) *LimitRangeInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LimitRangeInterface_List_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (*corev1.LimitRangeList, error)) *LimitRangeInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *LimitRangeInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1.LimitRange, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *corev1.LimitRange
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*corev1.LimitRange, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *corev1.LimitRange); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.LimitRange)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LimitRangeInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type LimitRangeInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts metav1.PatchOptions
//   - subresources ...string
func (_e *LimitRangeInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *LimitRangeInterface_Patch_Call {
	return &LimitRangeInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *LimitRangeInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string)) *LimitRangeInterface_Patch_Call {
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

func (_c *LimitRangeInterface_Patch_Call) Return(result *corev1.LimitRange, err error) *LimitRangeInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *LimitRangeInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*corev1.LimitRange, error)) *LimitRangeInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, limitRange, opts
func (_m *LimitRangeInterface) Update(ctx context.Context, limitRange *corev1.LimitRange, opts metav1.UpdateOptions) (*corev1.LimitRange, error) {
	ret := _m.Called(ctx, limitRange, opts)

	var r0 *corev1.LimitRange
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.LimitRange, metav1.UpdateOptions) (*corev1.LimitRange, error)); ok {
		return rf(ctx, limitRange, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.LimitRange, metav1.UpdateOptions) *corev1.LimitRange); ok {
		r0 = rf(ctx, limitRange, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.LimitRange)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *corev1.LimitRange, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, limitRange, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LimitRangeInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type LimitRangeInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - limitRange *corev1.LimitRange
//   - opts metav1.UpdateOptions
func (_e *LimitRangeInterface_Expecter) Update(ctx interface{}, limitRange interface{}, opts interface{}) *LimitRangeInterface_Update_Call {
	return &LimitRangeInterface_Update_Call{Call: _e.mock.On("Update", ctx, limitRange, opts)}
}

func (_c *LimitRangeInterface_Update_Call) Run(run func(ctx context.Context, limitRange *corev1.LimitRange, opts metav1.UpdateOptions)) *LimitRangeInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.LimitRange), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *LimitRangeInterface_Update_Call) Return(_a0 *corev1.LimitRange, _a1 error) *LimitRangeInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LimitRangeInterface_Update_Call) RunAndReturn(run func(context.Context, *corev1.LimitRange, metav1.UpdateOptions) (*corev1.LimitRange, error)) *LimitRangeInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *LimitRangeInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
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

// LimitRangeInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type LimitRangeInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *LimitRangeInterface_Expecter) Watch(ctx interface{}, opts interface{}) *LimitRangeInterface_Watch_Call {
	return &LimitRangeInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *LimitRangeInterface_Watch_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *LimitRangeInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *LimitRangeInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *LimitRangeInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LimitRangeInterface_Watch_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (watch.Interface, error)) *LimitRangeInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewLimitRangeInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewLimitRangeInterface creates a new instance of LimitRangeInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLimitRangeInterface(t mockConstructorTestingTNewLimitRangeInterface) *LimitRangeInterface {
	mock := &LimitRangeInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
