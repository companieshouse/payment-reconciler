// Code generated by MockGen. DO NOT EDIT.
// Source: dao/dao.go

// Package dao is a generated GoMock package.
package dao

import (
	models "github.com/companieshouse/payment-reconciler/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
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

// GetTransactionsData mocks base method
func (m *MockDAO) GetTransactionsData(reconciliationMetaData *models.ReconciliationMetaData) (models.TransactionsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionsData", reconciliationMetaData)
	ret0, _ := ret[0].(models.TransactionsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionsData indicates an expected call of GetTransactionsData
func (mr *MockDAOMockRecorder) GetTransactionsData(reconciliationMetaData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionsData", reflect.TypeOf((*MockDAO)(nil).GetTransactionsData), reconciliationMetaData)
}

// GetProductsData mocks base method
func (m *MockDAO) GetProductsData(reconciliationMetaData *models.ReconciliationMetaData) (models.ProductsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsData", reconciliationMetaData)
	ret0, _ := ret[0].(models.ProductsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsData indicates an expected call of GetProductsData
func (mr *MockDAOMockRecorder) GetProductsData(reconciliationMetaData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsData", reflect.TypeOf((*MockDAO)(nil).GetProductsData), reconciliationMetaData)
}

// GetRefundsData mocks base method
func (m *MockDAO) GetRefundsData(reconciliationMetaData *models.ReconciliationMetaData) (models.RefundsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRefundsData", reconciliationMetaData)
	ret0, _ := ret[0].(models.RefundsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRefundsData indicates an expected call of GetRefundsData
func (mr *MockDAOMockRecorder) GetRefundsData(reconciliationMetaData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRefundsData", reflect.TypeOf((*MockDAO)(nil).GetRefundsData), reconciliationMetaData)
}

// GetAutoRefundsData mocks base method
func (m *MockDAO) GetAutoRefundsData(reconciliationMetaData *models.ReconciliationMetaData) (models.RefundsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAutoRefundsData", reconciliationMetaData)
	ret0, _ := ret[0].(models.RefundsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAutoRefundsData indicates an expected call of GetRefundsData
func (mr *MockDAOMockRecorder) GetAutoRefundsData(reconciliationMetaData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutoRefundsData", reflect.TypeOf((*MockDAO)(nil).GetAutoRefundsData), reconciliationMetaData)
}
