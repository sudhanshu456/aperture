// Code generated by MockGen. DO NOT EDIT.
// Source: metrics-registry.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	prometheus "github.com/prometheus/client_golang/prometheus"
)

// MockResponseMetricsAPI is a mock of ResponseMetricsAPI interface.
type MockResponseMetricsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockResponseMetricsAPIMockRecorder
}

// MockResponseMetricsAPIMockRecorder is the mock recorder for MockResponseMetricsAPI.
type MockResponseMetricsAPIMockRecorder struct {
	mock *MockResponseMetricsAPI
}

// NewMockResponseMetricsAPI creates a new mock instance.
func NewMockResponseMetricsAPI(ctrl *gomock.Controller) *MockResponseMetricsAPI {
	mock := &MockResponseMetricsAPI{ctrl: ctrl}
	mock.recorder = &MockResponseMetricsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResponseMetricsAPI) EXPECT() *MockResponseMetricsAPIMockRecorder {
	return m.recorder
}

// DeleteTokenLatencyHistogram mocks base method.
func (m *MockResponseMetricsAPI) DeleteTokenLatencyHistogram(labels map[string]string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTokenLatencyHistogram", labels)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DeleteTokenLatencyHistogram indicates an expected call of DeleteTokenLatencyHistogram.
func (mr *MockResponseMetricsAPIMockRecorder) DeleteTokenLatencyHistogram(labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTokenLatencyHistogram", reflect.TypeOf((*MockResponseMetricsAPI)(nil).DeleteTokenLatencyHistogram), labels)
}

// GetTokenLatencyHistogram mocks base method.
func (m *MockResponseMetricsAPI) GetTokenLatencyHistogram(labels map[string]string) (prometheus.Observer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenLatencyHistogram", labels)
	ret0, _ := ret[0].(prometheus.Observer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTokenLatencyHistogram indicates an expected call of GetTokenLatencyHistogram.
func (mr *MockResponseMetricsAPIMockRecorder) GetTokenLatencyHistogram(labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenLatencyHistogram", reflect.TypeOf((*MockResponseMetricsAPI)(nil).GetTokenLatencyHistogram), labels)
}