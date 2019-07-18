package filetransfer

import "github.com/companieshouse/payment-reconciler/app/models"

// FileTransfer provides an interface by which to upload CSV files
type FileTransfer interface {
	UploadCSVFiles(csvs []models.CSV) error
}
