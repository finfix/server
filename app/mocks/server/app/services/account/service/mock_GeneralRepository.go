// Code generated by mockery v2.42.2. DO NOT EDIT.

package service

import (
	context "context"
	checker "server/app/services/generalRepository/checker"

	decimal "github.com/shopspring/decimal"

	mock "github.com/stretchr/testify/mock"
)

// MockGeneralRepository is an autogenerated mock type for the GeneralRepository type
type MockGeneralRepository struct {
	mock.Mock
}

type MockGeneralRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGeneralRepository) EXPECT() *MockGeneralRepository_Expecter {
	return &MockGeneralRepository_Expecter{mock: &_m.Mock}
}

// CheckUserAccessToObjects provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *MockGeneralRepository) CheckUserAccessToObjects(_a0 context.Context, _a1 checker.CheckType, _a2 uint32, _a3 []uint32) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for CheckUserAccessToObjects")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, checker.CheckType, uint32, []uint32) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockGeneralRepository_CheckUserAccessToObjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckUserAccessToObjects'
type MockGeneralRepository_CheckUserAccessToObjects_Call struct {
	*mock.Call
}

// CheckUserAccessToObjects is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 checker.CheckType
//   - _a2 uint32
//   - _a3 []uint32
func (_e *MockGeneralRepository_Expecter) CheckUserAccessToObjects(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *MockGeneralRepository_CheckUserAccessToObjects_Call {
	return &MockGeneralRepository_CheckUserAccessToObjects_Call{Call: _e.mock.On("CheckUserAccessToObjects", _a0, _a1, _a2, _a3)}
}

func (_c *MockGeneralRepository_CheckUserAccessToObjects_Call) Run(run func(_a0 context.Context, _a1 checker.CheckType, _a2 uint32, _a3 []uint32)) *MockGeneralRepository_CheckUserAccessToObjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(checker.CheckType), args[2].(uint32), args[3].([]uint32))
	})
	return _c
}

func (_c *MockGeneralRepository_CheckUserAccessToObjects_Call) Return(_a0 error) *MockGeneralRepository_CheckUserAccessToObjects_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGeneralRepository_CheckUserAccessToObjects_Call) RunAndReturn(run func(context.Context, checker.CheckType, uint32, []uint32) error) *MockGeneralRepository_CheckUserAccessToObjects_Call {
	_c.Call.Return(run)
	return _c
}

// GetAvailableAccountGroups provides a mock function with given fields: userID
func (_m *MockGeneralRepository) GetAvailableAccountGroups(userID uint32) []uint32 {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetAvailableAccountGroups")
	}

	var r0 []uint32
	if rf, ok := ret.Get(0).(func(uint32) []uint32); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint32)
		}
	}

	return r0
}

// MockGeneralRepository_GetAvailableAccountGroups_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAvailableAccountGroups'
type MockGeneralRepository_GetAvailableAccountGroups_Call struct {
	*mock.Call
}

// GetAvailableAccountGroups is a helper method to define mock.On call
//   - userID uint32
func (_e *MockGeneralRepository_Expecter) GetAvailableAccountGroups(userID interface{}) *MockGeneralRepository_GetAvailableAccountGroups_Call {
	return &MockGeneralRepository_GetAvailableAccountGroups_Call{Call: _e.mock.On("GetAvailableAccountGroups", userID)}
}

func (_c *MockGeneralRepository_GetAvailableAccountGroups_Call) Run(run func(userID uint32)) *MockGeneralRepository_GetAvailableAccountGroups_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *MockGeneralRepository_GetAvailableAccountGroups_Call) Return(_a0 []uint32) *MockGeneralRepository_GetAvailableAccountGroups_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGeneralRepository_GetAvailableAccountGroups_Call) RunAndReturn(run func(uint32) []uint32) *MockGeneralRepository_GetAvailableAccountGroups_Call {
	_c.Call.Return(run)
	return _c
}

// GetCurrencies provides a mock function with given fields: _a0
func (_m *MockGeneralRepository) GetCurrencies(_a0 context.Context) (map[string]decimal.Decimal, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetCurrencies")
	}

	var r0 map[string]decimal.Decimal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (map[string]decimal.Decimal, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) map[string]decimal.Decimal); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]decimal.Decimal)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGeneralRepository_GetCurrencies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCurrencies'
type MockGeneralRepository_GetCurrencies_Call struct {
	*mock.Call
}

// GetCurrencies is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockGeneralRepository_Expecter) GetCurrencies(_a0 interface{}) *MockGeneralRepository_GetCurrencies_Call {
	return &MockGeneralRepository_GetCurrencies_Call{Call: _e.mock.On("GetCurrencies", _a0)}
}

func (_c *MockGeneralRepository_GetCurrencies_Call) Run(run func(_a0 context.Context)) *MockGeneralRepository_GetCurrencies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockGeneralRepository_GetCurrencies_Call) Return(_a0 map[string]decimal.Decimal, _a1 error) *MockGeneralRepository_GetCurrencies_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGeneralRepository_GetCurrencies_Call) RunAndReturn(run func(context.Context) (map[string]decimal.Decimal, error)) *MockGeneralRepository_GetCurrencies_Call {
	_c.Call.Return(run)
	return _c
}

// WithinTransaction provides a mock function with given fields: ctx, callback
func (_m *MockGeneralRepository) WithinTransaction(ctx context.Context, callback func(context.Context) error) error {
	ret := _m.Called(ctx, callback)

	if len(ret) == 0 {
		panic("no return value specified for WithinTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) error) error); ok {
		r0 = rf(ctx, callback)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockGeneralRepository_WithinTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithinTransaction'
type MockGeneralRepository_WithinTransaction_Call struct {
	*mock.Call
}

// WithinTransaction is a helper method to define mock.On call
//   - ctx context.Context
//   - callback func(context.Context) error
func (_e *MockGeneralRepository_Expecter) WithinTransaction(ctx interface{}, callback interface{}) *MockGeneralRepository_WithinTransaction_Call {
	return &MockGeneralRepository_WithinTransaction_Call{Call: _e.mock.On("WithinTransaction", ctx, callback)}
}

func (_c *MockGeneralRepository_WithinTransaction_Call) Run(run func(ctx context.Context, callback func(context.Context) error)) *MockGeneralRepository_WithinTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(context.Context) error))
	})
	return _c
}

func (_c *MockGeneralRepository_WithinTransaction_Call) Return(_a0 error) *MockGeneralRepository_WithinTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGeneralRepository_WithinTransaction_Call) RunAndReturn(run func(context.Context, func(context.Context) error) error) *MockGeneralRepository_WithinTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGeneralRepository creates a new instance of MockGeneralRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGeneralRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGeneralRepository {
	mock := &MockGeneralRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}