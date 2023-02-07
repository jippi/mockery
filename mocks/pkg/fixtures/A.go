// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	test "github.com/vektra/mockery/v2/pkg/fixtures"
)

// A is an autogenerated mock type for the A type
type A struct {
	mock.Mock
}

// Call provides a mock function with given fields:
func (_m *A) Call() (test.B, error) {
	ret := _m.Called()

	var r0 test.B
	var r1 error
	if rf, ok := ret.Get(0).(func() (test.B, error)); ok {
		r0, r1 = rf()
	} else {
		if rf, ok := ret.Get(0).(func() test.B); ok {
			r0 = rf()
		} else {
			r0 = ret.Get(0).(test.B)
		}

		if rf, ok := ret.Get(1).(func() error); ok {
			r1 = rf()
		} else {
			r1 = ret.Error(1)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewA interface {
	mock.TestingT
	Cleanup(func())
}

// NewA creates a new instance of A. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewA(t mockConstructorTestingTNewA) *A {
	mock := &A{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
