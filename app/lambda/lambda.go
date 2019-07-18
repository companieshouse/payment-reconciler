package lambda

import (
	"github.com/companieshouse/chs.go/log"
	"github.com/companieshouse/payment-reconciler/app/config"
	"github.com/companieshouse/payment-reconciler/app/filetransfer"
	"github.com/companieshouse/payment-reconciler/app/models"
	"github.com/companieshouse/payment-reconciler/app/service"
)

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

	log.Info("Payment reconciliation lambda executing. Creating transactions CSV.")

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

	log.Info("Products CSV constructed. Preparing to upload CSV's.")
	log.Trace("Products CSV", log.Data{"products_csv": productsCSV})

	err = lambda.FileTransfer.UploadCSVFiles([]models.CSV{transactionsCSV, productsCSV})
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("CSV's successfully uploaded. Lambda execution finished.")

	return nil
}
