package lambda

import (
	"time"

	"github.com/companieshouse/chs.go/log"
	"github.com/companieshouse/payment-reconciler/config"
	"github.com/companieshouse/payment-reconciler/filetransfer"
	"github.com/companieshouse/payment-reconciler/models"
	"github.com/companieshouse/payment-reconciler/service"
)

const dateFormat = "2006-01-02"

// Lambda provides AWS lambda execution functionality
type Lambda struct {
	Config       *config.Config
	Service      service.Service
	FileTransfer filetransfer.FileTransfer
}

// New returns a new Lambda using the provided configs
func New(cfg *config.Config) *Lambda {

	return &Lambda{
		Config:       cfg,
		Service:      service.New(cfg),
		FileTransfer: filetransfer.New(cfg),
	}
}

// Execute handles lambda execution
func (lambda *Lambda) Execute(reconciliationMetaData *models.ReconciliationMetaData) error {
	if reconciliationMetaData.ReconciliationDate == "" {

		// If no date/time supplied, set start and end to beginning and end of previous day
		reconciliationDateTime := time.Now().Add(time.Duration(-24) * time.Hour)
		reconciliationMetaData.ReconciliationDate = reconciliationDateTime.Format(dateFormat)

		startTime := reconciliationDateTime.Truncate(24 * time.Hour)
		reconciliationMetaData.StartTime = startTime
		reconciliationMetaData.EndTime = startTime.Add(24 * time.Hour)
	} else {

		startTime, err := time.Parse(dateFormat, reconciliationMetaData.ReconciliationDate)
		if err != nil {
			log.Error(err)
			return err
		}

		reconciliationMetaData.StartTime = startTime
		reconciliationMetaData.EndTime = startTime.Add(24 * time.Hour)
	}

	log.Info("Payment reconciliation lambda executing. Reconciling payments for date: " + reconciliationMetaData.ReconciliationDate + ". Creating transactions CSV.")

	transactionsCSV, err := lambda.Service.GetTransactionsCSV(reconciliationMetaData)
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("Transactions CSV constructed. Creating products CSV.")
	log.Trace("Transactions CSV", log.Data{"transactions_csv": transactionsCSV})

	productsCSV, err := lambda.Service.GetProductsCSV(reconciliationMetaData)
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("Products CSV constructed. Creating refunds CSV.")
	log.Trace("Products CSV", log.Data{"products_csv": productsCSV})

	refundsCSV, err := lambda.Service.GetRefundsCSV(reconciliationMetaData)
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("Refunds CSV constructed. Preparing to upload CSV's.")
	log.Trace("Refunds CSV", log.Data{"refunds_csv": refundsCSV})

	err = lambda.FileTransfer.UploadCSVFiles([]models.CSV{transactionsCSV, productsCSV, refundsCSV})
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("CSV's successfully uploaded. Lambda execution finished.")

	return nil
}
