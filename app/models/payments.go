package models

import "reflect"

type PaymentResourceData struct {
	Amount        string `bson:"amount"`
	TransactionId string `bson:"transaction_id"`
}

func (resource *PaymentResourceData) ToSlice() []string {

	val := reflect.ValueOf(resource).Elem()

	slice := make([]string, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		slice[i] = val.Field(i).String()
	}

	return slice
}

func (resource *PaymentResourceData) GetReconciliationDataHeaders() []string {

	val := reflect.ValueOf(resource).Elem()

	headers := make([]string, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		headers[i] = val.Type().Field(i).Name
	}

	return headers
}
