package usecase

import (
	"context"
	"lender-service/cmd/worker/tasks"
	"lender-service/package/structs"
)

func (u InvestorUsecase) Invest(ctx context.Context, req structs.RequestInvestment) (structs.ResponseLoanInvestment,error) {
	payload := req.NewInvestment()
	if err := u.repo.DBrepo.Investor.StoreInvest(ctx, payload); err != nil{
		return structs.ResponseLoanInvestment{}, err
	}

	job := tasks.LoanInvestments{
		LoanId: req.LoanID,
		InvestorId: req.InvestorID,
		Amount: req.Amount,
	}

	// Send job
	_, err := u.workerClient.EnqueueContext(ctx, job.Dispatch(), nil)
	if err != nil {
		return structs.ResponseLoanInvestment{}, err
	}
	return structs.ResponseLoanInvestment{
		LoanID: req.LoanID,
		CurrentInvestedAmount: req.Amount,
		State: "Invested",
	}, nil
}