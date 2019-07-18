package service

import (
	"github.com/companieshouse/chs.go/log"
	"github.com/companieshouse/payment-reconciler/app/config"
	"github.com/companieshouse/payment-reconciler/app/dao"
	"github.com/companieshouse/payment-reconciler/app/models"
)

const transactionsFileNamePrefix string = "CHS_PaymentTransactions_"
const productsFileNamePrefix string = "CHS_PaymentProducts_"
const csvFileSuffix = ".csv"

type Service interface {
	GetTransactionsCSV(reconciliationMetaData *models.ReconciliationMetaData) (models.CSV, error)
	GetProductsCSV(reconciliationMetaData *models.ReconciliationMetaData) (models.CSV, error)
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

// GetTransactionsCSV retrieves transactions data and constructs a CSV
func (s *ServiceImpl) GetTransactionsCSV(reconciliationMetaData *models.ReconciliationMetaData) (models.CSV, error) {

	var csv models.CSV

	log.Info("Fetching transactions data.")

	transactions, err := s.DAO.GetTransactionsData()
	if err != nil {
		return csv, err
	}

	log.Info("Successfully retrieved transactions data.")
	log.Trace("Transactions data", log.Data{"transactions_data": transactions})

	csv = constructCSV(transactions, transactionsFileNamePrefix, reconciliationMetaData)

	return csv, nil
}

// GetProductsCSV retrieves products data and constructs a CSV
func (s *ServiceImpl) GetProductsCSV(reconciliationMetaData *models.ReconciliationMetaData) (models.CSV, error) {

	var csv models.CSV

	log.Info("Fetching products data.")

	products, err := s.DAO.GetProductsData()
	if err != nil {
		return csv, err
	}

	log.Info("Successfully retrieved products data.")
	log.Trace("Products data", log.Data{"products_data": products})

	csv = constructCSV(products, productsFileNamePrefix, reconciliationMetaData)

	return csv, nil
}

// constructCSV marshalls CSVable data into a CSV, accompanied by a file name
func constructCSV(data models.CSVable, fileNamePrefix string, reconciliationMetaData *models.ReconciliationMetaData) models.CSV {

	return models.CSV{
		Data:     data,
		FileName: fileNamePrefix + reconciliationMetaData.Date + csvFileSuffix,
	}
}
