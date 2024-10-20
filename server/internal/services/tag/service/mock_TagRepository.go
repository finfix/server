// Code generated by mockery v2.46.2. DO NOT EDIT.

package service

import (
	context "context"
	model "server/internal/services/tag/repository/model"

	mock "github.com/stretchr/testify/mock"

	tagmodel "server/internal/services/tag/model"
)

// MockTagRepository is an autogenerated mock type for the TagRepository type
type MockTagRepository struct {
	mock.Mock
}

// CheckAccess provides a mock function with given fields: ctx, accountGroupIDs, tagIDs
func (_m *MockTagRepository) CheckAccess(ctx context.Context, accountGroupIDs []uint32, tagIDs []uint32) error {
	ret := _m.Called(ctx, accountGroupIDs, tagIDs)

	if len(ret) == 0 {
		panic("no return value specified for CheckAccess")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []uint32, []uint32) error); ok {
		r0 = rf(ctx, accountGroupIDs, tagIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTag provides a mock function with given fields: _a0, _a1
func (_m *MockTagRepository) CreateTag(_a0 context.Context, _a1 model.CreateTagReq) (uint32, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateTag")
	}

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateTagReq) (uint32, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateTagReq) uint32); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CreateTagReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTag provides a mock function with given fields: ctx, id, userID
func (_m *MockTagRepository) DeleteTag(ctx context.Context, id uint32, userID uint32) error {
	ret := _m.Called(ctx, id, userID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32) error); ok {
		r0 = rf(ctx, id, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTags provides a mock function with given fields: _a0, _a1
func (_m *MockTagRepository) GetTags(_a0 context.Context, _a1 tagmodel.GetTagsReq) ([]tagmodel.Tag, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetTags")
	}

	var r0 []tagmodel.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, tagmodel.GetTagsReq) ([]tagmodel.Tag, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, tagmodel.GetTagsReq) []tagmodel.Tag); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]tagmodel.Tag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, tagmodel.GetTagsReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTagsToTransactions provides a mock function with given fields: ctx, req
func (_m *MockTagRepository) GetTagsToTransactions(ctx context.Context, req tagmodel.GetTagsToTransactionsReq) ([]tagmodel.TagToTransaction, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetTagsToTransactions")
	}

	var r0 []tagmodel.TagToTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, tagmodel.GetTagsToTransactionsReq) ([]tagmodel.TagToTransaction, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, tagmodel.GetTagsToTransactionsReq) []tagmodel.TagToTransaction); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]tagmodel.TagToTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, tagmodel.GetTagsToTransactionsReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTag provides a mock function with given fields: _a0, _a1
func (_m *MockTagRepository) UpdateTag(_a0 context.Context, _a1 tagmodel.UpdateTagReq) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, tagmodel.UpdateTagReq) error); ok {
		r0 = rf(_a0, _a1)
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
