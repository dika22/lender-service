package usecase

import (
	"context"
	"lender-service/package/di/repositories"
	"lender-service/package/structs"

	"github.com/hibiken/asynq"
)

type InvestorUsecase struct{
	repo repositories.WrappedRepositories
	workerClient *asynq.Client
}

type Investor interface {
	Invest(ctx context.Context, req structs.RequestInvestment) (structs.ResponseLoanInvestment,error)
}

func NewUsecaseInvestor(repo repositories.WrappedRepositories, workerClient *asynq.Client) Investor {
	return &InvestorUsecase{
		repo: repo,
		workerClient: workerClient,
	}
}
