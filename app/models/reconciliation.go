package models

type Reconciliation struct {
	Date string `json:"date"`
}

type ReconciliationData struct {
	Payments []PaymentResourceData
}

func (reconciliationData *ReconciliationData) ToCSV() [][]string {

	payments := reconciliationData.Payments

	csvData := make([][]string, len(payments)+1)

	for i := 0; i < len(payments); i++ {
		if i == 0 {
			csvData[i] = payments[i].GetReconciliationDataHeaders()
		}
		csvData[i+1] = payments[i].ToSlice()
	}

	return csvData
}
