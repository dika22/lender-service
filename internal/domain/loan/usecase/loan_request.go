package usecase

import (
	"context"
	"lender-service/package/structs"
)

func (u LoanUsecase) LoanRequest(ctx context.Context, req *structs.LoanRequest) (*structs.ResponseLoan, error)  {
	loanPayload := req.NewLoan()
	err := u.repo.DBrepo.Loan.StoreLoan(ctx, loanPayload)
	if err != nil {
		return nil, err
	}
	return &structs.ResponseLoan{
		LoanID: loanPayload.Id,
		State: "Purposed",
	}, nil
}