package tasks

import (
	"context"
	"lender-service/package/config"
	"lender-service/package/di/repositories"

	"github.com/hibiken/asynq"
)

const (
	TaskInvestLoan = "invest:loan"
)

type AsyncTask struct{
	conf *config.Config
	repos repositories.WrappedRepositories
	workerClient *asynq.Client
}

type Tasks interface {
	InvestLoan(ctx context.Context, task *asynq.Task) error
}

func NewAsynqTask(conf *config.Config, repos repositories.WrappedRepositories, workClient *asynq.Client) Tasks {
	return AsyncTask{
		conf: conf,
		repos: repos,
		workerClient: workClient,
	}
}