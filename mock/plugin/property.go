// Code generated by MockGen. DO NOT EDIT.
// Source: ./plugin/property.go

// Package plugin is a generated GoMock package.
package plugin

import (
	models "github.com/baetyl/baetyl-cloud/v2/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProperty is a mock of Property interface.
type MockProperty struct {
	ctrl     *gomock.Controller
	recorder *MockPropertyMockRecorder
}

// MockPropertyMockRecorder is the mock recorder for MockProperty.
type MockPropertyMockRecorder struct {
	mock *MockProperty
}

// NewMockProperty creates a new mock instance.
func NewMockProperty(ctrl *gomock.Controller) *MockProperty {
	mock := &MockProperty{ctrl: ctrl}
	mock.recorder = &MockPropertyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProperty) EXPECT() *MockPropertyMockRecorder {
	return m.recorder
}

// CreateProperty mocks base method.
func (m *MockProperty) CreateProperty(property *models.Property) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProperty", property)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProperty indicates an expected call of CreateProperty.
func (mr *MockPropertyMockRecorder) CreateProperty(property interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProperty", reflect.TypeOf((*MockProperty)(nil).CreateProperty), property)
}

// DeleteProperty mocks base method.
func (m *MockProperty) DeleteProperty(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProperty", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProperty indicates an expected call of DeleteProperty.
func (mr *MockPropertyMockRecorder) DeleteProperty(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProperty", reflect.TypeOf((*MockProperty)(nil).DeleteProperty), key)
}

// ListProperty mocks base method.
func (m *MockProperty) ListProperty(page *models.Filter) ([]models.Property, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProperty", page)
	ret0, _ := ret[0].([]models.Property)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProperty indicates an expected call of ListProperty.
func (mr *MockPropertyMockRecorder) ListProperty(page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProperty", reflect.TypeOf((*MockProperty)(nil).ListProperty), page)
}

// CountProperty mocks base method.
func (m *MockProperty) CountProperty(key string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountProperty", key)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountProperty indicates an expected call of CountProperty.
func (mr *MockPropertyMockRecorder) CountProperty(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountProperty", reflect.TypeOf((*MockProperty)(nil).CountProperty), key)
}

// UpdateProperty mocks base method.
func (m *MockProperty) UpdateProperty(property *models.Property) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProperty", property)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProperty indicates an expected call of UpdateProperty.
func (mr *MockPropertyMockRecorder) UpdateProperty(property interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProperty", reflect.TypeOf((*MockProperty)(nil).UpdateProperty), property)
}

// GetPropertyValue mocks base method.
func (m *MockProperty) GetPropertyValue(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPropertyValue", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPropertyValue indicates an expected call of GetPropertyValue.
func (mr *MockPropertyMockRecorder) GetPropertyValue(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPropertyValue", reflect.TypeOf((*MockProperty)(nil).GetPropertyValue), key)
}

// Close mocks base method.
func (m *MockProperty) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockPropertyMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockProperty)(nil).Close))
}