// Code generated by mockery v2.42.2. DO NOT EDIT.

package service

import (
	context "context"
	model "server/app/services/user/model"

	mock "github.com/stretchr/testify/mock"

	repositorymodel "server/app/services/user/repository/model"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

type MockUserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserRepository) EXPECT() *MockUserRepository_Expecter {
	return &MockUserRepository_Expecter{mock: &_m.Mock}
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

// MockUserRepository_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockUserRepository_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.CreateReq
func (_e *MockUserRepository_Expecter) CreateUser(_a0 interface{}, _a1 interface{}) *MockUserRepository_CreateUser_Call {
	return &MockUserRepository_CreateUser_Call{Call: _e.mock.On("CreateUser", _a0, _a1)}
}

func (_c *MockUserRepository_CreateUser_Call) Run(run func(_a0 context.Context, _a1 model.CreateReq)) *MockUserRepository_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.CreateReq))
	})
	return _c
}

func (_c *MockUserRepository_CreateUser_Call) Return(_a0 uint32, _a1 error) *MockUserRepository_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_CreateUser_Call) RunAndReturn(run func(context.Context, model.CreateReq) (uint32, error)) *MockUserRepository_CreateUser_Call {
	_c.Call.Return(run)
	return _c
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

// MockUserRepository_GetDevices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDevices'
type MockUserRepository_GetDevices_Call struct {
	*mock.Call
}

// GetDevices is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 repositorymodel.GetDevicesReq
func (_e *MockUserRepository_Expecter) GetDevices(_a0 interface{}, _a1 interface{}) *MockUserRepository_GetDevices_Call {
	return &MockUserRepository_GetDevices_Call{Call: _e.mock.On("GetDevices", _a0, _a1)}
}

func (_c *MockUserRepository_GetDevices_Call) Run(run func(_a0 context.Context, _a1 repositorymodel.GetDevicesReq)) *MockUserRepository_GetDevices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(repositorymodel.GetDevicesReq))
	})
	return _c
}

func (_c *MockUserRepository_GetDevices_Call) Return(_a0 []model.Device, _a1 error) *MockUserRepository_GetDevices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetDevices_Call) RunAndReturn(run func(context.Context, repositorymodel.GetDevicesReq) ([]model.Device, error)) *MockUserRepository_GetDevices_Call {
	_c.Call.Return(run)
	return _c
}

// GetUsers provides a mock function with given fields: _a0, _a1
func (_m *MockUserRepository) GetUsers(_a0 context.Context, _a1 model.GetReq) ([]model.User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetUsers")
	}

	var r0 []model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetReq) ([]model.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetReq) []model.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_GetUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUsers'
type MockUserRepository_GetUsers_Call struct {
	*mock.Call
}

// GetUsers is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.GetReq
func (_e *MockUserRepository_Expecter) GetUsers(_a0 interface{}, _a1 interface{}) *MockUserRepository_GetUsers_Call {
	return &MockUserRepository_GetUsers_Call{Call: _e.mock.On("GetUsers", _a0, _a1)}
}

func (_c *MockUserRepository_GetUsers_Call) Run(run func(_a0 context.Context, _a1 model.GetReq)) *MockUserRepository_GetUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.GetReq))
	})
	return _c
}

func (_c *MockUserRepository_GetUsers_Call) Return(_a0 []model.User, _a1 error) *MockUserRepository_GetUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetUsers_Call) RunAndReturn(run func(context.Context, model.GetReq) ([]model.User, error)) *MockUserRepository_GetUsers_Call {
	_c.Call.Return(run)
	return _c
}

// LinkUserToAccountGroup provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockUserRepository) LinkUserToAccountGroup(_a0 context.Context, _a1 uint32, _a2 uint32) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for LinkUserToAccountGroup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_LinkUserToAccountGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LinkUserToAccountGroup'
type MockUserRepository_LinkUserToAccountGroup_Call struct {
	*mock.Call
}

// LinkUserToAccountGroup is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint32
//   - _a2 uint32
func (_e *MockUserRepository_Expecter) LinkUserToAccountGroup(_a0 interface{}, _a1 interface{}, _a2 interface{}) *MockUserRepository_LinkUserToAccountGroup_Call {
	return &MockUserRepository_LinkUserToAccountGroup_Call{Call: _e.mock.On("LinkUserToAccountGroup", _a0, _a1, _a2)}
}

func (_c *MockUserRepository_LinkUserToAccountGroup_Call) Run(run func(_a0 context.Context, _a1 uint32, _a2 uint32)) *MockUserRepository_LinkUserToAccountGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint32))
	})
	return _c
}

func (_c *MockUserRepository_LinkUserToAccountGroup_Call) Return(_a0 error) *MockUserRepository_LinkUserToAccountGroup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_LinkUserToAccountGroup_Call) RunAndReturn(run func(context.Context, uint32, uint32) error) *MockUserRepository_LinkUserToAccountGroup_Call {
	_c.Call.Return(run)
	return _c
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

// MockUserRepository_UpdateDevice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDevice'
type MockUserRepository_UpdateDevice_Call struct {
	*mock.Call
}

// UpdateDevice is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 repositorymodel.UpdateDeviceReq
func (_e *MockUserRepository_Expecter) UpdateDevice(_a0 interface{}, _a1 interface{}) *MockUserRepository_UpdateDevice_Call {
	return &MockUserRepository_UpdateDevice_Call{Call: _e.mock.On("UpdateDevice", _a0, _a1)}
}

func (_c *MockUserRepository_UpdateDevice_Call) Run(run func(_a0 context.Context, _a1 repositorymodel.UpdateDeviceReq)) *MockUserRepository_UpdateDevice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(repositorymodel.UpdateDeviceReq))
	})
	return _c
}

func (_c *MockUserRepository_UpdateDevice_Call) Return(_a0 error) *MockUserRepository_UpdateDevice_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_UpdateDevice_Call) RunAndReturn(run func(context.Context, repositorymodel.UpdateDeviceReq) error) *MockUserRepository_UpdateDevice_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: _a0, _a1
func (_m *MockUserRepository) UpdateUser(_a0 context.Context, _a1 repositorymodel.UpdateUserReq) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repositorymodel.UpdateUserReq) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type MockUserRepository_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 repositorymodel.UpdateUserReq
func (_e *MockUserRepository_Expecter) UpdateUser(_a0 interface{}, _a1 interface{}) *MockUserRepository_UpdateUser_Call {
	return &MockUserRepository_UpdateUser_Call{Call: _e.mock.On("UpdateUser", _a0, _a1)}
}

func (_c *MockUserRepository_UpdateUser_Call) Run(run func(_a0 context.Context, _a1 repositorymodel.UpdateUserReq)) *MockUserRepository_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(repositorymodel.UpdateUserReq))
	})
	return _c
}

func (_c *MockUserRepository_UpdateUser_Call) Return(_a0 error) *MockUserRepository_UpdateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_UpdateUser_Call) RunAndReturn(run func(context.Context, repositorymodel.UpdateUserReq) error) *MockUserRepository_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
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
