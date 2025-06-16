package tasks

import (
	"context"
	"encoding/json"
	"errors"
	"lender-service/internal/constant"
	"lender-service/package/structs"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type LoanInvestments struct {
    LoanId 		string `json:"loan_id"`
    InvestorId 	string `json:"investor_id"`
    Amount 		int64  `json:"amount"`
}

func (i LoanInvestments) Dispatch() *asynq.Task {
	marshal, err := json.Marshal(i)
	if err != nil {
		return nil
	}
	return asynq.NewTask(
		TaskInvestLoan, 
		marshal,
		asynq.Queue("lender:queue:low"),
	)
}

func (a AsyncTask) InvestLoan(ctx context.Context, task *asynq.Task) error {
	var param LoanInvestments
	if err := json.Unmarshal(task.Payload(), &param); err != nil {
		return err
	}


	dest := structs.Loans{}
	if err := a.repos.DBrepo.Loan.GetByID(ctx, param.LoanId, dest); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	if dest.PrincipalAmount == param.Amount {
		if err := a.repos.DBrepo.Loan.LoanUpdateState(ctx, param.LoanId, constant.Invested); err != nil {
			return err
		}
	} else {
		destInvestment := []structs.LoanInvestments{}
		if err := a.repos.DBrepo.Investor.GetByLoanID(ctx, param.InvestorId, param.LoanId, &destInvestment); err != nil {
			return err
		}

		totalAmount := int64(0)
		for _, inv := range destInvestment {
			totalAmount += inv.Amount
		}

		if totalAmount == param.Amount {
			if err := a.repos.DBrepo.Loan.LoanUpdateState(ctx, param.LoanId, constant.Invested); err != nil {
				return err
			}
		}
	}

	return nil
}