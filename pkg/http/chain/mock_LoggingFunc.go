// Code generated by mockery v2.42.2. DO NOT EDIT.

package chain

import mock "github.com/stretchr/testify/mock"

// MockLoggingFunc is an autogenerated mock type for the LoggingFunc type
type MockLoggingFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *MockLoggingFunc) Execute(_a0 error) {
	_m.Called(_a0)
}

// NewMockLoggingFunc creates a new instance of MockLoggingFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLoggingFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLoggingFunc {
	mock := &MockLoggingFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
