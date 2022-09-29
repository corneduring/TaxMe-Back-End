package datastruct

import "time"

const CalculationsTableName = "calculations"

type calculation struct {
	CalculationID    int
	UserID           int
	PaymentFrequency string `json:"Frequency"`
	Salary           string
	MonthlyTax       string
	YearlyTax        string
	Timestamp        time.Time
}
