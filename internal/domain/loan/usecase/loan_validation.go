package usecase

import (
	"context"
	"lender-service/internal/constant"
	"lender-service/package/structs"
)

func (u LoanUsecase) ValidationLoan(ctx context.Context, req structs.RequestLoanValidator) (structs.ResponseLoan, error){
	count := int64(0)
	if err:= u.repo.DBrepo.Loan.Count(ctx, req.LoanID, &count); err !=nil{
		return structs.ResponseLoan{}, err
	}
	loanApproval := req.NewLoanValidator(constant.Approved)
	statusApproved := constant.StateApproved
	if count > 0 {
		return structs.ResponseLoan{
			LoanID: req.LoanID,
			State:  statusApproved,
		}, nil
	}
	err := u.repo.DBrepo.Loan.LoanApproval(ctx, loanApproval)
	if err != nil {
		return structs.ResponseLoan{}, err
	}
	err = u.repo.DBrepo.Loan.LoanUpdateState(ctx, req.LoanID, constant.Approved)
	if  err != nil {
		return structs.ResponseLoan{}, err
	}
	return structs.ResponseLoan{
		LoanID: req.LoanID,
		State:  statusApproved,
	}, nil
}