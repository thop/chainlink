// Code generated by mockery v2.1.0. DO NOT EDIT.

package mocks

import (
	models "github.com/smartcontractkit/chainlink/core/store/models"
	mock "github.com/stretchr/testify/mock"
)

// RunQueue is an autogenerated mock type for the RunQueue type
type RunQueue struct {
	mock.Mock
}

// Run provides a mock function with given fields: _a0
func (_m *RunQueue) Run(_a0 *models.JobRun) {
	_m.Called(_a0)
}

// Start provides a mock function with given fields:
func (_m *RunQueue) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *RunQueue) Stop() {
	_m.Called()
}

// WorkerCount provides a mock function with given fields:
func (_m *RunQueue) WorkerCount() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
