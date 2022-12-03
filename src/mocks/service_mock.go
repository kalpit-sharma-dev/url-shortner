// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/kalpit-sharma-dev/url-shortner/src/model"
	reflect "reflect"
)

// MockUrlService is a mock of UrlService interface
type MockUrlService struct {
	ctrl     *gomock.Controller
	recorder *MockUrlServiceMockRecorder
}

// MockUrlServiceMockRecorder is the mock recorder for MockUrlService
type MockUrlServiceMockRecorder struct {
	mock *MockUrlService
}

// NewMockUrlService creates a new mock instance
func NewMockUrlService(ctrl *gomock.Controller) *MockUrlService {
	mock := &MockUrlService{ctrl: ctrl}
	mock.recorder = &MockUrlServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUrlService) EXPECT() *MockUrlServiceMockRecorder {
	return m.recorder
}

// CreateUrl mocks base method
func (m *MockUrlService) CreateUrl(ctx context.Context, urlReq model.Request) (model.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUrl", ctx, urlReq)
	ret0, _ := ret[0].(model.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUrl indicates an expected call of CreateUrl
func (mr *MockUrlServiceMockRecorder) CreateUrl(ctx, urlReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUrl", reflect.TypeOf((*MockUrlService)(nil).CreateUrl), ctx, urlReq)
}

// GetUrl mocks base method
func (m *MockUrlService) GetUrl(ctx context.Context, url string) (model.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUrl", ctx, url)
	ret0, _ := ret[0].(model.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUrl indicates an expected call of GetUrl
func (mr *MockUrlServiceMockRecorder) GetUrl(ctx, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUrl", reflect.TypeOf((*MockUrlService)(nil).GetUrl), ctx, url)
}
