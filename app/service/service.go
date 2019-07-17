package service

import (
	"github.com/companieshouse/chs.go/log"
	"github.com/companieshouse/payment-reconciler/app/config"
	"github.com/companieshouse/payment-reconciler/app/dao"
	"github.com/companieshouse/payment-reconciler/app/models"
)

const paymentTransactionsFileNamePrefix string = "CHS_PaymentTransactions_"
const paymentProductsFileNamePrefix string = "CHS_PaymentProducts_"
const csvFileSuffix = ".csv"

type Service interface {
	GetPaymentTransactionsCSV(reconciliationMetaData *models.ReconciliationMetaData) (models.CSV, error)
	GetPaymentProductsCSV(reconciliationMetaData *models.ReconciliationMetaData) (models.CSV, error)
}

type ServiceImpl struct {
	Config *config.Config
	DAO    dao.DAO
}

// New returns a new, concrete implementation of the Service interface
func New(cfg *config.Config) *ServiceImpl {

	return &ServiceImpl{
		Config: cfg,
		DAO:    dao.New(cfg),
	}
}

// GetPaymentTransactionsCSV retrieves payment transactions data and constructs a CSV
func (s *ServiceImpl) GetPaymentTransactionsCSV(reconciliationMetaData *models.ReconciliationMetaData) (models.CSV, error) {

	var csv models.CSV

	log.Info("Fetching payment transactions data.")

	paymentTransactions, err := s.DAO.GetPaymentTransactionsData()
	if err != nil {
		return csv, err
	}

	log.Info("Successfully retrieved payment transactions data.")
	log.Trace("Payment transactions data", log.Data{"payment_transactions_data": paymentTransactions})

	csv = constructCSV(paymentTransactions, paymentTransactionsFileNamePrefix, reconciliationMetaData)

	return csv, nil
}

// GetPaymentProductsCSV retrieves payment products data and constructs a CSV
func (s *ServiceImpl) GetPaymentProductsCSV(reconciliationMetaData *models.ReconciliationMetaData) (models.CSV, error) {

	var csv models.CSV

	log.Info("Fetching payment products data.")

	paymentProducts, err := s.DAO.GetPaymentProductsData()
	if err != nil {
		return csv, err
	}

	log.Info("Successfully retrieved payment products data.")
	log.Trace("Payment products data", log.Data{"payment_products_data": paymentProducts})

	csv = constructCSV(paymentProducts, paymentProductsFileNamePrefix, reconciliationMetaData)

	return csv, nil
}

// constructCSV marshalls CSVable data into a CSV, accompanied by a file name
func constructCSV(data models.CSVable, fileNamePrefix string, reconciliationMetaData *models.ReconciliationMetaData) models.CSV {

	return models.CSV{
		Data:     data,
		FileName: fileNamePrefix + reconciliationMetaData.Date + csvFileSuffix,
	}
}
