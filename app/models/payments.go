package models

import "reflect"

type PaymentProductsData struct {
	PaymentProducts []PaymentProduct
}

type PaymentProduct struct {
	PaymentRef    string `bson:"payment_ref"`
	ProductCode   string `bson:"product_code"`
	CompanyNumber string `bson:"company_number"`
	FilingDate    string `bson:"filing_date"`
	MadeUpDate    string `bson:"made_up_date"`
}

type PaymentTransactionsData struct {
	PaymentTransactions []PaymentTransaction
}

type PaymentTransaction struct {
	TransactionID     string `bson:"transaction_id"`
	TransactionDate   string `bson:"transaction_date"`
	Email             string `bson:"email"`
	PaymentMethod     string `bson:"payment_method"`
	Amount            string `bson:"amount"`
	CompanyNumber     string `bson:"company_number"`
	TransactionType   string `bson:"transaction_type"`
	OrderReference    string `bson:"order_reference"`
	Status            string `bson:"status"`
	UserID            string `bson:"user_id"`
	OriginalReference string `bson:"original_reference"`
	DisputeDetails    string `bson:"dispute_details"`
}

func (paymentTransactions PaymentTransactionsData) ToCSV() [][]string {

	csv := make([][]string, len(paymentTransactions.PaymentTransactions)+1)

	for i := 0; i < len(paymentTransactions.PaymentTransactions); i++ {
		if i == 0 {
			csv[i] = getHeaders(paymentTransactions.PaymentTransactions[i])
		}
		csv[i+1] = getSlice(paymentTransactions.PaymentTransactions[i])
	}

	return csv
}

func (paymentProducts PaymentProductsData) ToCSV() [][]string {

	csv := make([][]string, len(paymentProducts.PaymentProducts))

	for i := 0; i < len(paymentProducts.PaymentProducts); i++ {
		csv[i] = getSlice(paymentProducts.PaymentProducts[i])
	}

	return csv
}

func getSlice(resource interface{}) []string {

	val := reflect.ValueOf(resource)

	slice := make([]string, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		slice[i] = val.Field(i).String()
	}

	return slice
}

func getHeaders(resource interface{}) []string {

	val := reflect.ValueOf(resource)

	headers := make([]string, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		headers[i] = val.Type().Field(i).Name
	}

	return headers
}
