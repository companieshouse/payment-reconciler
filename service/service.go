package service

import (
	"github.com/companieshouse/chs.go/log"
	"github.com/companieshouse/payment-reconciler/config"
	"github.com/companieshouse/payment-reconciler/dao"
	"github.com/companieshouse/payment-reconciler/models"
)

const transactionsFileNamePrefix = "CHS_PaymentTransactions_"
const productsFileNamePrefix = "CHS_PaymentProducts_"
const refundsFileNamePrefix = "CHS_Refunds_"
const csvFileSuffix = ".csv"

// Service provides functionality by which to fetch payment reconciliation CSV's
type Service interface {
	GetTransactionsCSV(reconciliationMetaData *models.ReconciliationMetaData) (models.CSV, error)
	GetProductsCSV(reconciliationMetaData *models.ReconciliationMetaData) (models.CSV, error)
	GetRefundsCSV(reconciliationMetaData *models.ReconciliationMetaData) (models.CSV, error)
}

// ServiceImpl provides a concrete implementation of the Service interface
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

	transactions, err := s.DAO.GetTransactionsData(reconciliationMetaData)
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

	products, err := s.DAO.GetProductsData(reconciliationMetaData)
	if err != nil {
		return csv, err
	}

	log.Info("Successfully retrieved products data.")
	log.Trace("Products data", log.Data{"products_data": products})

	csv = constructCSV(products, productsFileNamePrefix, reconciliationMetaData)

	return csv, nil
}

// GetRefundsCSV retrieves products data and constructs a CSV
func (s *ServiceImpl) GetRefundsCSV(reconciliationMetaData *models.ReconciliationMetaData) (models.CSV, error) {

	var csv models.CSV

	log.Info("Fetching refunds data.")

	refunds, err := s.DAO.GetRefundsData(reconciliationMetaData)
	if err != nil {
		return csv, err
	}

	log.Info("Successfully retrieved refunds data.")
	log.Trace("Refunds data", log.Data{"refunds_data": refunds})

	csv = constructCSV(refunds, refundsFileNamePrefix, reconciliationMetaData)

	return csv, nil
}

// constructCSV marshals CSVable data into a CSV, accompanied by a file name
func constructCSV(data models.CSVable, fileNamePrefix string, reconciliationMetaData *models.ReconciliationMetaData) models.CSV {

	return models.CSV{
		Data:     data,
		FileName: fileNamePrefix + reconciliationMetaData.ReconciliationDate + csvFileSuffix,
	}
}
