// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	context "context"

	client "sigs.k8s.io/controller-runtime/pkg/client"

	exposer "github.com/vmware-tanzu/velero/pkg/exposer"

	mock "github.com/stretchr/testify/mock"

	time "time"

	v1 "k8s.io/api/core/v1"
)

// GenericRestoreExposer is an autogenerated mock type for the GenericRestoreExposer type
type GenericRestoreExposer struct {
	mock.Mock
}

// CleanUp provides a mock function with given fields: _a0, _a1
func (_m *GenericRestoreExposer) CleanUp(_a0 context.Context, _a1 v1.ObjectReference) {
	_m.Called(_a0, _a1)
}

// DiagnoseExpose provides a mock function with given fields: _a0, _a1
func (_m *GenericRestoreExposer) DiagnoseExpose(_a0 context.Context, _a1 v1.ObjectReference) string {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DiagnoseExpose")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, v1.ObjectReference) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Expose provides a mock function with given fields: _a0, _a1, _a2
func (_m *GenericRestoreExposer) Expose(_a0 context.Context, _a1 v1.ObjectReference, _a2 exposer.GenericRestoreExposeParam) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for Expose")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ObjectReference, exposer.GenericRestoreExposeParam) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetExposed provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *GenericRestoreExposer) GetExposed(_a0 context.Context, _a1 v1.ObjectReference, _a2 client.Client, _a3 string, _a4 time.Duration) (*exposer.ExposeResult, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	if len(ret) == 0 {
		panic("no return value specified for GetExposed")
	}

	var r0 *exposer.ExposeResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ObjectReference, client.Client, string, time.Duration) (*exposer.ExposeResult, error)); ok {
		return rf(_a0, _a1, _a2, _a3, _a4)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1.ObjectReference, client.Client, string, time.Duration) *exposer.ExposeResult); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*exposer.ExposeResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1.ObjectReference, client.Client, string, time.Duration) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PeekExposed provides a mock function with given fields: _a0, _a1
func (_m *GenericRestoreExposer) PeekExposed(_a0 context.Context, _a1 v1.ObjectReference) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for PeekExposed")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ObjectReference) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RebindVolume provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *GenericRestoreExposer) RebindVolume(_a0 context.Context, _a1 v1.ObjectReference, _a2 string, _a3 string, _a4 time.Duration) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	if len(ret) == 0 {
		panic("no return value specified for RebindVolume")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ObjectReference, string, string, time.Duration) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewGenericRestoreExposer creates a new instance of GenericRestoreExposer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGenericRestoreExposer(t interface {
	mock.TestingT
	Cleanup(func())
}) *GenericRestoreExposer {
	mock := &GenericRestoreExposer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
