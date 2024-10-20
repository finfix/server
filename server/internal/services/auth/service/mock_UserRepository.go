// Code generated by mockery v2.46.2. DO NOT EDIT.

package service

import (
	context "context"
	model "server/internal/services/user/model"

	mock "github.com/stretchr/testify/mock"

	repositorymodel "server/internal/services/user/repository/model"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

// CreateDevice provides a mock function with given fields: _a0, _a1
func (_m *MockUserRepository) CreateDevice(_a0 context.Context, _a1 model.Device) (uint32, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateDevice")
	}

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Device) (uint32, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Device) uint32); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Device) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: _a0, _a1
func (_m *MockUserRepository) CreateUser(_a0 context.Context, _a1 model.CreateReq) (uint32, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateReq) (uint32, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateReq) uint32); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CreateReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDevice provides a mock function with given fields: ctx, userID, deviceID
func (_m *MockUserRepository) DeleteDevice(ctx context.Context, userID uint32, deviceID string) error {
	ret := _m.Called(ctx, userID, deviceID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDevice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, string) error); ok {
		r0 = rf(ctx, userID, deviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDevices provides a mock function with given fields: _a0, _a1
func (_m *MockUserRepository) GetDevices(_a0 context.Context, _a1 repositorymodel.GetDevicesReq) ([]model.Device, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetDevices")
	}

	var r0 []model.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repositorymodel.GetDevicesReq) ([]model.Device, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repositorymodel.GetDevicesReq) []model.Device); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, repositorymodel.GetDevicesReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: _a0, _a1
func (_m *MockUserRepository) GetUsers(_a0 context.Context, _a1 model.GetUsersReq) ([]model.User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetUsers")
	}

	var r0 []model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetUsersReq) ([]model.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetUsersReq) []model.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetUsersReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDevice provides a mock function with given fields: _a0, _a1
func (_m *MockUserRepository) UpdateDevice(_a0 context.Context, _a1 repositorymodel.UpdateDeviceReq) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDevice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repositorymodel.UpdateDeviceReq) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockUserRepository creates a new instance of MockUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserRepository {
	mock := &MockUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}