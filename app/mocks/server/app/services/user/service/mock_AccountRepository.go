// Code generated by mockery v2.42.2. DO NOT EDIT.

package service

import (
	context "context"
	accountmodel "server/app/services/account/model"

	mock "github.com/stretchr/testify/mock"

	model "server/app/services/account/repository/model"
)

// MockAccountRepository is an autogenerated mock type for the AccountRepository type
type MockAccountRepository struct {
	mock.Mock
}

type MockAccountRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccountRepository) EXPECT() *MockAccountRepository_Expecter {
	return &MockAccountRepository_Expecter{mock: &_m.Mock}
}

// CreateAccount provides a mock function with given fields: ctx, req
func (_m *MockAccountRepository) CreateAccount(ctx context.Context, req model.CreateAccountReq) (uint32, uint32, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccount")
	}

	var r0 uint32
	var r1 uint32
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateAccountReq) (uint32, uint32, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateAccountReq) uint32); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CreateAccountReq) uint32); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	if rf, ok := ret.Get(2).(func(context.Context, model.CreateAccountReq) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockAccountRepository_CreateAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAccount'
type MockAccountRepository_CreateAccount_Call struct {
	*mock.Call
}

// CreateAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - req model.CreateAccountReq
func (_e *MockAccountRepository_Expecter) CreateAccount(ctx interface{}, req interface{}) *MockAccountRepository_CreateAccount_Call {
	return &MockAccountRepository_CreateAccount_Call{Call: _e.mock.On("CreateAccount", ctx, req)}
}

func (_c *MockAccountRepository_CreateAccount_Call) Run(run func(ctx context.Context, req model.CreateAccountReq)) *MockAccountRepository_CreateAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.CreateAccountReq))
	})
	return _c
}

func (_c *MockAccountRepository_CreateAccount_Call) Return(_a0 uint32, _a1 uint32, _a2 error) *MockAccountRepository_CreateAccount_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockAccountRepository_CreateAccount_Call) RunAndReturn(run func(context.Context, model.CreateAccountReq) (uint32, uint32, error)) *MockAccountRepository_CreateAccount_Call {
	_c.Call.Return(run)
	return _c
}

// CreateAccountGroup provides a mock function with given fields: _a0, _a1
func (_m *MockAccountRepository) CreateAccountGroup(_a0 context.Context, _a1 accountmodel.CreateAccountGroupReq) (uint32, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccountGroup")
	}

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, accountmodel.CreateAccountGroupReq) (uint32, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, accountmodel.CreateAccountGroupReq) uint32); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, accountmodel.CreateAccountGroupReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountRepository_CreateAccountGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAccountGroup'
type MockAccountRepository_CreateAccountGroup_Call struct {
	*mock.Call
}

// CreateAccountGroup is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 accountmodel.CreateAccountGroupReq
func (_e *MockAccountRepository_Expecter) CreateAccountGroup(_a0 interface{}, _a1 interface{}) *MockAccountRepository_CreateAccountGroup_Call {
	return &MockAccountRepository_CreateAccountGroup_Call{Call: _e.mock.On("CreateAccountGroup", _a0, _a1)}
}

func (_c *MockAccountRepository_CreateAccountGroup_Call) Run(run func(_a0 context.Context, _a1 accountmodel.CreateAccountGroupReq)) *MockAccountRepository_CreateAccountGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(accountmodel.CreateAccountGroupReq))
	})
	return _c
}

func (_c *MockAccountRepository_CreateAccountGroup_Call) Return(_a0 uint32, _a1 error) *MockAccountRepository_CreateAccountGroup_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountRepository_CreateAccountGroup_Call) RunAndReturn(run func(context.Context, accountmodel.CreateAccountGroupReq) (uint32, error)) *MockAccountRepository_CreateAccountGroup_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAccountRepository creates a new instance of MockAccountRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAccountRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountRepository {
	mock := &MockAccountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}