package repository

import (
	"context"
	"lender-service/package/structs"

	"gorm.io/gorm"
)

type InvestRepository struct {
	db *gorm.DB
}

type InvestDatabase interface{
	StoreInvest(ctx context.Context, payload structs.LoanInvestments) error
	GetByLoanID(ctx context.Context, investorID string, loanId string, dest interface{}) error
}


func NewInvestRepository(g *gorm.DB) InvestDatabase {
	return &InvestRepository{
		db: g,
	}
}