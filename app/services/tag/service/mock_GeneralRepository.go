// Code generated by mockery v2.42.2. DO NOT EDIT.

package service

import (
	context "context"
	checker "server/app/services/generalRepository/checker"

	mock "github.com/stretchr/testify/mock"
)

// MockGeneralRepository is an autogenerated mock type for the GeneralRepository type
type MockGeneralRepository struct {
	mock.Mock
}

// CheckUserAccessToObjects provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *MockGeneralRepository) CheckUserAccessToObjects(_a0 context.Context, _a1 checker.CheckType, _a2 uint32, _a3 []uint32) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for CheckUserAccessToObjects")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, checker.CheckType, uint32, []uint32) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAvailableAccountGroups provides a mock function with given fields: userID
func (_m *MockGeneralRepository) GetAvailableAccountGroups(userID uint32) []uint32 {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetAvailableAccountGroups")
	}

	var r0 []uint32
	if rf, ok := ret.Get(0).(func(uint32) []uint32); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint32)
		}
	}

	return r0
}

// WithinTransaction provides a mock function with given fields: ctx, callback
func (_m *MockGeneralRepository) WithinTransaction(ctx context.Context, callback func(context.Context) error) error {
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

// NewMockGeneralRepository creates a new instance of MockGeneralRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGeneralRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGeneralRepository {
	mock := &MockGeneralRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
