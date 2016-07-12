// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/citwild/wfe/stores (interfaces: Accounts)

package mock_stores

import (
	api "github.com/citwild/wfe/api"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
)

// Mock of Accounts interface
type MockAccounts struct {
	ctrl     *gomock.Controller
	recorder *_MockAccountsRecorder
}

// Recorder for MockAccounts (not exported)
type _MockAccountsRecorder struct {
	mock *MockAccounts
}

func NewMockAccounts(ctrl *gomock.Controller) *MockAccounts {
	mock := &MockAccounts{ctrl: ctrl}
	mock.recorder = &_MockAccountsRecorder{mock}
	return mock
}

func (_m *MockAccounts) EXPECT() *_MockAccountsRecorder {
	return _m.recorder
}

func (_m *MockAccounts) Create(_param0 context.Context, _param1 *api.User, _param2 *api.EmailAddr) (*api.User, error) {
	ret := _m.ctrl.Call(_m, "Create", _param0, _param1, _param2)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccountsRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0, arg1, arg2)
}