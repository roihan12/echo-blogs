// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	category "echo-blog/bussiness/category"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: categoryDomain
func (_m *Usecase) Create(categoryDomain *category.Domain) category.Domain {
	ret := _m.Called(categoryDomain)

	var r0 category.Domain
	if rf, ok := ret.Get(0).(func(*category.Domain) category.Domain); ok {
		r0 = rf(categoryDomain)
	} else {
		r0 = ret.Get(0).(category.Domain)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *Usecase) Delete(id string) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Usecase) GetAll() []category.Domain {
	ret := _m.Called()

	var r0 []category.Domain
	if rf, ok := ret.Get(0).(func() []category.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]category.Domain)
		}
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *Usecase) GetByID(id string) category.Domain {
	ret := _m.Called(id)

	var r0 category.Domain
	if rf, ok := ret.Get(0).(func(string) category.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(category.Domain)
	}

	return r0
}

// Update provides a mock function with given fields: id, categoryDomain
func (_m *Usecase) Update(id string, categoryDomain *category.Domain) category.Domain {
	ret := _m.Called(id, categoryDomain)

	var r0 category.Domain
	if rf, ok := ret.Get(0).(func(string, *category.Domain) category.Domain); ok {
		r0 = rf(id, categoryDomain)
	} else {
		r0 = ret.Get(0).(category.Domain)
	}

	return r0
}

type mockConstructorTestingTNewUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsecase creates a new instance of Usecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsecase(t mockConstructorTestingTNewUsecase) *Usecase {
	mock := &Usecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}