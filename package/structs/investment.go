package structs

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// LoanInvestments
type LoanInvestments struct {
    Id 			string `json:"id"`
    LoanId 		string `json:"loan_id"`
    InvestorId 	string `json:"investor_id"`
    Amount 		int64  `json:"amount"`
    CreatedAt 	time.Time `json:"created_at"`
}

func (p RequestInvestment) NewInvestment() LoanInvestments {
	return LoanInvestments{
		Id: uuid.New().String(),
		LoanId: p.LoanID,
		InvestorId: fmt.Sprintf("inv-%d", GenerateInvestrorID()),
		Amount: int64(p.Amount),
		CreatedAt: time.Now(),
	}
}

func GenerateInvestrorID() int64  {
	return time.Now().UnixMilli()
}