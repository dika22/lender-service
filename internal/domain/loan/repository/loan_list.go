package repository

import (
	"context"
)

func (r LoanRepository) GetAll(ctx context.Context, dest interface{}) error {
	return r.db.Find(dest).Error
}