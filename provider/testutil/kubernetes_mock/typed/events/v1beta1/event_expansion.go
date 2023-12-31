// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/api/events/v1beta1"
)

// EventExpansion is an autogenerated mock type for the EventExpansion type
type EventExpansion struct {
	mock.Mock
}

type EventExpansion_Expecter struct {
	mock *mock.Mock
}

func (_m *EventExpansion) EXPECT() *EventExpansion_Expecter {
	return &EventExpansion_Expecter{mock: &_m.Mock}
}

// CreateWithEventNamespace provides a mock function with given fields: event
func (_m *EventExpansion) CreateWithEventNamespace(event *v1beta1.Event) (*v1beta1.Event, error) {
	ret := _m.Called(event)

	var r0 *v1beta1.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(*v1beta1.Event) (*v1beta1.Event, error)); ok {
		return rf(event)
	}
	if rf, ok := ret.Get(0).(func(*v1beta1.Event) *v1beta1.Event); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(*v1beta1.Event) error); ok {
		r1 = rf(event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventExpansion_CreateWithEventNamespace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateWithEventNamespace'
type EventExpansion_CreateWithEventNamespace_Call struct {
	*mock.Call
}

// CreateWithEventNamespace is a helper method to define mock.On call
//   - event *v1beta1.Event
func (_e *EventExpansion_Expecter) CreateWithEventNamespace(event interface{}) *EventExpansion_CreateWithEventNamespace_Call {
	return &EventExpansion_CreateWithEventNamespace_Call{Call: _e.mock.On("CreateWithEventNamespace", event)}
}

func (_c *EventExpansion_CreateWithEventNamespace_Call) Run(run func(event *v1beta1.Event)) *EventExpansion_CreateWithEventNamespace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*v1beta1.Event))
	})
	return _c
}

func (_c *EventExpansion_CreateWithEventNamespace_Call) Return(_a0 *v1beta1.Event, _a1 error) *EventExpansion_CreateWithEventNamespace_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EventExpansion_CreateWithEventNamespace_Call) RunAndReturn(run func(*v1beta1.Event) (*v1beta1.Event, error)) *EventExpansion_CreateWithEventNamespace_Call {
	_c.Call.Return(run)
	return _c
}

// PatchWithEventNamespace provides a mock function with given fields: event, data
func (_m *EventExpansion) PatchWithEventNamespace(event *v1beta1.Event, data []byte) (*v1beta1.Event, error) {
	ret := _m.Called(event, data)

	var r0 *v1beta1.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(*v1beta1.Event, []byte) (*v1beta1.Event, error)); ok {
		return rf(event, data)
	}
	if rf, ok := ret.Get(0).(func(*v1beta1.Event, []byte) *v1beta1.Event); ok {
		r0 = rf(event, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(*v1beta1.Event, []byte) error); ok {
		r1 = rf(event, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventExpansion_PatchWithEventNamespace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchWithEventNamespace'
type EventExpansion_PatchWithEventNamespace_Call struct {
	*mock.Call
}

// PatchWithEventNamespace is a helper method to define mock.On call
//   - event *v1beta1.Event
//   - data []byte
func (_e *EventExpansion_Expecter) PatchWithEventNamespace(event interface{}, data interface{}) *EventExpansion_PatchWithEventNamespace_Call {
	return &EventExpansion_PatchWithEventNamespace_Call{Call: _e.mock.On("PatchWithEventNamespace", event, data)}
}

func (_c *EventExpansion_PatchWithEventNamespace_Call) Run(run func(event *v1beta1.Event, data []byte)) *EventExpansion_PatchWithEventNamespace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*v1beta1.Event), args[1].([]byte))
	})
	return _c
}

func (_c *EventExpansion_PatchWithEventNamespace_Call) Return(_a0 *v1beta1.Event, _a1 error) *EventExpansion_PatchWithEventNamespace_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EventExpansion_PatchWithEventNamespace_Call) RunAndReturn(run func(*v1beta1.Event, []byte) (*v1beta1.Event, error)) *EventExpansion_PatchWithEventNamespace_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateWithEventNamespace provides a mock function with given fields: event
func (_m *EventExpansion) UpdateWithEventNamespace(event *v1beta1.Event) (*v1beta1.Event, error) {
	ret := _m.Called(event)

	var r0 *v1beta1.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(*v1beta1.Event) (*v1beta1.Event, error)); ok {
		return rf(event)
	}
	if rf, ok := ret.Get(0).(func(*v1beta1.Event) *v1beta1.Event); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(*v1beta1.Event) error); ok {
		r1 = rf(event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventExpansion_UpdateWithEventNamespace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateWithEventNamespace'
type EventExpansion_UpdateWithEventNamespace_Call struct {
	*mock.Call
}

// UpdateWithEventNamespace is a helper method to define mock.On call
//   - event *v1beta1.Event
func (_e *EventExpansion_Expecter) UpdateWithEventNamespace(event interface{}) *EventExpansion_UpdateWithEventNamespace_Call {
	return &EventExpansion_UpdateWithEventNamespace_Call{Call: _e.mock.On("UpdateWithEventNamespace", event)}
}

func (_c *EventExpansion_UpdateWithEventNamespace_Call) Run(run func(event *v1beta1.Event)) *EventExpansion_UpdateWithEventNamespace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*v1beta1.Event))
	})
	return _c
}

func (_c *EventExpansion_UpdateWithEventNamespace_Call) Return(_a0 *v1beta1.Event, _a1 error) *EventExpansion_UpdateWithEventNamespace_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EventExpansion_UpdateWithEventNamespace_Call) RunAndReturn(run func(*v1beta1.Event) (*v1beta1.Event, error)) *EventExpansion_UpdateWithEventNamespace_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewEventExpansion interface {
	mock.TestingT
	Cleanup(func())
}

// NewEventExpansion creates a new instance of EventExpansion. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventExpansion(t mockConstructorTestingTNewEventExpansion) *EventExpansion {
	mock := &EventExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
