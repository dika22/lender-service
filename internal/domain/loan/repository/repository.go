package repository

import (
	"context"
	"lender-service/package/structs"

	"gorm.io/gorm"
)

type LoanDatabase interface{
	// StoreDocument(ctx context.Context) error
	StoreLoan(ctx context.Context, payload structs.Loans) error
	LoanApproval(ctx context.Context, payload structs.LoanHistory) error
	GetAll(ctx context.Context, dest interface{}) error
	GetByID(ctx context.Context, loan_id string, dest interface{}) error
	DisbursementLoan(ctx context.Context, payload structs.LoanHistory) error
	Count(ctx context.Context, loan_id string, count *int64) error
	LoanUpdateState(ctx context.Context, loanID string, state int) error 
}

type LoanRepository struct{
	db *gorm.DB
}

type LoanRedisRepository interface {

}

func NewLoanRepository(conn *gorm.DB) LoanDatabase {
	return LoanRepository{
		db: conn,
	}
}