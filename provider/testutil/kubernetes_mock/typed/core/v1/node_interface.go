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

// NodeInterface is an autogenerated mock type for the NodeInterface type
type NodeInterface struct {
	mock.Mock
}

type NodeInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *NodeInterface) EXPECT() *NodeInterface_Expecter {
	return &NodeInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, node, opts
func (_m *NodeInterface) Apply(ctx context.Context, node *v1.NodeApplyConfiguration, opts metav1.ApplyOptions) (*corev1.Node, error) {
	ret := _m.Called(ctx, node, opts)

	var r0 *corev1.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.NodeApplyConfiguration, metav1.ApplyOptions) (*corev1.Node, error)); ok {
		return rf(ctx, node, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.NodeApplyConfiguration, metav1.ApplyOptions) *corev1.Node); ok {
		r0 = rf(ctx, node, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.NodeApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, node, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type NodeInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - node *v1.NodeApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *NodeInterface_Expecter) Apply(ctx interface{}, node interface{}, opts interface{}) *NodeInterface_Apply_Call {
	return &NodeInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, node, opts)}
}

func (_c *NodeInterface_Apply_Call) Run(run func(ctx context.Context, node *v1.NodeApplyConfiguration, opts metav1.ApplyOptions)) *NodeInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.NodeApplyConfiguration), args[2].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *NodeInterface_Apply_Call) Return(result *corev1.Node, err error) *NodeInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *NodeInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1.NodeApplyConfiguration, metav1.ApplyOptions) (*corev1.Node, error)) *NodeInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// ApplyStatus provides a mock function with given fields: ctx, node, opts
func (_m *NodeInterface) ApplyStatus(ctx context.Context, node *v1.NodeApplyConfiguration, opts metav1.ApplyOptions) (*corev1.Node, error) {
	ret := _m.Called(ctx, node, opts)

	var r0 *corev1.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.NodeApplyConfiguration, metav1.ApplyOptions) (*corev1.Node, error)); ok {
		return rf(ctx, node, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.NodeApplyConfiguration, metav1.ApplyOptions) *corev1.Node); ok {
		r0 = rf(ctx, node, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.NodeApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, node, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeInterface_ApplyStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyStatus'
type NodeInterface_ApplyStatus_Call struct {
	*mock.Call
}

// ApplyStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - node *v1.NodeApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *NodeInterface_Expecter) ApplyStatus(ctx interface{}, node interface{}, opts interface{}) *NodeInterface_ApplyStatus_Call {
	return &NodeInterface_ApplyStatus_Call{Call: _e.mock.On("ApplyStatus", ctx, node, opts)}
}

func (_c *NodeInterface_ApplyStatus_Call) Run(run func(ctx context.Context, node *v1.NodeApplyConfiguration, opts metav1.ApplyOptions)) *NodeInterface_ApplyStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.NodeApplyConfiguration), args[2].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *NodeInterface_ApplyStatus_Call) Return(result *corev1.Node, err error) *NodeInterface_ApplyStatus_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *NodeInterface_ApplyStatus_Call) RunAndReturn(run func(context.Context, *v1.NodeApplyConfiguration, metav1.ApplyOptions) (*corev1.Node, error)) *NodeInterface_ApplyStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, node, opts
func (_m *NodeInterface) Create(ctx context.Context, node *corev1.Node, opts metav1.CreateOptions) (*corev1.Node, error) {
	ret := _m.Called(ctx, node, opts)

	var r0 *corev1.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Node, metav1.CreateOptions) (*corev1.Node, error)); ok {
		return rf(ctx, node, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Node, metav1.CreateOptions) *corev1.Node); ok {
		r0 = rf(ctx, node, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *corev1.Node, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, node, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type NodeInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - node *corev1.Node
//   - opts metav1.CreateOptions
func (_e *NodeInterface_Expecter) Create(ctx interface{}, node interface{}, opts interface{}) *NodeInterface_Create_Call {
	return &NodeInterface_Create_Call{Call: _e.mock.On("Create", ctx, node, opts)}
}

func (_c *NodeInterface_Create_Call) Run(run func(ctx context.Context, node *corev1.Node, opts metav1.CreateOptions)) *NodeInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.Node), args[2].(metav1.CreateOptions))
	})
	return _c
}

func (_c *NodeInterface_Create_Call) Return(_a0 *corev1.Node, _a1 error) *NodeInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NodeInterface_Create_Call) RunAndReturn(run func(context.Context, *corev1.Node, metav1.CreateOptions) (*corev1.Node, error)) *NodeInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *NodeInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NodeInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type NodeInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.DeleteOptions
func (_e *NodeInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *NodeInterface_Delete_Call {
	return &NodeInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *NodeInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts metav1.DeleteOptions)) *NodeInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.DeleteOptions))
	})
	return _c
}

func (_c *NodeInterface_Delete_Call) Return(_a0 error) *NodeInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NodeInterface_Delete_Call) RunAndReturn(run func(context.Context, string, metav1.DeleteOptions) error) *NodeInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *NodeInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NodeInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type NodeInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.DeleteOptions
//   - listOpts metav1.ListOptions
func (_e *NodeInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *NodeInterface_DeleteCollection_Call {
	return &NodeInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *NodeInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions)) *NodeInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.DeleteOptions), args[2].(metav1.ListOptions))
	})
	return _c
}

func (_c *NodeInterface_DeleteCollection_Call) Return(_a0 error) *NodeInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NodeInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error) *NodeInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *NodeInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.Node, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *corev1.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) (*corev1.Node, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *corev1.Node); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type NodeInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.GetOptions
func (_e *NodeInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *NodeInterface_Get_Call {
	return &NodeInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *NodeInterface_Get_Call) Run(run func(ctx context.Context, name string, opts metav1.GetOptions)) *NodeInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.GetOptions))
	})
	return _c
}

func (_c *NodeInterface_Get_Call) Return(_a0 *corev1.Node, _a1 error) *NodeInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NodeInterface_Get_Call) RunAndReturn(run func(context.Context, string, metav1.GetOptions) (*corev1.Node, error)) *NodeInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *NodeInterface) List(ctx context.Context, opts metav1.ListOptions) (*corev1.NodeList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *corev1.NodeList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) (*corev1.NodeList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *corev1.NodeList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.NodeList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type NodeInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *NodeInterface_Expecter) List(ctx interface{}, opts interface{}) *NodeInterface_List_Call {
	return &NodeInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *NodeInterface_List_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *NodeInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *NodeInterface_List_Call) Return(_a0 *corev1.NodeList, _a1 error) *NodeInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NodeInterface_List_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (*corev1.NodeList, error)) *NodeInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *NodeInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1.Node, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *corev1.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*corev1.Node, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *corev1.Node); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type NodeInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts metav1.PatchOptions
//   - subresources ...string
func (_e *NodeInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *NodeInterface_Patch_Call {
	return &NodeInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *NodeInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string)) *NodeInterface_Patch_Call {
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

func (_c *NodeInterface_Patch_Call) Return(result *corev1.Node, err error) *NodeInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *NodeInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*corev1.Node, error)) *NodeInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// PatchStatus provides a mock function with given fields: ctx, nodeName, data
func (_m *NodeInterface) PatchStatus(ctx context.Context, nodeName string, data []byte) (*corev1.Node, error) {
	ret := _m.Called(ctx, nodeName, data)

	var r0 *corev1.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) (*corev1.Node, error)); ok {
		return rf(ctx, nodeName, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) *corev1.Node); ok {
		r0 = rf(ctx, nodeName, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []byte) error); ok {
		r1 = rf(ctx, nodeName, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeInterface_PatchStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchStatus'
type NodeInterface_PatchStatus_Call struct {
	*mock.Call
}

// PatchStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeName string
//   - data []byte
func (_e *NodeInterface_Expecter) PatchStatus(ctx interface{}, nodeName interface{}, data interface{}) *NodeInterface_PatchStatus_Call {
	return &NodeInterface_PatchStatus_Call{Call: _e.mock.On("PatchStatus", ctx, nodeName, data)}
}

func (_c *NodeInterface_PatchStatus_Call) Run(run func(ctx context.Context, nodeName string, data []byte)) *NodeInterface_PatchStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]byte))
	})
	return _c
}

func (_c *NodeInterface_PatchStatus_Call) Return(_a0 *corev1.Node, _a1 error) *NodeInterface_PatchStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NodeInterface_PatchStatus_Call) RunAndReturn(run func(context.Context, string, []byte) (*corev1.Node, error)) *NodeInterface_PatchStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, node, opts
func (_m *NodeInterface) Update(ctx context.Context, node *corev1.Node, opts metav1.UpdateOptions) (*corev1.Node, error) {
	ret := _m.Called(ctx, node, opts)

	var r0 *corev1.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Node, metav1.UpdateOptions) (*corev1.Node, error)); ok {
		return rf(ctx, node, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Node, metav1.UpdateOptions) *corev1.Node); ok {
		r0 = rf(ctx, node, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *corev1.Node, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, node, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type NodeInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - node *corev1.Node
//   - opts metav1.UpdateOptions
func (_e *NodeInterface_Expecter) Update(ctx interface{}, node interface{}, opts interface{}) *NodeInterface_Update_Call {
	return &NodeInterface_Update_Call{Call: _e.mock.On("Update", ctx, node, opts)}
}

func (_c *NodeInterface_Update_Call) Run(run func(ctx context.Context, node *corev1.Node, opts metav1.UpdateOptions)) *NodeInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.Node), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *NodeInterface_Update_Call) Return(_a0 *corev1.Node, _a1 error) *NodeInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NodeInterface_Update_Call) RunAndReturn(run func(context.Context, *corev1.Node, metav1.UpdateOptions) (*corev1.Node, error)) *NodeInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStatus provides a mock function with given fields: ctx, node, opts
func (_m *NodeInterface) UpdateStatus(ctx context.Context, node *corev1.Node, opts metav1.UpdateOptions) (*corev1.Node, error) {
	ret := _m.Called(ctx, node, opts)

	var r0 *corev1.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Node, metav1.UpdateOptions) (*corev1.Node, error)); ok {
		return rf(ctx, node, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Node, metav1.UpdateOptions) *corev1.Node); ok {
		r0 = rf(ctx, node, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *corev1.Node, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, node, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeInterface_UpdateStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStatus'
type NodeInterface_UpdateStatus_Call struct {
	*mock.Call
}

// UpdateStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - node *corev1.Node
//   - opts metav1.UpdateOptions
func (_e *NodeInterface_Expecter) UpdateStatus(ctx interface{}, node interface{}, opts interface{}) *NodeInterface_UpdateStatus_Call {
	return &NodeInterface_UpdateStatus_Call{Call: _e.mock.On("UpdateStatus", ctx, node, opts)}
}

func (_c *NodeInterface_UpdateStatus_Call) Run(run func(ctx context.Context, node *corev1.Node, opts metav1.UpdateOptions)) *NodeInterface_UpdateStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.Node), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *NodeInterface_UpdateStatus_Call) Return(_a0 *corev1.Node, _a1 error) *NodeInterface_UpdateStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NodeInterface_UpdateStatus_Call) RunAndReturn(run func(context.Context, *corev1.Node, metav1.UpdateOptions) (*corev1.Node, error)) *NodeInterface_UpdateStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *NodeInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
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

// NodeInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type NodeInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *NodeInterface_Expecter) Watch(ctx interface{}, opts interface{}) *NodeInterface_Watch_Call {
	return &NodeInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *NodeInterface_Watch_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *NodeInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *NodeInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *NodeInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NodeInterface_Watch_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (watch.Interface, error)) *NodeInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewNodeInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewNodeInterface creates a new instance of NodeInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNodeInterface(t mockConstructorTestingTNewNodeInterface) *NodeInterface {
	mock := &NodeInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
