package repository

import (
	"context"
	"lender-service/package/structs"
)

func (r LoanRepository) LoanApproval(ctx context.Context, payload structs.LoanHistory) error{
	return r.db.Table("loan_history").Create(&payload).Error
}