package models

import "time"

// ReconciliationMetaData holds meta data regarding the date for which to reconcile payments
type ReconciliationMetaData struct {
	ReconciliationDate string `json:"reconciliation_date"`
	StartTime          time.Time
	EndTime            time.Time
}
