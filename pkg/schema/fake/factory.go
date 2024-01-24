// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rancher/steve/pkg/schema (interfaces: Factory)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "gitlab.devops.telekom.de/caas/rancher/apiserver/pkg/types"
	schema "github.com/rancher/steve/pkg/schema"
	schema0 "k8s.io/apimachinery/pkg/runtime/schema"
	user "k8s.io/apiserver/pkg/authentication/user"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// AddTemplate mocks base method.
func (m *MockFactory) AddTemplate(arg0 ...schema.Template) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddTemplate", varargs...)
}

// AddTemplate indicates an expected call of AddTemplate.
func (mr *MockFactoryMockRecorder) AddTemplate(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTemplate", reflect.TypeOf((*MockFactory)(nil).AddTemplate), arg0...)
}

// ByGVK mocks base method.
func (m *MockFactory) ByGVK(arg0 schema0.GroupVersionKind) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByGVK", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// ByGVK indicates an expected call of ByGVK.
func (mr *MockFactoryMockRecorder) ByGVK(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByGVK", reflect.TypeOf((*MockFactory)(nil).ByGVK), arg0)
}

// ByGVR mocks base method.
func (m *MockFactory) ByGVR(arg0 schema0.GroupVersionResource) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByGVR", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// ByGVR indicates an expected call of ByGVR.
func (mr *MockFactoryMockRecorder) ByGVR(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByGVR", reflect.TypeOf((*MockFactory)(nil).ByGVR), arg0)
}

// OnChange mocks base method.
func (m *MockFactory) OnChange(arg0 context.Context, arg1 func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnChange", arg0, arg1)
}

// OnChange indicates an expected call of OnChange.
func (mr *MockFactoryMockRecorder) OnChange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnChange", reflect.TypeOf((*MockFactory)(nil).OnChange), arg0, arg1)
}

// Schemas mocks base method.
func (m *MockFactory) Schemas(arg0 user.Info) (*types.APISchemas, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Schemas", arg0)
	ret0, _ := ret[0].(*types.APISchemas)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Schemas indicates an expected call of Schemas.
func (mr *MockFactoryMockRecorder) Schemas(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Schemas", reflect.TypeOf((*MockFactory)(nil).Schemas), arg0)
}
