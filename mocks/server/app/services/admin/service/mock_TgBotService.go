// Code generated by mockery v2.42.2. DO NOT EDIT.

package service

import (
	context "context"
	model "server/app/services/tgBot/model"

	mock "github.com/stretchr/testify/mock"
)

// MockTgBotService is an autogenerated mock type for the TgBotService type
type MockTgBotService struct {
	mock.Mock
}

type MockTgBotService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTgBotService) EXPECT() *MockTgBotService_Expecter {
	return &MockTgBotService_Expecter{mock: &_m.Mock}
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

// MockTgBotService_SendMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMessage'
type MockTgBotService_SendMessage_Call struct {
	*mock.Call
}

// SendMessage is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.SendMessageReq
func (_e *MockTgBotService_Expecter) SendMessage(_a0 interface{}, _a1 interface{}) *MockTgBotService_SendMessage_Call {
	return &MockTgBotService_SendMessage_Call{Call: _e.mock.On("SendMessage", _a0, _a1)}
}

func (_c *MockTgBotService_SendMessage_Call) Run(run func(_a0 context.Context, _a1 model.SendMessageReq)) *MockTgBotService_SendMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SendMessageReq))
	})
	return _c
}

func (_c *MockTgBotService_SendMessage_Call) Return(_a0 error) *MockTgBotService_SendMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTgBotService_SendMessage_Call) RunAndReturn(run func(context.Context, model.SendMessageReq) error) *MockTgBotService_SendMessage_Call {
	_c.Call.Return(run)
	return _c
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
