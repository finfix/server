// Code generated by mockery v2.42.2. DO NOT EDIT.

package server

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MockEncodeErrorFunc is an autogenerated mock type for the EncodeErrorFunc type
type MockEncodeErrorFunc struct {
	mock.Mock
}

type MockEncodeErrorFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEncodeErrorFunc) EXPECT() *MockEncodeErrorFunc_Expecter {
	return &MockEncodeErrorFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockEncodeErrorFunc) Execute(_a0 context.Context, _a1 http.ResponseWriter, _a2 error) {
	_m.Called(_a0, _a1, _a2)
}

// MockEncodeErrorFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockEncodeErrorFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 http.ResponseWriter
//   - _a2 error
func (_e *MockEncodeErrorFunc_Expecter) Execute(_a0 interface{}, _a1 interface{}, _a2 interface{}) *MockEncodeErrorFunc_Execute_Call {
	return &MockEncodeErrorFunc_Execute_Call{Call: _e.mock.On("Execute", _a0, _a1, _a2)}
}

func (_c *MockEncodeErrorFunc_Execute_Call) Run(run func(_a0 context.Context, _a1 http.ResponseWriter, _a2 error)) *MockEncodeErrorFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(http.ResponseWriter), args[2].(error))
	})
	return _c
}

func (_c *MockEncodeErrorFunc_Execute_Call) Return() *MockEncodeErrorFunc_Execute_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockEncodeErrorFunc_Execute_Call) RunAndReturn(run func(context.Context, http.ResponseWriter, error)) *MockEncodeErrorFunc_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEncodeErrorFunc creates a new instance of MockEncodeErrorFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEncodeErrorFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEncodeErrorFunc {
	mock := &MockEncodeErrorFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}