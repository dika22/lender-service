package repository

import (
	"context"
)

func (r LoanRepository) GetByID(ctx context.Context, id_loan string, dest interface{}) error {
	return r.db.Where("id = ?", id_loan).Find(&dest).Error
}