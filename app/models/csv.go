package models

type CSV struct {
	Data     CSVable
	FileName string
}

type CSVable interface {
	ToCSV() [][]string
}
