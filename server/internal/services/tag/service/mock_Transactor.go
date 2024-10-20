// Code generated by mockery v2.46.2. DO NOT EDIT.

package service

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockTransactor is an autogenerated mock type for the Transactor type
type MockTransactor struct {
	mock.Mock
}

// WithinTransaction provides a mock function with given fields: ctx, callback
func (_m *MockTransactor) WithinTransaction(ctx context.Context, callback func(context.Context) error) error {
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

// NewMockTransactor creates a new instance of MockTransactor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTransactor(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTransactor {
	mock := &MockTransactor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
