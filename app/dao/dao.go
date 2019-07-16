package dao

import (
	"github.com/companieshouse/payment-reconciler/app/models"
)

type DAO interface {
	GetReconciliationData() (*models.ReconciliationData, error)
}
