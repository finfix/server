// Code generated by mockery v2.46.2. DO NOT EDIT.

package endpoint

import (
	context "context"
	model "server/internal/services/accountGroup/model"

	mock "github.com/stretchr/testify/mock"
)

// mockAccountGroupService is an autogenerated mock type for the accountGroupService type
type mockAccountGroupService struct {
	mock.Mock
}

// CreateAccountGroup provides a mock function with given fields: _a0, _a1
func (_m *mockAccountGroupService) CreateAccountGroup(_a0 context.Context, _a1 model.CreateAccountGroupReq) (model.CreateAccountGroupRes, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccountGroup")
	}

	var r0 model.CreateAccountGroupRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateAccountGroupReq) (model.CreateAccountGroupRes, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateAccountGroupReq) model.CreateAccountGroupRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(model.CreateAccountGroupRes)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CreateAccountGroupReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAccountGroup provides a mock function with given fields: _a0, _a1
func (_m *mockAccountGroupService) DeleteAccountGroup(_a0 context.Context, _a1 model.DeleteAccountGroupReq) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAccountGroup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.DeleteAccountGroupReq) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAccountGroups provides a mock function with given fields: _a0, _a1
func (_m *mockAccountGroupService) GetAccountGroups(_a0 context.Context, _a1 model.GetAccountGroupsReq) ([]model.AccountGroup, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAccountGroups")
	}

	var r0 []model.AccountGroup
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetAccountGroupsReq) ([]model.AccountGroup, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetAccountGroupsReq) []model.AccountGroup); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.AccountGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetAccountGroupsReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAccountGroup provides a mock function with given fields: _a0, _a1
func (_m *mockAccountGroupService) UpdateAccountGroup(_a0 context.Context, _a1 model.UpdateAccountGroupReq) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAccountGroup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateAccountGroupReq) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockAccountGroupService creates a new instance of mockAccountGroupService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockAccountGroupService(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockAccountGroupService {
	mock := &mockAccountGroupService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
