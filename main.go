package main

import (
	"lender-service/package/config"
	"lender-service/package/connection/cache"
	"lender-service/package/connection/database"
	"lender-service/package/di/repositories"
	"lender-service/package/di/usecases"
	"lender-service/package/validator"
	"log"
	"os"

	api "lender-service/cmd/api"
	"lender-service/cmd/worker"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/urfave/cli/v2"
)

func main() {

  dbConf := config.NewDatabase()
  conf := config.NewConfig()
  cacheConf := config.NewCache()
  conn := database.LenderDB
  dbConn := database.NewDatabase(conn, dbConf)

  rds := cache.NewRedis("cacheLender", cacheConf)

  nrApp, err := newrelic.NewApplication(
    newrelic.ConfigAppName("fusio"),
    newrelic.ConfigLicense(conf.NewRelicLicense),
)
if err != nil {
    log.Print("ERROR INIT NEWRELIC", err)
}


  workerClient := worker.WorkerClient(cacheConf)
  dbRepo := repositories.NewDatabaseRepositories(dbConn)
  redisRepo := repositories.NewCacheRepositories(conf, rds, dbRepo)
  repos := repositories.NewWrappedRepositories(dbRepo, redisRepo)
  usecase := usecases.NewUsecase(repos, workerClient)
  validate := validator.NewValidator()
  cmds := []*cli.Command{}
  cmds = append(cmds, api.ServeAPI(usecase, validate, cacheConf)...)
  cmds = append(cmds, worker.StartWorker(conf, cacheConf, repos, workerClient, nrApp)...)

  app := &cli.App{
    Name: "lender-service",
    Commands: cmds,
  }

  if err := app.Run(os.Args); err != nil {
    panic(err)
  }
}