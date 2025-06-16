package repository

import "context"


func (r InvestRepository) GetByLoanID(ctx context.Context, investorID string, loanId string, dest interface{}) error{
	return r.db.Where("investor_id = ? AND loan_id = ?", investorID, loanId).Find(dest).Error
}