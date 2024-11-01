// Code generated by mockery v2.46.2. DO NOT EDIT.

package service

import (
	context "context"
	model "server/internal/services/tgBot/model"

	mock "github.com/stretchr/testify/mock"
)

// MockTgBotService is an autogenerated mock type for the TgBotService type
type MockTgBotService struct {
	mock.Mock
}

// SendMessage provides a mock function with given fields: _a0, _a1
func (_m *MockTgBotService) SendMessage(_a0 context.Context, _a1 model.SendMessageReq) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SendMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.SendMessageReq) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockTgBotService creates a new instance of MockTgBotService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTgBotService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTgBotService {
	mock := &MockTgBotService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
