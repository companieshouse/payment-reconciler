package filetransfer

type FileTransfer interface {
	UploadCSV(csvData [][]string) error
}
