package structs

type RequestInvestment struct{
	InvestorID string `json:"investor_id"`
	LoanID     string `json:"loan_id"`
	Amount     int64  `json:"amount"`
}