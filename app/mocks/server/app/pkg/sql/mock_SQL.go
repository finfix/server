// Code generated by mockery v2.42.2. DO NOT EDIT.

package sql

import (
	context "context"
	sql "server/app/pkg/sql"

	mock "github.com/stretchr/testify/mock"
)

// MockSQL is an autogenerated mock type for the SQL type
type MockSQL struct {
	mock.Mock
}

type MockSQL_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSQL) EXPECT() *MockSQL_Expecter {
	return &MockSQL_Expecter{mock: &_m.Mock}
}

// Begin provides a mock function with given fields: _a0
func (_m *MockSQL) Begin(_a0 context.Context) (*sql.Tx, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Begin")
	}

	var r0 *sql.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*sql.Tx, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *sql.Tx); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSQL_Begin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Begin'
type MockSQL_Begin_Call struct {
	*mock.Call
}

// Begin is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockSQL_Expecter) Begin(_a0 interface{}) *MockSQL_Begin_Call {
	return &MockSQL_Begin_Call{Call: _e.mock.On("Begin", _a0)}
}

func (_c *MockSQL_Begin_Call) Run(run func(_a0 context.Context)) *MockSQL_Begin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSQL_Begin_Call) Return(_a0 *sql.Tx, _a1 error) *MockSQL_Begin_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSQL_Begin_Call) RunAndReturn(run func(context.Context) (*sql.Tx, error)) *MockSQL_Begin_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockSQL) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSQL_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockSQL_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockSQL_Expecter) Close() *MockSQL_Close_Call {
	return &MockSQL_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockSQL_Close_Call) Run(run func()) *MockSQL_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSQL_Close_Call) Return(_a0 error) *MockSQL_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSQL_Close_Call) RunAndReturn(run func() error) *MockSQL_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Exec provides a mock function with given fields: ctx, query, args
func (_m *MockSQL) Exec(ctx context.Context, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) error); ok {
		r0 = rf(ctx, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSQL_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockSQL_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockSQL_Expecter) Exec(ctx interface{}, query interface{}, args ...interface{}) *MockSQL_Exec_Call {
	return &MockSQL_Exec_Call{Call: _e.mock.On("Exec",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockSQL_Exec_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockSQL_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSQL_Exec_Call) Return(_a0 error) *MockSQL_Exec_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSQL_Exec_Call) RunAndReturn(run func(context.Context, string, ...interface{}) error) *MockSQL_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// ExecWithLastInsertID provides a mock function with given fields: ctx, query, args
func (_m *MockSQL) ExecWithLastInsertID(ctx context.Context, query string, args ...interface{}) (uint32, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExecWithLastInsertID")
	}

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (uint32, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) uint32); ok {
		r0 = rf(ctx, query, args...)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSQL_ExecWithLastInsertID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecWithLastInsertID'
type MockSQL_ExecWithLastInsertID_Call struct {
	*mock.Call
}

// ExecWithLastInsertID is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockSQL_Expecter) ExecWithLastInsertID(ctx interface{}, query interface{}, args ...interface{}) *MockSQL_ExecWithLastInsertID_Call {
	return &MockSQL_ExecWithLastInsertID_Call{Call: _e.mock.On("ExecWithLastInsertID",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockSQL_ExecWithLastInsertID_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockSQL_ExecWithLastInsertID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSQL_ExecWithLastInsertID_Call) Return(_a0 uint32, _a1 error) *MockSQL_ExecWithLastInsertID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSQL_ExecWithLastInsertID_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (uint32, error)) *MockSQL_ExecWithLastInsertID_Call {
	_c.Call.Return(run)
	return _c
}

// ExecWithRowsAffected provides a mock function with given fields: ctx, query, args
func (_m *MockSQL) ExecWithRowsAffected(ctx context.Context, query string, args ...interface{}) (uint32, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExecWithRowsAffected")
	}

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (uint32, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) uint32); ok {
		r0 = rf(ctx, query, args...)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSQL_ExecWithRowsAffected_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecWithRowsAffected'
type MockSQL_ExecWithRowsAffected_Call struct {
	*mock.Call
}

// ExecWithRowsAffected is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockSQL_Expecter) ExecWithRowsAffected(ctx interface{}, query interface{}, args ...interface{}) *MockSQL_ExecWithRowsAffected_Call {
	return &MockSQL_ExecWithRowsAffected_Call{Call: _e.mock.On("ExecWithRowsAffected",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockSQL_ExecWithRowsAffected_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockSQL_ExecWithRowsAffected_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSQL_ExecWithRowsAffected_Call) Return(_a0 uint32, _a1 error) *MockSQL_ExecWithRowsAffected_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSQL_ExecWithRowsAffected_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (uint32, error)) *MockSQL_ExecWithRowsAffected_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, dest, query, args
func (_m *MockSQL) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSQL_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockSQL_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - dest interface{}
//   - query string
//   - args ...interface{}
func (_e *MockSQL_Expecter) Get(ctx interface{}, dest interface{}, query interface{}, args ...interface{}) *MockSQL_Get_Call {
	return &MockSQL_Get_Call{Call: _e.mock.On("Get",
		append([]interface{}{ctx, dest, query}, args...)...)}
}

func (_c *MockSQL_Get_Call) Run(run func(ctx context.Context, dest interface{}, query string, args ...interface{})) *MockSQL_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(interface{}), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSQL_Get_Call) Return(_a0 error) *MockSQL_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSQL_Get_Call) RunAndReturn(run func(context.Context, interface{}, string, ...interface{}) error) *MockSQL_Get_Call {
	_c.Call.Return(run)
	return _c
}

// In provides a mock function with given fields: query, args
func (_m *MockSQL) In(query string, args ...interface{}) (string, []interface{}, error) {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for In")
	}

	var r0 string
	var r1 []interface{}
	var r2 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) (string, []interface{}, error)); ok {
		return rf(query, args...)
	}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) string); ok {
		r0 = rf(query, args...)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, ...interface{}) []interface{}); ok {
		r1 = rf(query, args...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]interface{})
		}
	}

	if rf, ok := ret.Get(2).(func(string, ...interface{}) error); ok {
		r2 = rf(query, args...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockSQL_In_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'In'
type MockSQL_In_Call struct {
	*mock.Call
}

// In is a helper method to define mock.On call
//   - query string
//   - args ...interface{}
func (_e *MockSQL_Expecter) In(query interface{}, args ...interface{}) *MockSQL_In_Call {
	return &MockSQL_In_Call{Call: _e.mock.On("In",
		append([]interface{}{query}, args...)...)}
}

func (_c *MockSQL_In_Call) Run(run func(query string, args ...interface{})) *MockSQL_In_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSQL_In_Call) Return(_a0 string, _a1 []interface{}, _a2 error) *MockSQL_In_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockSQL_In_Call) RunAndReturn(run func(string, ...interface{}) (string, []interface{}, error)) *MockSQL_In_Call {
	_c.Call.Return(run)
	return _c
}

// Ping provides a mock function with given fields:
func (_m *MockSQL) Ping() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ping")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSQL_Ping_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ping'
type MockSQL_Ping_Call struct {
	*mock.Call
}

// Ping is a helper method to define mock.On call
func (_e *MockSQL_Expecter) Ping() *MockSQL_Ping_Call {
	return &MockSQL_Ping_Call{Call: _e.mock.On("Ping")}
}

func (_c *MockSQL_Ping_Call) Run(run func()) *MockSQL_Ping_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSQL_Ping_Call) Return(_a0 error) *MockSQL_Ping_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSQL_Ping_Call) RunAndReturn(run func() error) *MockSQL_Ping_Call {
	_c.Call.Return(run)
	return _c
}

// Prepare provides a mock function with given fields: ctx, query
func (_m *MockSQL) Prepare(ctx context.Context, query string) (*sql.Stmt, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for Prepare")
	}

	var r0 *sql.Stmt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sql.Stmt, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sql.Stmt); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Stmt)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSQL_Prepare_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Prepare'
type MockSQL_Prepare_Call struct {
	*mock.Call
}

// Prepare is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
func (_e *MockSQL_Expecter) Prepare(ctx interface{}, query interface{}) *MockSQL_Prepare_Call {
	return &MockSQL_Prepare_Call{Call: _e.mock.On("Prepare", ctx, query)}
}

func (_c *MockSQL_Prepare_Call) Run(run func(ctx context.Context, query string)) *MockSQL_Prepare_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSQL_Prepare_Call) Return(_a0 *sql.Stmt, _a1 error) *MockSQL_Prepare_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSQL_Prepare_Call) RunAndReturn(run func(context.Context, string) (*sql.Stmt, error)) *MockSQL_Prepare_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: ctx, query, args
func (_m *MockSQL) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 *sql.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (*sql.Rows, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sql.Rows); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSQL_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type MockSQL_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockSQL_Expecter) Query(ctx interface{}, query interface{}, args ...interface{}) *MockSQL_Query_Call {
	return &MockSQL_Query_Call{Call: _e.mock.On("Query",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockSQL_Query_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockSQL_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSQL_Query_Call) Return(_a0 *sql.Rows, _a1 error) *MockSQL_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSQL_Query_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (*sql.Rows, error)) *MockSQL_Query_Call {
	_c.Call.Return(run)
	return _c
}

// QueryRow provides a mock function with given fields: ctx, query, args
func (_m *MockSQL) QueryRow(ctx context.Context, query string, args ...interface{}) (*sql.Row, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryRow")
	}

	var r0 *sql.Row
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (*sql.Row, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sql.Row); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSQL_QueryRow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryRow'
type MockSQL_QueryRow_Call struct {
	*mock.Call
}

// QueryRow is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockSQL_Expecter) QueryRow(ctx interface{}, query interface{}, args ...interface{}) *MockSQL_QueryRow_Call {
	return &MockSQL_QueryRow_Call{Call: _e.mock.On("QueryRow",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockSQL_QueryRow_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockSQL_QueryRow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSQL_QueryRow_Call) Return(_a0 *sql.Row, _a1 error) *MockSQL_QueryRow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSQL_QueryRow_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (*sql.Row, error)) *MockSQL_QueryRow_Call {
	_c.Call.Return(run)
	return _c
}

// Select provides a mock function with given fields: ctx, dest, query, args
func (_m *MockSQL) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Select")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSQL_Select_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Select'
type MockSQL_Select_Call struct {
	*mock.Call
}

// Select is a helper method to define mock.On call
//   - ctx context.Context
//   - dest interface{}
//   - query string
//   - args ...interface{}
func (_e *MockSQL_Expecter) Select(ctx interface{}, dest interface{}, query interface{}, args ...interface{}) *MockSQL_Select_Call {
	return &MockSQL_Select_Call{Call: _e.mock.On("Select",
		append([]interface{}{ctx, dest, query}, args...)...)}
}

func (_c *MockSQL_Select_Call) Run(run func(ctx context.Context, dest interface{}, query string, args ...interface{})) *MockSQL_Select_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(interface{}), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSQL_Select_Call) Return(_a0 error) *MockSQL_Select_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSQL_Select_Call) RunAndReturn(run func(context.Context, interface{}, string, ...interface{}) error) *MockSQL_Select_Call {
	_c.Call.Return(run)
	return _c
}

// Unsafe provides a mock function with given fields:
func (_m *MockSQL) Unsafe() *sql.DB {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Unsafe")
	}

	var r0 *sql.DB
	if rf, ok := ret.Get(0).(func() *sql.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
		}
	}

	return r0
}

// MockSQL_Unsafe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unsafe'
type MockSQL_Unsafe_Call struct {
	*mock.Call
}

// Unsafe is a helper method to define mock.On call
func (_e *MockSQL_Expecter) Unsafe() *MockSQL_Unsafe_Call {
	return &MockSQL_Unsafe_Call{Call: _e.mock.On("Unsafe")}
}

func (_c *MockSQL_Unsafe_Call) Run(run func()) *MockSQL_Unsafe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSQL_Unsafe_Call) Return(_a0 *sql.DB) *MockSQL_Unsafe_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSQL_Unsafe_Call) RunAndReturn(run func() *sql.DB) *MockSQL_Unsafe_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSQL creates a new instance of MockSQL. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSQL(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSQL {
	mock := &MockSQL{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}