package worker

import (
	"context"
	"fmt"
	"lender-service/cmd/worker/tasks"
	"lender-service/internal/constant"
	"lender-service/package/config"
	"lender-service/package/di/repositories"
	"log"

	"github.com/hibiken/asynq"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/urfave/cli/v2"
)


type (
	WorkerHandler func(context.Context, *asynq.Task) error
	Worker struct {
		conf *config.Config
		cacheConf *config.Cache
		tasks tasks.Tasks
		nr *newrelic.Application
	}
)

const CmdServerWorker = "start-worker"
func WorkerClient(cf *config.Cache) *asynq.Client{
	return asynq.NewClient(asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%s", cf.WorkerRedisHost, cf.WorkerRedisPort)})
}

func (w Worker) StartWorker(*cli.Context) error  {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{},
		asynq.Config{},
	)
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TaskInvestLoan, w.NrWorkerMiddleware(w.tasks.InvestLoan, tasks.TaskInvestLoan))
	if err := srv.Run(mux); err!= nil {
		log.Fatalf("could not run serve: %v", err)
		return err
	}
	return nil
}

func (w Worker) NrWorkerMiddleware(f WorkerHandler, taskName string) func(ctx context.Context, task *asynq.Task) error {
	return func(ctx context.Context, task *asynq.Task) error {
		txn := w.nr.StartTransaction(taskName)
		defer txn.End()
		ctx = context.WithValue(ctx, constant.NewRelicTransactionCtx, txn)
		return f(ctx, task)
	}
 }
 

func StartWorker(conf *config.Config, c *config.Cache, repos repositories.WrappedRepositories, workerClient *asynq.Client) []*cli.Command {
	w := Worker{conf: conf,cacheConf: c}
	return []*cli.Command{
		{
			Name: CmdServerWorker,
			Usage: "Serve Lender Worker",
			Action: w.StartWorker,
		},
	}
}