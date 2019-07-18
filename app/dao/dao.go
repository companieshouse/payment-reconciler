package dao

import (
	"github.com/companieshouse/payment-reconciler/app/models"
)

type DAO interface {
	GetTransactionsData() (models.TransactionsData, error)
	GetProductsData() (models.ProductsData, error)
}
