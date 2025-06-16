package usecase

import (
	"context"
	"lender-service/package/structs"
)


func (u LoanUsecase) LoanDetail(ctx context.Context, idLoan string) (structs.Loans, error){
	dest := structs.Loans{}
	if err := u.repo.DBrepo.Loan.GetByID(ctx, idLoan, &dest); err != nil {
		return structs.Loans{}, nil
	}
	return dest, nil
}