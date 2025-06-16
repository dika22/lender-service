package repository

import (
	"context"
	"lender-service/package/structs"
)

func (r LoanRepository) Count(ctx context.Context, loan_id string,  count *int64) error {
	return r.db.Model(&structs.Loans{}).Where("id = ? AND state = 2", loan_id).Count(count).Error
}
