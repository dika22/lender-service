package usecase

import (
	"context"
	"lender-service/package/structs"
)

func (u LoanUsecase) ListLoan(ctx context.Context) (structs.ResponseLoanList, error){
	dest := []structs.Loans{}
	if err := u.repo.DBrepo.Loan.GetAll(ctx, &dest); err != nil {
		return structs.ResponseLoanList{}, nil
	}
	return structs.ResponseLoanList{
		LoanList: dest,
	}, nil
	
}