// Code generated by mockery v2.46.2. DO NOT EDIT.

package service

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockAccountService is an autogenerated mock type for the AccountService type
type MockAccountService struct {
	mock.Mock
}

// CheckAccess provides a mock function with given fields: ctx, userID, accountIDs
func (_m *MockAccountService) CheckAccess(ctx context.Context, userID uint32, accountIDs []uint32) error {
	ret := _m.Called(ctx, userID, accountIDs)

	if len(ret) == 0 {
		panic("no return value specified for CheckAccess")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, []uint32) error); ok {
		r0 = rf(ctx, userID, accountIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockAccountService creates a new instance of MockAccountService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAccountService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountService {
	mock := &MockAccountService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
