// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	http "net/http"

	fixtureshttp "github.com/vektra/mockery/v2/pkg/fixtures/http"

	mock "github.com/stretchr/testify/mock"
)

// HasConflictingNestedImports is an autogenerated mock type for the HasConflictingNestedImports type
type HasConflictingNestedImports struct {
	mock.Mock
}

// Get provides a mock function with given fields: path
func (_m *HasConflictingNestedImports) Get(path string) (http.Response, error) {
	ret := _m.Called(path)

	var r0 http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (http.Response, error)); ok {
		r0, r1 = rf(path)
	} else {
		if rf, ok := ret.Get(0).(func(string) http.Response); ok {
			r0 = rf(path)
		} else {
			r0 = ret.Get(0).(http.Response)
		}

		if rf, ok := ret.Get(1).(func(string) error); ok {
			r1 = rf(path)
		} else {
			r1 = ret.Error(1)
		}
	}

	return r0, r1
}

// Z provides a mock function with given fields:
func (_m *HasConflictingNestedImports) Z() fixtureshttp.MyStruct {
	ret := _m.Called()

	var r0 fixtureshttp.MyStruct
	if rf, ok := ret.Get(0).(func() fixtureshttp.MyStruct); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(fixtureshttp.MyStruct)
	}

	return r0
}

type mockConstructorTestingTNewHasConflictingNestedImports interface {
	mock.TestingT
	Cleanup(func())
}

// NewHasConflictingNestedImports creates a new instance of HasConflictingNestedImports. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHasConflictingNestedImports(t mockConstructorTestingTNewHasConflictingNestedImports) *HasConflictingNestedImports {
	mock := &HasConflictingNestedImports{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
