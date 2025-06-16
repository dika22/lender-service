package structs

type ResponseLoanInvestment struct {
	LoanID                string `json:"loan_id"`
	CurrentInvestedAmount int64   `json:"current_invested_amount"`
	State                 string `json:"state"`
}