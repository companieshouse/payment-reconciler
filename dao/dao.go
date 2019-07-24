package dao

import (
	"github.com/companieshouse/payment-reconciler/models"
)

// DAO provides access to the database
type DAO interface {
	GetTransactionsData() (models.TransactionsList, error)
	GetProductsData() (models.ProductsList, error)
}
