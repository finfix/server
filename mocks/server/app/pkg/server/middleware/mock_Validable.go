// Code generated by mockery v2.42.2. DO NOT EDIT.

package middleware

import mock "github.com/stretchr/testify/mock"

// MockValidable is an autogenerated mock type for the Validable type
type MockValidable struct {
	mock.Mock
}

type MockValidable_Expecter struct {
	mock *mock.Mock
}

func (_m *MockValidable) EXPECT() *MockValidable_Expecter {
	return &MockValidable_Expecter{mock: &_m.Mock}
}

// Validate provides a mock function with given fields:
func (_m *MockValidable) Validate() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockValidable_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type MockValidable_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
func (_e *MockValidable_Expecter) Validate() *MockValidable_Validate_Call {
	return &MockValidable_Validate_Call{Call: _e.mock.On("Validate")}
}

func (_c *MockValidable_Validate_Call) Run(run func()) *MockValidable_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockValidable_Validate_Call) Return(_a0 error) *MockValidable_Validate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockValidable_Validate_Call) RunAndReturn(run func() error) *MockValidable_Validate_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockValidable creates a new instance of MockValidable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockValidable(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockValidable {
	mock := &MockValidable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}