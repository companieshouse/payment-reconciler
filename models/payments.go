package models

import "reflect"

// ProductsData holds an array of payment products
type ProductsData struct {
	Products []Product
}

// Product contains data of a payment product
type Product struct {
	PaymentRef    string `bson:"payment_ref"`
	ProductCode   string `bson:"product_code"`
	CompanyNumber string `bson:"company_number"`
	FilingDate    string `bson:"filing_date"`
	MadeUpDate    string `bson:"made_up_date"`
}

// TransactionsData holds an array of payment transactions
type TransactionsData struct {
	Transactions []Transaction
}

// Transaction contains data of a payment transaction
type Transaction struct {
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

// ToCSV converts TransactionsData into CSV-writable data
func (transactions TransactionsData) ToCSV() [][]string {

	csv := make([][]string, len(transactions.Transactions)+1)

	for i := 0; i < len(transactions.Transactions); i++ {
		if i == 0 {
			csv[i] = getHeaders(transactions.Transactions[i])
		}
		csv[i+1] = getSlice(transactions.Transactions[i])
	}

	return csv
}

// ToCSV converts ProductsData into CSV-writable data
func (products ProductsData) ToCSV() [][]string {

	csv := make([][]string, len(products.Products))

	for i := 0; i < len(products.Products); i++ {
		csv[i] = getSlice(products.Products[i])
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
