// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	discoveryv1 "k8s.io/api/discovery/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/discovery/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// EndpointSliceInterface is an autogenerated mock type for the EndpointSliceInterface type
type EndpointSliceInterface struct {
	mock.Mock
}

type EndpointSliceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *EndpointSliceInterface) EXPECT() *EndpointSliceInterface_Expecter {
	return &EndpointSliceInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, endpointSlice, opts
func (_m *EndpointSliceInterface) Apply(ctx context.Context, endpointSlice *v1.EndpointSliceApplyConfiguration, opts metav1.ApplyOptions) (*discoveryv1.EndpointSlice, error) {
	ret := _m.Called(ctx, endpointSlice, opts)

	var r0 *discoveryv1.EndpointSlice
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.EndpointSliceApplyConfiguration, metav1.ApplyOptions) (*discoveryv1.EndpointSlice, error)); ok {
		return rf(ctx, endpointSlice, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.EndpointSliceApplyConfiguration, metav1.ApplyOptions) *discoveryv1.EndpointSlice); ok {
		r0 = rf(ctx, endpointSlice, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discoveryv1.EndpointSlice)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.EndpointSliceApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, endpointSlice, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EndpointSliceInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type EndpointSliceInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - endpointSlice *v1.EndpointSliceApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *EndpointSliceInterface_Expecter) Apply(ctx interface{}, endpointSlice interface{}, opts interface{}) *EndpointSliceInterface_Apply_Call {
	return &EndpointSliceInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, endpointSlice, opts)}
}

func (_c *EndpointSliceInterface_Apply_Call) Run(run func(ctx context.Context, endpointSlice *v1.EndpointSliceApplyConfiguration, opts metav1.ApplyOptions)) *EndpointSliceInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.EndpointSliceApplyConfiguration), args[2].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *EndpointSliceInterface_Apply_Call) Return(result *discoveryv1.EndpointSlice, err error) *EndpointSliceInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *EndpointSliceInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1.EndpointSliceApplyConfiguration, metav1.ApplyOptions) (*discoveryv1.EndpointSlice, error)) *EndpointSliceInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, endpointSlice, opts
func (_m *EndpointSliceInterface) Create(ctx context.Context, endpointSlice *discoveryv1.EndpointSlice, opts metav1.CreateOptions) (*discoveryv1.EndpointSlice, error) {
	ret := _m.Called(ctx, endpointSlice, opts)

	var r0 *discoveryv1.EndpointSlice
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *discoveryv1.EndpointSlice, metav1.CreateOptions) (*discoveryv1.EndpointSlice, error)); ok {
		return rf(ctx, endpointSlice, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *discoveryv1.EndpointSlice, metav1.CreateOptions) *discoveryv1.EndpointSlice); ok {
		r0 = rf(ctx, endpointSlice, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discoveryv1.EndpointSlice)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *discoveryv1.EndpointSlice, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, endpointSlice, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EndpointSliceInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type EndpointSliceInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - endpointSlice *discoveryv1.EndpointSlice
//   - opts metav1.CreateOptions
func (_e *EndpointSliceInterface_Expecter) Create(ctx interface{}, endpointSlice interface{}, opts interface{}) *EndpointSliceInterface_Create_Call {
	return &EndpointSliceInterface_Create_Call{Call: _e.mock.On("Create", ctx, endpointSlice, opts)}
}

func (_c *EndpointSliceInterface_Create_Call) Run(run func(ctx context.Context, endpointSlice *discoveryv1.EndpointSlice, opts metav1.CreateOptions)) *EndpointSliceInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*discoveryv1.EndpointSlice), args[2].(metav1.CreateOptions))
	})
	return _c
}

func (_c *EndpointSliceInterface_Create_Call) Return(_a0 *discoveryv1.EndpointSlice, _a1 error) *EndpointSliceInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EndpointSliceInterface_Create_Call) RunAndReturn(run func(context.Context, *discoveryv1.EndpointSlice, metav1.CreateOptions) (*discoveryv1.EndpointSlice, error)) *EndpointSliceInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *EndpointSliceInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EndpointSliceInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type EndpointSliceInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.DeleteOptions
func (_e *EndpointSliceInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *EndpointSliceInterface_Delete_Call {
	return &EndpointSliceInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *EndpointSliceInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts metav1.DeleteOptions)) *EndpointSliceInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.DeleteOptions))
	})
	return _c
}

func (_c *EndpointSliceInterface_Delete_Call) Return(_a0 error) *EndpointSliceInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EndpointSliceInterface_Delete_Call) RunAndReturn(run func(context.Context, string, metav1.DeleteOptions) error) *EndpointSliceInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *EndpointSliceInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EndpointSliceInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type EndpointSliceInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.DeleteOptions
//   - listOpts metav1.ListOptions
func (_e *EndpointSliceInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *EndpointSliceInterface_DeleteCollection_Call {
	return &EndpointSliceInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *EndpointSliceInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions)) *EndpointSliceInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.DeleteOptions), args[2].(metav1.ListOptions))
	})
	return _c
}

func (_c *EndpointSliceInterface_DeleteCollection_Call) Return(_a0 error) *EndpointSliceInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EndpointSliceInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error) *EndpointSliceInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *EndpointSliceInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*discoveryv1.EndpointSlice, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *discoveryv1.EndpointSlice
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) (*discoveryv1.EndpointSlice, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *discoveryv1.EndpointSlice); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discoveryv1.EndpointSlice)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EndpointSliceInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type EndpointSliceInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.GetOptions
func (_e *EndpointSliceInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *EndpointSliceInterface_Get_Call {
	return &EndpointSliceInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *EndpointSliceInterface_Get_Call) Run(run func(ctx context.Context, name string, opts metav1.GetOptions)) *EndpointSliceInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.GetOptions))
	})
	return _c
}

func (_c *EndpointSliceInterface_Get_Call) Return(_a0 *discoveryv1.EndpointSlice, _a1 error) *EndpointSliceInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EndpointSliceInterface_Get_Call) RunAndReturn(run func(context.Context, string, metav1.GetOptions) (*discoveryv1.EndpointSlice, error)) *EndpointSliceInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *EndpointSliceInterface) List(ctx context.Context, opts metav1.ListOptions) (*discoveryv1.EndpointSliceList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *discoveryv1.EndpointSliceList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) (*discoveryv1.EndpointSliceList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *discoveryv1.EndpointSliceList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discoveryv1.EndpointSliceList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EndpointSliceInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type EndpointSliceInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *EndpointSliceInterface_Expecter) List(ctx interface{}, opts interface{}) *EndpointSliceInterface_List_Call {
	return &EndpointSliceInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *EndpointSliceInterface_List_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *EndpointSliceInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *EndpointSliceInterface_List_Call) Return(_a0 *discoveryv1.EndpointSliceList, _a1 error) *EndpointSliceInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EndpointSliceInterface_List_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (*discoveryv1.EndpointSliceList, error)) *EndpointSliceInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *EndpointSliceInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*discoveryv1.EndpointSlice, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *discoveryv1.EndpointSlice
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*discoveryv1.EndpointSlice, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *discoveryv1.EndpointSlice); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discoveryv1.EndpointSlice)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EndpointSliceInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type EndpointSliceInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts metav1.PatchOptions
//   - subresources ...string
func (_e *EndpointSliceInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *EndpointSliceInterface_Patch_Call {
	return &EndpointSliceInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *EndpointSliceInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string)) *EndpointSliceInterface_Patch_Call {
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

func (_c *EndpointSliceInterface_Patch_Call) Return(result *discoveryv1.EndpointSlice, err error) *EndpointSliceInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *EndpointSliceInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*discoveryv1.EndpointSlice, error)) *EndpointSliceInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, endpointSlice, opts
func (_m *EndpointSliceInterface) Update(ctx context.Context, endpointSlice *discoveryv1.EndpointSlice, opts metav1.UpdateOptions) (*discoveryv1.EndpointSlice, error) {
	ret := _m.Called(ctx, endpointSlice, opts)

	var r0 *discoveryv1.EndpointSlice
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *discoveryv1.EndpointSlice, metav1.UpdateOptions) (*discoveryv1.EndpointSlice, error)); ok {
		return rf(ctx, endpointSlice, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *discoveryv1.EndpointSlice, metav1.UpdateOptions) *discoveryv1.EndpointSlice); ok {
		r0 = rf(ctx, endpointSlice, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discoveryv1.EndpointSlice)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *discoveryv1.EndpointSlice, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, endpointSlice, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EndpointSliceInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type EndpointSliceInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - endpointSlice *discoveryv1.EndpointSlice
//   - opts metav1.UpdateOptions
func (_e *EndpointSliceInterface_Expecter) Update(ctx interface{}, endpointSlice interface{}, opts interface{}) *EndpointSliceInterface_Update_Call {
	return &EndpointSliceInterface_Update_Call{Call: _e.mock.On("Update", ctx, endpointSlice, opts)}
}

func (_c *EndpointSliceInterface_Update_Call) Run(run func(ctx context.Context, endpointSlice *discoveryv1.EndpointSlice, opts metav1.UpdateOptions)) *EndpointSliceInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*discoveryv1.EndpointSlice), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *EndpointSliceInterface_Update_Call) Return(_a0 *discoveryv1.EndpointSlice, _a1 error) *EndpointSliceInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EndpointSliceInterface_Update_Call) RunAndReturn(run func(context.Context, *discoveryv1.EndpointSlice, metav1.UpdateOptions) (*discoveryv1.EndpointSlice, error)) *EndpointSliceInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *EndpointSliceInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
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

// EndpointSliceInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type EndpointSliceInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *EndpointSliceInterface_Expecter) Watch(ctx interface{}, opts interface{}) *EndpointSliceInterface_Watch_Call {
	return &EndpointSliceInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *EndpointSliceInterface_Watch_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *EndpointSliceInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *EndpointSliceInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *EndpointSliceInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EndpointSliceInterface_Watch_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (watch.Interface, error)) *EndpointSliceInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewEndpointSliceInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewEndpointSliceInterface creates a new instance of EndpointSliceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEndpointSliceInterface(t mockConstructorTestingTNewEndpointSliceInterface) *EndpointSliceInterface {
	mock := &EndpointSliceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}