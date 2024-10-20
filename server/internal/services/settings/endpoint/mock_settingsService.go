// Code generated by mockery v2.46.2. DO NOT EDIT.

package endpoint

import (
	context "context"
	applicationType "server/internal/services/settings/model/applicationType"

	mock "github.com/stretchr/testify/mock"

	model "server/internal/services/settings/model"
)

// mockSettingsService is an autogenerated mock type for the settingsService type
type mockSettingsService struct {
	mock.Mock
}

// GetCurrencies provides a mock function with given fields: _a0
func (_m *mockSettingsService) GetCurrencies(_a0 context.Context) ([]model.Currency, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetCurrencies")
	}

	var r0 []model.Currency
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.Currency, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.Currency); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Currency)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIcons provides a mock function with given fields: _a0
func (_m *mockSettingsService) GetIcons(_a0 context.Context) ([]model.Icon, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetIcons")
	}

	var r0 []model.Icon
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.Icon, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.Icon); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Icon)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVersion provides a mock function with given fields: _a0, _a1
func (_m *mockSettingsService) GetVersion(_a0 context.Context, _a1 applicationType.Type) (model.Version, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetVersion")
	}

	var r0 model.Version
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, applicationType.Type) (model.Version, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, applicationType.Type) model.Version); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(model.Version)
	}

	if rf, ok := ret.Get(1).(func(context.Context, applicationType.Type) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendNotification provides a mock function with given fields: _a0, _a1
func (_m *mockSettingsService) SendNotification(_a0 context.Context, _a1 model.SendNotificationReq) (model.SendNotificationRes, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SendNotification")
	}

	var r0 model.SendNotificationRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.SendNotificationReq) (model.SendNotificationRes, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.SendNotificationReq) model.SendNotificationRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(model.SendNotificationRes)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.SendNotificationReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCurrencies provides a mock function with given fields: _a0, _a1
func (_m *mockSettingsService) UpdateCurrencies(_a0 context.Context, _a1 model.UpdateCurrenciesReq) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCurrencies")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateCurrenciesReq) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockSettingsService creates a new instance of mockSettingsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockSettingsService(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockSettingsService {
	mock := &mockSettingsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
