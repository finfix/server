// Code generated by mockery v2.46.2. DO NOT EDIT.

package service

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockTagService is an autogenerated mock type for the TagService type
type MockTagService struct {
	mock.Mock
}

// CheckAccess provides a mock function with given fields: ctx, userID, tagIDs
func (_m *MockTagService) CheckAccess(ctx context.Context, userID uint32, tagIDs []uint32) error {
	ret := _m.Called(ctx, userID, tagIDs)

	if len(ret) == 0 {
		panic("no return value specified for CheckAccess")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, []uint32) error); ok {
		r0 = rf(ctx, userID, tagIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockTagService creates a new instance of MockTagService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTagService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTagService {
	mock := &MockTagService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
