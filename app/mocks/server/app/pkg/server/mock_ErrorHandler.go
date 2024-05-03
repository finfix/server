// Code generated by mockery v2.42.2. DO NOT EDIT.

package server

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockErrorHandler is an autogenerated mock type for the ErrorHandler type
type MockErrorHandler struct {
	mock.Mock
}

type MockErrorHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockErrorHandler) EXPECT() *MockErrorHandler_Expecter {
	return &MockErrorHandler_Expecter{mock: &_m.Mock}
}

// Handle provides a mock function with given fields: ctx, err
func (_m *MockErrorHandler) Handle(ctx context.Context, err error) {
	_m.Called(ctx, err)
}

// MockErrorHandler_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type MockErrorHandler_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - ctx context.Context
//   - err error
func (_e *MockErrorHandler_Expecter) Handle(ctx interface{}, err interface{}) *MockErrorHandler_Handle_Call {
	return &MockErrorHandler_Handle_Call{Call: _e.mock.On("Handle", ctx, err)}
}

func (_c *MockErrorHandler_Handle_Call) Run(run func(ctx context.Context, err error)) *MockErrorHandler_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(error))
	})
	return _c
}

func (_c *MockErrorHandler_Handle_Call) Return() *MockErrorHandler_Handle_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockErrorHandler_Handle_Call) RunAndReturn(run func(context.Context, error)) *MockErrorHandler_Handle_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockErrorHandler creates a new instance of MockErrorHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockErrorHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockErrorHandler {
	mock := &MockErrorHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
