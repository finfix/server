// Code generated by mockery v2.42.2. DO NOT EDIT.

package endpoint

import (
	context "context"
	model "server/internal/services/tag/model"

	mock "github.com/stretchr/testify/mock"
)

// mockTagService is an autogenerated mock type for the tagService type
type mockTagService struct {
	mock.Mock
}

// CreateTag provides a mock function with given fields: _a0, _a1
func (_m *mockTagService) CreateTag(_a0 context.Context, _a1 model.CreateTagReq) (uint32, error) {
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

// DeleteTag provides a mock function with given fields: _a0, _a1
func (_m *mockTagService) DeleteTag(_a0 context.Context, _a1 model.DeleteTagReq) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.DeleteTagReq) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTags provides a mock function with given fields: _a0, _a1
func (_m *mockTagService) GetTags(_a0 context.Context, _a1 model.GetTagsReq) ([]model.Tag, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetTags")
	}

	var r0 []model.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetTagsReq) ([]model.Tag, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetTagsReq) []model.Tag); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Tag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetTagsReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTagsToTransactions provides a mock function with given fields: ctx, req
func (_m *mockTagService) GetTagsToTransactions(ctx context.Context, req model.GetTagsToTransactionsReq) ([]model.TagToTransaction, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetTagsToTransactions")
	}

	var r0 []model.TagToTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetTagsToTransactionsReq) ([]model.TagToTransaction, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetTagsToTransactionsReq) []model.TagToTransaction); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.TagToTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetTagsToTransactionsReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTag provides a mock function with given fields: _a0, _a1
func (_m *mockTagService) UpdateTag(_a0 context.Context, _a1 model.UpdateTagReq) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateTagReq) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockTagService creates a new instance of mockTagService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockTagService(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockTagService {
	mock := &mockTagService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
