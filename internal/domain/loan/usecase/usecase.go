package usecase

import (
	"context"
	"lender-service/package/di/repositories"
	"lender-service/package/structs"
)

type ILoan interface{
	UploadDocument(ctx context.Context) error
	LoanRequest(ctx context.Context, req *structs.LoanRequest) (*structs.ResponseLoan, error)
	ValidationLoan(ctx context.Context, req structs.RequestLoanValidator) (structs.ResponseLoan, error)
	ListLoan(ctx context.Context) (structs.ResponseLoanList, error)
	LoanDetail(ctx context.Context, idLoan string) (structs.Loans, error)
	LoanDisbursement(ctx context.Context, req structs.RequestLoanValidator) (structs.ResponseLoan, error)
}

type LoanUsecase struct{
	repo repositories.WrappedRepositories
}


func NewLoanUsecase(repos repositories.WrappedRepositories) ILoan  {
	return &LoanUsecase{
		repo: repos,
	}
}
