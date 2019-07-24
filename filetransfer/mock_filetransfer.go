// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/companieshouse/payment-reconciler/app/filetransfer (interfaces: FileTransfer)
package filetransfer

import (
	"reflect"

	"github.com/companieshouse/payment-reconciler/models"
	"github.com/golang/mock/gomock"
)

// MockFileTransfer is a mock of FileTransfer interface
type MockFileTransfer struct {
	ctrl     *gomock.Controller
	recorder *MockFileTransferMockRecorder
}

// MockFileTransferMockRecorder is the mock recorder for MockFileTransfer
type MockFileTransferMockRecorder struct {
	mock *MockFileTransfer
}

// NewMockFileTransfer creates a new mock instance
func NewMockFileTransfer(ctrl *gomock.Controller) *MockFileTransfer {
	mock := &MockFileTransfer{ctrl: ctrl}
	mock.recorder = &MockFileTransferMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileTransfer) EXPECT() *MockFileTransferMockRecorder {
	return m.recorder
}

// UploadCSVFiles mocks base method
func (m *MockFileTransfer) UploadCSVFiles(arg0 []models.CSV) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadCSVFiles", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadCSVFiles indicates an expected call of UploadCSVFiles
func (mr *MockFileTransferMockRecorder) UploadCSVFiles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadCSVFiles", reflect.TypeOf((*MockFileTransfer)(nil).UploadCSVFiles), arg0)
}
