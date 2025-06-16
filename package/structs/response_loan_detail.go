package structs

type ResponseLoanDetail struct {
	LoanID          string  `json:"loan_id"`
	BorrowerID      string  `json:"borrower_id"`
	PrincipalAmount int     `json:"principal_amount"`
	Rate            float64 `json:"rate"`
	Roi             float64 `json:"roi"`
	AgreementURL    string  `json:"agreement_url"`
	State           string  `json:"state"`
}