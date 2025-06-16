package repository

import (
	"context"
	"lender-service/package/structs"
)

func (r InvestRepository) StoreInvest(ctx context.Context, payload structs.LoanInvestments) error {
	return r.db.Create(payload).Error
}