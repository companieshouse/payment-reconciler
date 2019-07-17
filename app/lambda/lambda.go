package lambda

import (
	"github.com/companieshouse/chs.go/log"
	"github.com/companieshouse/payment-reconciler/app/config"
	"github.com/companieshouse/payment-reconciler/app/filetransfer"
	"github.com/companieshouse/payment-reconciler/app/models"
	"github.com/companieshouse/payment-reconciler/app/service"
)

type Lambda struct {
	Config       *config.Config
	Service      service.Service
	FileTransfer filetransfer.FileTransfer
}

func New(cfg *config.Config) *Lambda {

	return &Lambda{
		Config:       cfg,
		Service:      service.New(cfg),
		FileTransfer: filetransfer.New(cfg),
	}
}

// Execute handles lambda execution
func (lambda *Lambda) Execute(reconciliationMetaData *models.ReconciliationMetaData) error {

	log.Info("Payment reconciliation lambda executing. Creating payment transactions CSV.")

	paymentTransactionsCSV, err := lambda.Service.GetPaymentTransactionsCSV(reconciliationMetaData)
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("Payment transactions CSV constructed. Creating payment products CSV.")
	log.Trace("Payment transactions CSV", log.Data{"payment_transactions_csv": paymentTransactionsCSV})

	paymentProductsCSV, err := lambda.Service.GetPaymentProductsCSV(reconciliationMetaData)
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("Payment products CSV constructed. Preparing to upload CSV's.")
	log.Trace("Payment products CSV", log.Data{"payment_products_csv": paymentProductsCSV})

	err = lambda.FileTransfer.UploadCSVFiles([]models.CSV{paymentTransactionsCSV, paymentProductsCSV})
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("CSV's successfully uploaded. Lambda execution finished.")

	return nil
}
