// Code generated by mockery v2.42.2. DO NOT EDIT.

package log

import (
	context "context"
	log "server/app/pkg/log"

	mock "github.com/stretchr/testify/mock"
)

// MockHandler is an autogenerated mock type for the Handler type
type MockHandler struct {
	mock.Mock
}

type MockHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHandler) EXPECT() *MockHandler_Expecter {
	return &MockHandler_Expecter{mock: &_m.Mock}
}

// handle provides a mock function with given fields: ctx, level, _a2, opts
func (_m *MockHandler) handle(ctx context.Context, level log.LogLevel, _a2 interface{}, opts ...log.Option) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, level, _a2)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockHandler_handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'handle'
type MockHandler_handle_Call struct {
	*mock.Call
}

// handle is a helper method to define mock.On call
//   - ctx context.Context
//   - level log.LogLevel
//   - _a2 interface{}
//   - opts ...log.Option
func (_e *MockHandler_Expecter) handle(ctx interface{}, level interface{}, _a2 interface{}, opts ...interface{}) *MockHandler_handle_Call {
	return &MockHandler_handle_Call{Call: _e.mock.On("handle",
		append([]interface{}{ctx, level, _a2}, opts...)...)}
}

func (_c *MockHandler_handle_Call) Run(run func(ctx context.Context, level log.LogLevel, _a2 interface{}, opts ...log.Option)) *MockHandler_handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]log.Option, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(log.Option)
			}
		}
		run(args[0].(context.Context), args[1].(log.LogLevel), args[2].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *MockHandler_handle_Call) Return() *MockHandler_handle_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockHandler_handle_Call) RunAndReturn(run func(context.Context, log.LogLevel, interface{}, ...log.Option)) *MockHandler_handle_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockHandler creates a new instance of MockHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHandler {
	mock := &MockHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}