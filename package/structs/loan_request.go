package structs

import (
	"fmt"
	"time"
)

type LoanRequest struct{
	BorrowerID      string  `json:"borrower_id"`
	PrincipalAmount int64    `json:"principal_amount"`
	Rate            float64 `json:"rate"`
	Roi             float64 `json:"roi"`
}

func GenerateCodeLoan(borrowId string) string {
	date := time.Now().Day()
	return  fmt.Sprintf("AM%s%d",borrowId, date)
}