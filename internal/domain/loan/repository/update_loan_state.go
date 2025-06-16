package repository

import (
	"context"
)

func (r LoanRepository) LoanUpdateState(ctx context.Context, loanID string, state int) error {
	return r.db.Table("loans").Where("id = ?", loanID).UpdateColumn("state", state).Error
}