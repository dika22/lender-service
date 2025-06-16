package usecase

import (
	"context"
	"lender-service/internal/constant"
	"lender-service/package/structs"
)

func (u LoanUsecase) LoanDisbursement(ctx context.Context, req structs.RequestLoanValidator) (structs.ResponseLoan, error){
	loanDisbursement := req.NewLoanValidator(constant.Disbursed)
	err := u.repo.DBrepo.Loan.DisbursementLoan(ctx, loanDisbursement)
	if err != nil {
		return structs.ResponseLoan{}, err
	}
	return structs.ResponseLoan{
		LoanID: req.LoanID,
		State:  constant.StateDisbursed,
	}, nil
}