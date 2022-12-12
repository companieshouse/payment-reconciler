package models

import (
	"time"
)

// RefundsList holds an array of refund products
type RefundsList struct {
	Refunds []Refund
}

// AutoRefundsList holds an array of refund products
type AutoRefundsList struct {
	Refunds []AutoRefund
}

// Refund contains data of a refund
type Refund struct {
	TransactionID     string    `bson:"transaction_id"`
	TransactionDate   time.Time `bson:"transaction_date"`
	Email             string    `bson:"email"`
	PaymentMethod     string    `bson:"payment_method"`
	Amount            string    `bson:"amount"`
	CompanyNumber     string    `bson:"company_number"`
	TransactionType   string    `bson:"transaction_type"`
	OrderReference    string    `bson:"order_reference"`
	Status            string    `bson:"status"`
	UserID            string    `bson:"user_id"`
	OriginalReference string    `bson:"original_reference"`
	DisputeDetails    string    `bson:"dispute_details"`
	ProductCode       int       `bson:"product_code"`
}

type AutoRefund struct {
	TransactionID   string    `bson:"transaction_id"`
	TransactionDate time.Time `bson:"transaction_date"`
	TransactionType string    `bson:"transaction_type"`
	PaymentMethod   string    `bson:"payment_method"`
	Amount          string    `bson:"amount"`
	Status          string    `bson:"status"`
	CompanyNumber   string    `bson:"company_number"`
	ProductCode     int       `bson:"product_code"`
}

// ToCSV converts RefundsList into CSV-writable data
func (refunds RefundsList) ToCSV() [][]string {

	csv := make([][]string, len(refunds.Refunds)+1)

	for i := 0; i < len(refunds.Refunds); i++ {
		if i == 0 {
			csv[i] = getHeaders(refunds.Refunds[i])
		}
		csv[i+1] = getSlice(refunds.Refunds[i])
	}

	return csv
}
