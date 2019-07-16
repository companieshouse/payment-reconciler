package filetransfer

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// UploadCSV mocks base method
func (m *MockFileTransfer) UploadCSV(arg0 [][]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadCSV", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadCSV indicates an expected call of UploadCSV
func (mr *MockFileTransferMockRecorder) UploadCSV(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadCSV", reflect.TypeOf((*MockFileTransfer)(nil).UploadCSV), arg0)
}
