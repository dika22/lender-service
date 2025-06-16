package repository

import (
	"context"
	"lender-service/package/structs"
)

func (r LoanRepository) StoreLoan(ctx context.Context, dest structs.Loans) error {
	return r.db.Create(dest).Error
}