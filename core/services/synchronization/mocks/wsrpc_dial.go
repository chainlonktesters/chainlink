// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	wsrpc "github.com/smartcontractkit/wsrpc"
)

// WsrpcDial is an autogenerated mock type for the WsrpcDial type
type WsrpcDial struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *WsrpcDial) Execute(_a0 string, _a1 ...wsrpc.DialOption) (*wsrpc.ClientConn, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *wsrpc.ClientConn
	if rf, ok := ret.Get(0).(func(string, ...wsrpc.DialOption) *wsrpc.ClientConn); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wsrpc.ClientConn)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...wsrpc.DialOption) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
