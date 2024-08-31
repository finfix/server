// Code generated by mockery v2.42.2. DO NOT EDIT.

package service

import (
	context "context"
	model "server/app/services/tag/model"

	mock "github.com/stretchr/testify/mock"
)

// MockTagRepository is an autogenerated mock type for the TagRepository type
type MockTagRepository struct {
	mock.Mock
}

// GetTagsToTransactions provides a mock function with given fields: _a0, _a1
func (_m *MockTagRepository) GetTagsToTransactions(_a0 context.Context, _a1 model.GetTagsToTransactionsReq) ([]model.TagToTransaction, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetTagsToTransactions")
	}

	var r0 []model.TagToTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetTagsToTransactionsReq) ([]model.TagToTransaction, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetTagsToTransactionsReq) []model.TagToTransaction); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.TagToTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetTagsToTransactionsReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkTagsToTransaction provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockTagRepository) LinkTagsToTransaction(_a0 context.Context, _a1 []uint32, _a2 uint32) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for LinkTagsToTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []uint32, uint32) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnlinkTagsFromTransaction provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockTagRepository) UnlinkTagsFromTransaction(_a0 context.Context, _a1 []uint32, _a2 uint32) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for UnlinkTagsFromTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []uint32, uint32) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockTagRepository creates a new instance of MockTagRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTagRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTagRepository {
	mock := &MockTagRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
