// Code generated by mockery v2.42.1. DO NOT EDIT.

package repository_mocks

import (
	model "self-projects/pismo/internal/app/pismo/model"

	mock "github.com/stretchr/testify/mock"
)

// ITransactionRepo is an autogenerated mock type for the ITransactionRepo type
type ITransactionRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, tableMame
func (_m *ITransactionRepo) Create(_a0 interface{}, tableMame string) error {
	ret := _m.Called(_a0, tableMame)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string) error); ok {
		r0 = rf(_a0, tableMame)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchAllTransactionsWithAmountLessThan provides a mock function with given fields: accountId, amount
func (_m *ITransactionRepo) FetchAllTransactionsWithAmountLessThan(accountId int, amount float64) ([]model.Transaction, error) {
	ret := _m.Called(accountId, amount)

	if len(ret) == 0 {
		panic("no return value specified for FetchAllTransactionsWithAmountLessThan")
	}

	var r0 []model.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(int, float64) ([]model.Transaction, error)); ok {
		return rf(accountId, amount)
	}
	if rf, ok := ret.Get(0).(func(int, float64) []model.Transaction); ok {
		r0 = rf(accountId, amount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(int, float64) error); ok {
		r1 = rf(accountId, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *ITransactionRepo) Update(_a0 interface{}) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewITransactionRepo creates a new instance of ITransactionRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewITransactionRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *ITransactionRepo {
	mock := &ITransactionRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
