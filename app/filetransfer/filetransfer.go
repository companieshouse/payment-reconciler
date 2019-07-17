package filetransfer

import "github.com/companieshouse/payment-reconciler/app/models"

type FileTransfer interface {
	UploadCSVFiles(csvs []models.CSV) error
}
