// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/citwild/wfe/store (interfaces: PasswordStore)

package mockstore

import (
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
)

// Mock of PasswordStore interface
type MockPasswordStore struct {
	ctrl     *gomock.Controller
	recorder *_MockPasswordStoreRecorder
}

// Recorder for MockPasswordStore (not exported)
type _MockPasswordStoreRecorder struct {
	mock *MockPasswordStore
}

func NewMockPasswordStore(ctrl *gomock.Controller) *MockPasswordStore {
	mock := &MockPasswordStore{ctrl: ctrl}
	mock.recorder = &_MockPasswordStoreRecorder{mock}
	return mock
}

func (_m *MockPasswordStore) EXPECT() *_MockPasswordStoreRecorder {
	return _m.recorder
}

func (_m *MockPasswordStore) SetPassword(_param0 context.Context, _param1 int32, _param2 string) error {
	ret := _m.ctrl.Call(_m, "SetPassword", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockPasswordStoreRecorder) SetPassword(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetPassword", arg0, arg1, arg2)
}