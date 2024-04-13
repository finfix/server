// Code generated by mockery v2.42.2. DO NOT EDIT.

package service

import (
	context "context"
	model "server/app/services/user/model"

	mock "github.com/stretchr/testify/mock"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

type MockUserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserRepository) EXPECT() *MockUserRepository_Expecter {
	return &MockUserRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *MockUserRepository) Create(_a0 context.Context, _a1 model.CreateReq) (uint32, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateReq) (uint32, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateReq) uint32); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CreateReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockUserRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.CreateReq
func (_e *MockUserRepository_Expecter) Create(_a0 interface{}, _a1 interface{}) *MockUserRepository_Create_Call {
	return &MockUserRepository_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *MockUserRepository_Create_Call) Run(run func(_a0 context.Context, _a1 model.CreateReq)) *MockUserRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.CreateReq))
	})
	return _c
}

func (_c *MockUserRepository_Create_Call) Return(_a0 uint32, _a1 error) *MockUserRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_Create_Call) RunAndReturn(run func(context.Context, model.CreateReq) (uint32, error)) *MockUserRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// GetCurrencies provides a mock function with given fields: _a0
func (_m *MockUserRepository) GetCurrencies(_a0 context.Context) ([]model.Currency, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetCurrencies")
	}

	var r0 []model.Currency
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.Currency, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.Currency); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Currency)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_GetCurrencies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCurrencies'
type MockUserRepository_GetCurrencies_Call struct {
	*mock.Call
}

// GetCurrencies is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockUserRepository_Expecter) GetCurrencies(_a0 interface{}) *MockUserRepository_GetCurrencies_Call {
	return &MockUserRepository_GetCurrencies_Call{Call: _e.mock.On("GetCurrencies", _a0)}
}

func (_c *MockUserRepository_GetCurrencies_Call) Run(run func(_a0 context.Context)) *MockUserRepository_GetCurrencies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockUserRepository_GetCurrencies_Call) Return(_a0 []model.Currency, _a1 error) *MockUserRepository_GetCurrencies_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetCurrencies_Call) RunAndReturn(run func(context.Context) ([]model.Currency, error)) *MockUserRepository_GetCurrencies_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransactions provides a mock function with given fields: _a0, _a1
func (_m *MockUserRepository) GetTransactions(_a0 context.Context, _a1 model.GetReq) ([]model.User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactions")
	}

	var r0 []model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetReq) ([]model.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetReq) []model.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_GetTransactions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransactions'
type MockUserRepository_GetTransactions_Call struct {
	*mock.Call
}

// GetTransactions is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.GetReq
func (_e *MockUserRepository_Expecter) GetTransactions(_a0 interface{}, _a1 interface{}) *MockUserRepository_GetTransactions_Call {
	return &MockUserRepository_GetTransactions_Call{Call: _e.mock.On("GetTransactions", _a0, _a1)}
}

func (_c *MockUserRepository_GetTransactions_Call) Run(run func(_a0 context.Context, _a1 model.GetReq)) *MockUserRepository_GetTransactions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.GetReq))
	})
	return _c
}

func (_c *MockUserRepository_GetTransactions_Call) Return(_a0 []model.User, _a1 error) *MockUserRepository_GetTransactions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetTransactions_Call) RunAndReturn(run func(context.Context, model.GetReq) ([]model.User, error)) *MockUserRepository_GetTransactions_Call {
	_c.Call.Return(run)
	return _c
}

// LinkUserToAccountGroup provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockUserRepository) LinkUserToAccountGroup(_a0 context.Context, _a1 uint32, _a2 uint32) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for LinkUserToAccountGroup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_LinkUserToAccountGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LinkUserToAccountGroup'
type MockUserRepository_LinkUserToAccountGroup_Call struct {
	*mock.Call
}

// LinkUserToAccountGroup is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint32
//   - _a2 uint32
func (_e *MockUserRepository_Expecter) LinkUserToAccountGroup(_a0 interface{}, _a1 interface{}, _a2 interface{}) *MockUserRepository_LinkUserToAccountGroup_Call {
	return &MockUserRepository_LinkUserToAccountGroup_Call{Call: _e.mock.On("LinkUserToAccountGroup", _a0, _a1, _a2)}
}

func (_c *MockUserRepository_LinkUserToAccountGroup_Call) Run(run func(_a0 context.Context, _a1 uint32, _a2 uint32)) *MockUserRepository_LinkUserToAccountGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint32))
	})
	return _c
}

func (_c *MockUserRepository_LinkUserToAccountGroup_Call) Return(_a0 error) *MockUserRepository_LinkUserToAccountGroup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_LinkUserToAccountGroup_Call) RunAndReturn(run func(context.Context, uint32, uint32) error) *MockUserRepository_LinkUserToAccountGroup_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserRepository creates a new instance of MockUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserRepository {
	mock := &MockUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}