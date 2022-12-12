package dao

import (
	"github.com/companieshouse/payment-reconciler/models"
)

// DAO provides access to the database
type DAO interface {
	GetTransactionsData(reconciliationMetaData *models.ReconciliationMetaData) (models.TransactionsList, error)
	GetProductsData(reconciliationMetaData *models.ReconciliationMetaData) (models.ProductsList, error)
	GetRefundsData(reconciliationMetaData *models.ReconciliationMetaData) (models.RefundsList, error)
	GetAutoRefundsData(reconciliationMetaData *models.ReconciliationMetaData) (models.RefundsList, error)
}
