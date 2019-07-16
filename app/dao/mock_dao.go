package dao

import (
	reflect "reflect"

	models "github.com/companieshouse/payment-reconciler/app/models"
	gomock "github.com/golang/mock/gomock"
)

// MockDAO is a mock of DAO interface
type MockDAO struct {
	ctrl     *gomock.Controller
	recorder *MockDAOMockRecorder
}

// MockDAOMockRecorder is the mock recorder for MockDAO
type MockDAOMockRecorder struct {
	mock *MockDAO
}

// NewMockDAO creates a new mock instance
func NewMockDAO(ctrl *gomock.Controller) *MockDAO {
	mock := &MockDAO{ctrl: ctrl}
	mock.recorder = &MockDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDAO) EXPECT() *MockDAOMockRecorder {
	return m.recorder
}

// GetReconciliationData mocks base method
func (m *MockDAO) GetReconciliationData() (*models.ReconciliationData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReconciliationData")
	ret0, _ := ret[0].(*models.ReconciliationData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReconciliationData indicates an expected call of GetReconciliationData
func (mr *MockDAOMockRecorder) GetReconciliationData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReconciliationData", reflect.TypeOf((*MockDAO)(nil).GetReconciliationData))
}
