package structs

import "time"

type Loans struct {
    Id 			    string `json:"id"`
    BorrowerId      string `json:"borrower_id"`
    PrincipalAmount int64 `json:"principal_amount"`
    Rate            float64 `json:"rate"`
    Roi             float64 `json:"roi"`
    State           int64 `json:"state"`
    AgreementUrl    string `json:"agreement_url"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}

func (p LoanRequest) NewLoan() Loans { 
	return Loans{
		Id: GenerateCodeLoan(p.BorrowerID),
		BorrowerId: p.BorrowerID,
		PrincipalAmount: p.PrincipalAmount,
		Rate: p.Rate,
		Roi: p.Roi,
		State: 1, // 1 Pengajuan
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}