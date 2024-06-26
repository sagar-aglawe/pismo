// Code generated by mockery v2.42.1. DO NOT EDIT.

package service_mocks

import (
	dto "self-projects/pismo/internal/app/pismo/dto"

	mock "github.com/stretchr/testify/mock"

	request_context "self-projects/pismo/pkg/request_context"
)

// IAccountService is an autogenerated mock type for the IAccountService type
type IAccountService struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: rCtx, reqBody
func (_m *IAccountService) CreateAccount(rCtx *request_context.ReqCtx, reqBody *dto.AccountCreateRequest) (*dto.AccountCreateResponse, error) {
	ret := _m.Called(rCtx, reqBody)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccount")
	}

	var r0 *dto.AccountCreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*request_context.ReqCtx, *dto.AccountCreateRequest) (*dto.AccountCreateResponse, error)); ok {
		return rf(rCtx, reqBody)
	}
	if rf, ok := ret.Get(0).(func(*request_context.ReqCtx, *dto.AccountCreateRequest) *dto.AccountCreateResponse); ok {
		r0 = rf(rCtx, reqBody)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.AccountCreateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*request_context.ReqCtx, *dto.AccountCreateRequest) error); ok {
		r1 = rf(rCtx, reqBody)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccount provides a mock function with given fields: rCtx, accountId
func (_m *IAccountService) GetAccount(rCtx *request_context.ReqCtx, accountId int) (*dto.AccountGetResponse, error) {
	ret := _m.Called(rCtx, accountId)

	if len(ret) == 0 {
		panic("no return value specified for GetAccount")
	}

	var r0 *dto.AccountGetResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*request_context.ReqCtx, int) (*dto.AccountGetResponse, error)); ok {
		return rf(rCtx, accountId)
	}
	if rf, ok := ret.Get(0).(func(*request_context.ReqCtx, int) *dto.AccountGetResponse); ok {
		r0 = rf(rCtx, accountId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.AccountGetResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*request_context.ReqCtx, int) error); ok {
		r1 = rf(rCtx, accountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIAccountService creates a new instance of IAccountService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIAccountService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAccountService {
	mock := &IAccountService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
