// Code generated by mockery v2.24.0. DO NOT EDIT.

package mocks

import (
	authz "github.com/cosmos/cosmos-sdk/x/authz"

	mock "github.com/stretchr/testify/mock"

	time "time"

	types "github.com/cosmos/cosmos-sdk/types"
)

// AuthzKeeper is an autogenerated mock type for the AuthzKeeper type
type AuthzKeeper struct {
	mock.Mock
}

type AuthzKeeper_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthzKeeper) EXPECT() *AuthzKeeper_Expecter {
	return &AuthzKeeper_Expecter{mock: &_m.Mock}
}

// DeleteGrant provides a mock function with given fields: ctx, grantee, granter, msgType
func (_m *AuthzKeeper) DeleteGrant(ctx types.Context, grantee types.AccAddress, granter types.AccAddress, msgType string) error {
	ret := _m.Called(ctx, grantee, granter, msgType)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, types.AccAddress, string) error); ok {
		r0 = rf(ctx, grantee, granter, msgType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthzKeeper_DeleteGrant_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteGrant'
type AuthzKeeper_DeleteGrant_Call struct {
	*mock.Call
}

// DeleteGrant is a helper method to define mock.On call
//   - ctx types.Context
//   - grantee types.AccAddress
//   - granter types.AccAddress
//   - msgType string
func (_e *AuthzKeeper_Expecter) DeleteGrant(ctx interface{}, grantee interface{}, granter interface{}, msgType interface{}) *AuthzKeeper_DeleteGrant_Call {
	return &AuthzKeeper_DeleteGrant_Call{Call: _e.mock.On("DeleteGrant", ctx, grantee, granter, msgType)}
}

func (_c *AuthzKeeper_DeleteGrant_Call) Run(run func(ctx types.Context, grantee types.AccAddress, granter types.AccAddress, msgType string)) *AuthzKeeper_DeleteGrant_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(types.AccAddress), args[2].(types.AccAddress), args[3].(string))
	})
	return _c
}

func (_c *AuthzKeeper_DeleteGrant_Call) Return(_a0 error) *AuthzKeeper_DeleteGrant_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthzKeeper_DeleteGrant_Call) RunAndReturn(run func(types.Context, types.AccAddress, types.AccAddress, string) error) *AuthzKeeper_DeleteGrant_Call {
	_c.Call.Return(run)
	return _c
}

// GetCleanAuthorization provides a mock function with given fields: ctx, grantee, granter, msgType
func (_m *AuthzKeeper) GetCleanAuthorization(ctx types.Context, grantee types.AccAddress, granter types.AccAddress, msgType string) (authz.Authorization, time.Time) {
	ret := _m.Called(ctx, grantee, granter, msgType)

	var r0 authz.Authorization
	var r1 time.Time
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, types.AccAddress, string) (authz.Authorization, time.Time)); ok {
		return rf(ctx, grantee, granter, msgType)
	}
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, types.AccAddress, string) authz.Authorization); ok {
		r0 = rf(ctx, grantee, granter, msgType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(authz.Authorization)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, types.AccAddress, types.AccAddress, string) time.Time); ok {
		r1 = rf(ctx, grantee, granter, msgType)
	} else {
		r1 = ret.Get(1).(time.Time)
	}

	return r0, r1
}

// AuthzKeeper_GetCleanAuthorization_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCleanAuthorization'
type AuthzKeeper_GetCleanAuthorization_Call struct {
	*mock.Call
}

// GetCleanAuthorization is a helper method to define mock.On call
//   - ctx types.Context
//   - grantee types.AccAddress
//   - granter types.AccAddress
//   - msgType string
func (_e *AuthzKeeper_Expecter) GetCleanAuthorization(ctx interface{}, grantee interface{}, granter interface{}, msgType interface{}) *AuthzKeeper_GetCleanAuthorization_Call {
	return &AuthzKeeper_GetCleanAuthorization_Call{Call: _e.mock.On("GetCleanAuthorization", ctx, grantee, granter, msgType)}
}

func (_c *AuthzKeeper_GetCleanAuthorization_Call) Run(run func(ctx types.Context, grantee types.AccAddress, granter types.AccAddress, msgType string)) *AuthzKeeper_GetCleanAuthorization_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(types.AccAddress), args[2].(types.AccAddress), args[3].(string))
	})
	return _c
}

func (_c *AuthzKeeper_GetCleanAuthorization_Call) Return(cap authz.Authorization, expiration time.Time) *AuthzKeeper_GetCleanAuthorization_Call {
	_c.Call.Return(cap, expiration)
	return _c
}

func (_c *AuthzKeeper_GetCleanAuthorization_Call) RunAndReturn(run func(types.Context, types.AccAddress, types.AccAddress, string) (authz.Authorization, time.Time)) *AuthzKeeper_GetCleanAuthorization_Call {
	_c.Call.Return(run)
	return _c
}

// SaveGrant provides a mock function with given fields: ctx, grantee, granter, authorization, expiration
func (_m *AuthzKeeper) SaveGrant(ctx types.Context, grantee types.AccAddress, granter types.AccAddress, authorization authz.Authorization, expiration time.Time) error {
	ret := _m.Called(ctx, grantee, granter, authorization, expiration)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, types.AccAddress, authz.Authorization, time.Time) error); ok {
		r0 = rf(ctx, grantee, granter, authorization, expiration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthzKeeper_SaveGrant_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveGrant'
type AuthzKeeper_SaveGrant_Call struct {
	*mock.Call
}

// SaveGrant is a helper method to define mock.On call
//   - ctx types.Context
//   - grantee types.AccAddress
//   - granter types.AccAddress
//   - authorization authz.Authorization
//   - expiration time.Time
func (_e *AuthzKeeper_Expecter) SaveGrant(ctx interface{}, grantee interface{}, granter interface{}, authorization interface{}, expiration interface{}) *AuthzKeeper_SaveGrant_Call {
	return &AuthzKeeper_SaveGrant_Call{Call: _e.mock.On("SaveGrant", ctx, grantee, granter, authorization, expiration)}
}

func (_c *AuthzKeeper_SaveGrant_Call) Run(run func(ctx types.Context, grantee types.AccAddress, granter types.AccAddress, authorization authz.Authorization, expiration time.Time)) *AuthzKeeper_SaveGrant_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(types.AccAddress), args[2].(types.AccAddress), args[3].(authz.Authorization), args[4].(time.Time))
	})
	return _c
}

func (_c *AuthzKeeper_SaveGrant_Call) Return(_a0 error) *AuthzKeeper_SaveGrant_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthzKeeper_SaveGrant_Call) RunAndReturn(run func(types.Context, types.AccAddress, types.AccAddress, authz.Authorization, time.Time) error) *AuthzKeeper_SaveGrant_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewAuthzKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthzKeeper creates a new instance of AuthzKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthzKeeper(t mockConstructorTestingTNewAuthzKeeper) *AuthzKeeper {
	mock := &AuthzKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}