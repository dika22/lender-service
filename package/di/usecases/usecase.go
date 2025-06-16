package usecases

import (
	in "lender-service/internal/domain/investor/usecase"
	le "lender-service/internal/domain/loan/usecase"
	"lender-service/package/di/repositories"

	"github.com/hibiken/asynq"
)

type Usecases struct {
	Loan le.ILoan
	Investor  in.Investor
}

func NewUsecase(repos repositories.WrappedRepositories, workerClient *asynq.Client) Usecases  {
	return Usecases{
		Loan: le.NewLoanUsecase(repos),
		Investor: in.NewUsecaseInvestor(repos, workerClient),
	}
}


