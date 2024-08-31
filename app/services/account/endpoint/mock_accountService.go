// Code generated by mockery v2.42.2. DO NOT EDIT.

package endpoint

import (
	context "context"
	model "server/app/services/account/model"

	mock "github.com/stretchr/testify/mock"
)

// mockAccountService is an autogenerated mock type for the accountService type
type mockAccountService struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: _a0, _a1
func (_m *mockAccountService) CreateAccount(_a0 context.Context, _a1 model.CreateAccountReq) (model.CreateAccountRes, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccount")
	}

	var r0 model.CreateAccountRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateAccountReq) (model.CreateAccountRes, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateAccountReq) model.CreateAccountRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(model.CreateAccountRes)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CreateAccountReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAccount provides a mock function with given fields: _a0, _a1
func (_m *mockAccountService) DeleteAccount(_a0 context.Context, _a1 model.DeleteAccountReq) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAccount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.DeleteAccountReq) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAccounts provides a mock function with given fields: _a0, _a1
func (_m *mockAccountService) GetAccounts(_a0 context.Context, _a1 model.GetAccountsReq) ([]model.Account, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAccounts")
	}

	var r0 []model.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetAccountsReq) ([]model.Account, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetAccountsReq) []model.Account); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetAccountsReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAccount provides a mock function with given fields: _a0, _a1
func (_m *mockAccountService) UpdateAccount(_a0 context.Context, _a1 model.UpdateAccountReq) (model.UpdateAccountRes, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAccount")
	}

	var r0 model.UpdateAccountRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateAccountReq) (model.UpdateAccountRes, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateAccountReq) model.UpdateAccountRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(model.UpdateAccountRes)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.UpdateAccountReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockAccountService creates a new instance of mockAccountService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockAccountService(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockAccountService {
	mock := &mockAccountService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
