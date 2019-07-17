package dao

import (
	"github.com/companieshouse/payment-reconciler/app/models"
)

type DAO interface {
	GetPaymentTransactionsData() (models.PaymentTransactionsData, error)
	GetPaymentProductsData() (models.PaymentProductsData, error)
}
