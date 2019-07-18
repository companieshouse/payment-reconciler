package models

// CSV holds CSV-able data, accompanied by a file name
type CSV struct {
	Data     CSVable
	FileName string
}

// CSVable provides functionality to convert a struct type to CSV-writable data
type CSVable interface {
	ToCSV() [][]string
}
