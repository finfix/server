// Code generated by mockery v2.42.2. DO NOT EDIT.

package migrator

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockMigrator is an autogenerated mock type for the Migrator type
type MockMigrator struct {
	mock.Mock
}

type MockMigrator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMigrator) EXPECT() *MockMigrator_Expecter {
	return &MockMigrator_Expecter{mock: &_m.Mock}
}

// Down provides a mock function with given fields: _a0
func (_m *MockMigrator) Down(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Down")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMigrator_Down_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Down'
type MockMigrator_Down_Call struct {
	*mock.Call
}

// Down is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockMigrator_Expecter) Down(_a0 interface{}) *MockMigrator_Down_Call {
	return &MockMigrator_Down_Call{Call: _e.mock.On("Down", _a0)}
}

func (_c *MockMigrator_Down_Call) Run(run func(_a0 context.Context)) *MockMigrator_Down_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockMigrator_Down_Call) Return(_a0 error) *MockMigrator_Down_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMigrator_Down_Call) RunAndReturn(run func(context.Context) error) *MockMigrator_Down_Call {
	_c.Call.Return(run)
	return _c
}

// Up provides a mock function with given fields: _a0
func (_m *MockMigrator) Up(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Up")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMigrator_Up_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Up'
type MockMigrator_Up_Call struct {
	*mock.Call
}

// Up is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockMigrator_Expecter) Up(_a0 interface{}) *MockMigrator_Up_Call {
	return &MockMigrator_Up_Call{Call: _e.mock.On("Up", _a0)}
}

func (_c *MockMigrator_Up_Call) Run(run func(_a0 context.Context)) *MockMigrator_Up_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockMigrator_Up_Call) Return(_a0 error) *MockMigrator_Up_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMigrator_Up_Call) RunAndReturn(run func(context.Context) error) *MockMigrator_Up_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMigrator creates a new instance of MockMigrator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMigrator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMigrator {
	mock := &MockMigrator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}