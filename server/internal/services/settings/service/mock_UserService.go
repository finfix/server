// Code generated by mockery v2.46.2. DO NOT EDIT.

package service

import (
	context "context"
	model "server/internal/services/user/model"

	mock "github.com/stretchr/testify/mock"
)

// MockUserService is an autogenerated mock type for the UserService type
type MockUserService struct {
	mock.Mock
}

// GetUsers provides a mock function with given fields: ctx, filters
func (_m *MockUserService) GetUsers(ctx context.Context, filters model.GetUsersReq) ([]model.User, error) {
	ret := _m.Called(ctx, filters)

	if len(ret) == 0 {
		panic("no return value specified for GetUsers")
	}

	var r0 []model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetUsersReq) ([]model.User, error)); ok {
		return rf(ctx, filters)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetUsersReq) []model.User); ok {
		r0 = rf(ctx, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetUsersReq) error); ok {
		r1 = rf(ctx, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendNotification provides a mock function with given fields: ctx, userID, push
func (_m *MockUserService) SendNotification(ctx context.Context, userID uint32, push model.Notification) (uint8, error) {
	ret := _m.Called(ctx, userID, push)

	if len(ret) == 0 {
		panic("no return value specified for SendNotification")
	}

	var r0 uint8
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, model.Notification) (uint8, error)); ok {
		return rf(ctx, userID, push)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, model.Notification) uint8); ok {
		r0 = rf(ctx, userID, push)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, model.Notification) error); ok {
		r1 = rf(ctx, userID, push)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockUserService creates a new instance of MockUserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserService {
	mock := &MockUserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
