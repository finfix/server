// Code generated by mockery v2.42.2. DO NOT EDIT.

package service

import (
	context "context"
	date "server/app/pkg/datetime/date"

	mock "github.com/stretchr/testify/mock"

	model "server/app/services/account/model"
)

// MockAccountRepository is an autogenerated mock type for the AccountRepository type
type MockAccountRepository struct {
	mock.Mock
}

type MockAccountRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccountRepository) EXPECT() *MockAccountRepository_Expecter {
	return &MockAccountRepository_Expecter{mock: &_m.Mock}
}

// CalculateBalancingAmount provides a mock function with given fields: ctx, accountGroupIDs, dateFrom, dateTo
func (_m *MockAccountRepository) CalculateBalancingAmount(ctx context.Context, accountGroupIDs []uint32, dateFrom date.Date, dateTo date.Date) ([]model.BalancingAmount, error) {
	ret := _m.Called(ctx, accountGroupIDs, dateFrom, dateTo)

	if len(ret) == 0 {
		panic("no return value specified for CalculateBalancingAmount")
	}

	var r0 []model.BalancingAmount
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []uint32, date.Date, date.Date) ([]model.BalancingAmount, error)); ok {
		return rf(ctx, accountGroupIDs, dateFrom, dateTo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uint32, date.Date, date.Date) []model.BalancingAmount); ok {
		r0 = rf(ctx, accountGroupIDs, dateFrom, dateTo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.BalancingAmount)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uint32, date.Date, date.Date) error); ok {
		r1 = rf(ctx, accountGroupIDs, dateFrom, dateTo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountRepository_CalculateBalancingAmount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CalculateBalancingAmount'
type MockAccountRepository_CalculateBalancingAmount_Call struct {
	*mock.Call
}

// CalculateBalancingAmount is a helper method to define mock.On call
//   - ctx context.Context
//   - accountGroupIDs []uint32
//   - dateFrom date.Date
//   - dateTo date.Date
func (_e *MockAccountRepository_Expecter) CalculateBalancingAmount(ctx interface{}, accountGroupIDs interface{}, dateFrom interface{}, dateTo interface{}) *MockAccountRepository_CalculateBalancingAmount_Call {
	return &MockAccountRepository_CalculateBalancingAmount_Call{Call: _e.mock.On("CalculateBalancingAmount", ctx, accountGroupIDs, dateFrom, dateTo)}
}

func (_c *MockAccountRepository_CalculateBalancingAmount_Call) Run(run func(ctx context.Context, accountGroupIDs []uint32, dateFrom date.Date, dateTo date.Date)) *MockAccountRepository_CalculateBalancingAmount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]uint32), args[2].(date.Date), args[3].(date.Date))
	})
	return _c
}

func (_c *MockAccountRepository_CalculateBalancingAmount_Call) Return(_a0 []model.BalancingAmount, _a1 error) *MockAccountRepository_CalculateBalancingAmount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountRepository_CalculateBalancingAmount_Call) RunAndReturn(run func(context.Context, []uint32, date.Date, date.Date) ([]model.BalancingAmount, error)) *MockAccountRepository_CalculateBalancingAmount_Call {
	_c.Call.Return(run)
	return _c
}

// CalculateExpensesAndEarnings provides a mock function with given fields: ctx, accountGroupIDs, dateFrom, dateTo
func (_m *MockAccountRepository) CalculateExpensesAndEarnings(ctx context.Context, accountGroupIDs []uint32, dateFrom date.Date, dateTo date.Date) (map[uint32]float64, error) {
	ret := _m.Called(ctx, accountGroupIDs, dateFrom, dateTo)

	if len(ret) == 0 {
		panic("no return value specified for CalculateExpensesAndEarnings")
	}

	var r0 map[uint32]float64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []uint32, date.Date, date.Date) (map[uint32]float64, error)); ok {
		return rf(ctx, accountGroupIDs, dateFrom, dateTo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uint32, date.Date, date.Date) map[uint32]float64); ok {
		r0 = rf(ctx, accountGroupIDs, dateFrom, dateTo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint32]float64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uint32, date.Date, date.Date) error); ok {
		r1 = rf(ctx, accountGroupIDs, dateFrom, dateTo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountRepository_CalculateExpensesAndEarnings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CalculateExpensesAndEarnings'
type MockAccountRepository_CalculateExpensesAndEarnings_Call struct {
	*mock.Call
}

// CalculateExpensesAndEarnings is a helper method to define mock.On call
//   - ctx context.Context
//   - accountGroupIDs []uint32
//   - dateFrom date.Date
//   - dateTo date.Date
func (_e *MockAccountRepository_Expecter) CalculateExpensesAndEarnings(ctx interface{}, accountGroupIDs interface{}, dateFrom interface{}, dateTo interface{}) *MockAccountRepository_CalculateExpensesAndEarnings_Call {
	return &MockAccountRepository_CalculateExpensesAndEarnings_Call{Call: _e.mock.On("CalculateExpensesAndEarnings", ctx, accountGroupIDs, dateFrom, dateTo)}
}

func (_c *MockAccountRepository_CalculateExpensesAndEarnings_Call) Run(run func(ctx context.Context, accountGroupIDs []uint32, dateFrom date.Date, dateTo date.Date)) *MockAccountRepository_CalculateExpensesAndEarnings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]uint32), args[2].(date.Date), args[3].(date.Date))
	})
	return _c
}

func (_c *MockAccountRepository_CalculateExpensesAndEarnings_Call) Return(_a0 map[uint32]float64, _a1 error) *MockAccountRepository_CalculateExpensesAndEarnings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountRepository_CalculateExpensesAndEarnings_Call) RunAndReturn(run func(context.Context, []uint32, date.Date, date.Date) (map[uint32]float64, error)) *MockAccountRepository_CalculateExpensesAndEarnings_Call {
	_c.Call.Return(run)
	return _c
}

// CalculateRemainderAccounts provides a mock function with given fields: ctx, accountGroupIDs, dateTo
func (_m *MockAccountRepository) CalculateRemainderAccounts(ctx context.Context, accountGroupIDs []uint32, dateTo *date.Date) (map[uint32]float64, error) {
	ret := _m.Called(ctx, accountGroupIDs, dateTo)

	if len(ret) == 0 {
		panic("no return value specified for CalculateRemainderAccounts")
	}

	var r0 map[uint32]float64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []uint32, *date.Date) (map[uint32]float64, error)); ok {
		return rf(ctx, accountGroupIDs, dateTo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uint32, *date.Date) map[uint32]float64); ok {
		r0 = rf(ctx, accountGroupIDs, dateTo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint32]float64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uint32, *date.Date) error); ok {
		r1 = rf(ctx, accountGroupIDs, dateTo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountRepository_CalculateRemainderAccounts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CalculateRemainderAccounts'
type MockAccountRepository_CalculateRemainderAccounts_Call struct {
	*mock.Call
}

// CalculateRemainderAccounts is a helper method to define mock.On call
//   - ctx context.Context
//   - accountGroupIDs []uint32
//   - dateTo *date.Date
func (_e *MockAccountRepository_Expecter) CalculateRemainderAccounts(ctx interface{}, accountGroupIDs interface{}, dateTo interface{}) *MockAccountRepository_CalculateRemainderAccounts_Call {
	return &MockAccountRepository_CalculateRemainderAccounts_Call{Call: _e.mock.On("CalculateRemainderAccounts", ctx, accountGroupIDs, dateTo)}
}

func (_c *MockAccountRepository_CalculateRemainderAccounts_Call) Run(run func(ctx context.Context, accountGroupIDs []uint32, dateTo *date.Date)) *MockAccountRepository_CalculateRemainderAccounts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]uint32), args[2].(*date.Date))
	})
	return _c
}

func (_c *MockAccountRepository_CalculateRemainderAccounts_Call) Return(_a0 map[uint32]float64, _a1 error) *MockAccountRepository_CalculateRemainderAccounts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountRepository_CalculateRemainderAccounts_Call) RunAndReturn(run func(context.Context, []uint32, *date.Date) (map[uint32]float64, error)) *MockAccountRepository_CalculateRemainderAccounts_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *MockAccountRepository) Create(_a0 context.Context, _a1 model.CreateReq) (uint32, uint32, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 uint32
	var r1 uint32
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateReq) (uint32, uint32, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateReq) uint32); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CreateReq) uint32); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	if rf, ok := ret.Get(2).(func(context.Context, model.CreateReq) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockAccountRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockAccountRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.CreateReq
func (_e *MockAccountRepository_Expecter) Create(_a0 interface{}, _a1 interface{}) *MockAccountRepository_Create_Call {
	return &MockAccountRepository_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *MockAccountRepository_Create_Call) Run(run func(_a0 context.Context, _a1 model.CreateReq)) *MockAccountRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.CreateReq))
	})
	return _c
}

func (_c *MockAccountRepository_Create_Call) Return(_a0 uint32, _a1 uint32, _a2 error) *MockAccountRepository_Create_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockAccountRepository_Create_Call) RunAndReturn(run func(context.Context, model.CreateReq) (uint32, uint32, error)) *MockAccountRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateAccountGroup provides a mock function with given fields: _a0, _a1
func (_m *MockAccountRepository) CreateAccountGroup(_a0 context.Context, _a1 model.CreateAccountGroupReq) (uint32, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccountGroup")
	}

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateAccountGroupReq) (uint32, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateAccountGroupReq) uint32); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CreateAccountGroupReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountRepository_CreateAccountGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAccountGroup'
type MockAccountRepository_CreateAccountGroup_Call struct {
	*mock.Call
}

// CreateAccountGroup is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.CreateAccountGroupReq
func (_e *MockAccountRepository_Expecter) CreateAccountGroup(_a0 interface{}, _a1 interface{}) *MockAccountRepository_CreateAccountGroup_Call {
	return &MockAccountRepository_CreateAccountGroup_Call{Call: _e.mock.On("CreateAccountGroup", _a0, _a1)}
}

func (_c *MockAccountRepository_CreateAccountGroup_Call) Run(run func(_a0 context.Context, _a1 model.CreateAccountGroupReq)) *MockAccountRepository_CreateAccountGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.CreateAccountGroupReq))
	})
	return _c
}

func (_c *MockAccountRepository_CreateAccountGroup_Call) Return(_a0 uint32, _a1 error) *MockAccountRepository_CreateAccountGroup_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountRepository_CreateAccountGroup_Call) RunAndReturn(run func(context.Context, model.CreateAccountGroupReq) (uint32, error)) *MockAccountRepository_CreateAccountGroup_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockAccountRepository) Delete(ctx context.Context, id uint32) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockAccountRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
func (_e *MockAccountRepository_Expecter) Delete(ctx interface{}, id interface{}) *MockAccountRepository_Delete_Call {
	return &MockAccountRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *MockAccountRepository_Delete_Call) Run(run func(ctx context.Context, id uint32)) *MockAccountRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *MockAccountRepository_Delete_Call) Return(_a0 error) *MockAccountRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountRepository_Delete_Call) RunAndReturn(run func(context.Context, uint32) error) *MockAccountRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *MockAccountRepository) Get(_a0 context.Context, _a1 model.GetReq) ([]model.Account, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []model.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetReq) ([]model.Account, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetReq) []model.Account); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockAccountRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.GetReq
func (_e *MockAccountRepository_Expecter) Get(_a0 interface{}, _a1 interface{}) *MockAccountRepository_Get_Call {
	return &MockAccountRepository_Get_Call{Call: _e.mock.On("Get", _a0, _a1)}
}

func (_c *MockAccountRepository_Get_Call) Run(run func(_a0 context.Context, _a1 model.GetReq)) *MockAccountRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.GetReq))
	})
	return _c
}

func (_c *MockAccountRepository_Get_Call) Return(_a0 []model.Account, _a1 error) *MockAccountRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountRepository_Get_Call) RunAndReturn(run func(context.Context, model.GetReq) ([]model.Account, error)) *MockAccountRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAccountGroups provides a mock function with given fields: _a0, _a1
func (_m *MockAccountRepository) GetAccountGroups(_a0 context.Context, _a1 model.GetAccountGroupsReq) ([]model.AccountGroup, error) {
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

// MockAccountRepository_GetAccountGroups_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccountGroups'
type MockAccountRepository_GetAccountGroups_Call struct {
	*mock.Call
}

// GetAccountGroups is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.GetAccountGroupsReq
func (_e *MockAccountRepository_Expecter) GetAccountGroups(_a0 interface{}, _a1 interface{}) *MockAccountRepository_GetAccountGroups_Call {
	return &MockAccountRepository_GetAccountGroups_Call{Call: _e.mock.On("GetAccountGroups", _a0, _a1)}
}

func (_c *MockAccountRepository_GetAccountGroups_Call) Run(run func(_a0 context.Context, _a1 model.GetAccountGroupsReq)) *MockAccountRepository_GetAccountGroups_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.GetAccountGroupsReq))
	})
	return _c
}

func (_c *MockAccountRepository_GetAccountGroups_Call) Return(_a0 []model.AccountGroup, _a1 error) *MockAccountRepository_GetAccountGroups_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountRepository_GetAccountGroups_Call) RunAndReturn(run func(context.Context, model.GetAccountGroupsReq) ([]model.AccountGroup, error)) *MockAccountRepository_GetAccountGroups_Call {
	_c.Call.Return(run)
	return _c
}

// GetRemainder provides a mock function with given fields: ctx, id
func (_m *MockAccountRepository) GetRemainder(ctx context.Context, id uint32) (float64, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetRemainder")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) (float64, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) float64); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountRepository_GetRemainder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRemainder'
type MockAccountRepository_GetRemainder_Call struct {
	*mock.Call
}

// GetRemainder is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
func (_e *MockAccountRepository_Expecter) GetRemainder(ctx interface{}, id interface{}) *MockAccountRepository_GetRemainder_Call {
	return &MockAccountRepository_GetRemainder_Call{Call: _e.mock.On("GetRemainder", ctx, id)}
}

func (_c *MockAccountRepository_GetRemainder_Call) Run(run func(ctx context.Context, id uint32)) *MockAccountRepository_GetRemainder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *MockAccountRepository_GetRemainder_Call) Return(_a0 float64, _a1 error) *MockAccountRepository_GetRemainder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountRepository_GetRemainder_Call) RunAndReturn(run func(context.Context, uint32) (float64, error)) *MockAccountRepository_GetRemainder_Call {
	_c.Call.Return(run)
	return _c
}

// Switch provides a mock function with given fields: ctx, id1, id2
func (_m *MockAccountRepository) Switch(ctx context.Context, id1 uint32, id2 uint32) error {
	ret := _m.Called(ctx, id1, id2)

	if len(ret) == 0 {
		panic("no return value specified for Switch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32) error); ok {
		r0 = rf(ctx, id1, id2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountRepository_Switch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Switch'
type MockAccountRepository_Switch_Call struct {
	*mock.Call
}

// Switch is a helper method to define mock.On call
//   - ctx context.Context
//   - id1 uint32
//   - id2 uint32
func (_e *MockAccountRepository_Expecter) Switch(ctx interface{}, id1 interface{}, id2 interface{}) *MockAccountRepository_Switch_Call {
	return &MockAccountRepository_Switch_Call{Call: _e.mock.On("Switch", ctx, id1, id2)}
}

func (_c *MockAccountRepository_Switch_Call) Run(run func(ctx context.Context, id1 uint32, id2 uint32)) *MockAccountRepository_Switch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint32))
	})
	return _c
}

func (_c *MockAccountRepository_Switch_Call) Return(_a0 error) *MockAccountRepository_Switch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountRepository_Switch_Call) RunAndReturn(run func(context.Context, uint32, uint32) error) *MockAccountRepository_Switch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *MockAccountRepository) Update(_a0 context.Context, _a1 model.UpdateReq) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateReq) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockAccountRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.UpdateReq
func (_e *MockAccountRepository_Expecter) Update(_a0 interface{}, _a1 interface{}) *MockAccountRepository_Update_Call {
	return &MockAccountRepository_Update_Call{Call: _e.mock.On("Update", _a0, _a1)}
}

func (_c *MockAccountRepository_Update_Call) Run(run func(_a0 context.Context, _a1 model.UpdateReq)) *MockAccountRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UpdateReq))
	})
	return _c
}

func (_c *MockAccountRepository_Update_Call) Return(_a0 error) *MockAccountRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountRepository_Update_Call) RunAndReturn(run func(context.Context, model.UpdateReq) error) *MockAccountRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAccountRepository creates a new instance of MockAccountRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAccountRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountRepository {
	mock := &MockAccountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
