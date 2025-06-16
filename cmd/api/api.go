package cmd

import (
	"context"
	"fmt"
	deliInvestor "lender-service/internal/domain/investor/delivery"
	"lender-service/internal/domain/loan/delivery"
	"lender-service/package/config"
	"lender-service/package/di/usecases"
	"lender-service/package/validator"
	"net/http"
	"os"
	"time"

	"os/signal"

	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/urfave/cli/v2"
)

const CmdServeHTTP = "serve-http"

type HTTP struct{
	usecase usecases.Usecases
	validate *validator.Validator
	cacheConf *config.Cache
}

func (h HTTP) ServeAPI(c *cli.Context) error  {
	e := echo.New();

	mon := asynqmon.New(asynqmon.Options{
		RootPath: "/monitoring/tasks",
		RedisConnOpt: asynq.RedisClientOpt{
			Addr: fmt.Sprintf("%s:%s", h.cacheConf.WorkerRedisHost, h.cacheConf.WorkerRedisPort),
		},
	})
	e.Any("/monitoring/tasks/*", echo.WrapHandler(mon))

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	loanAPI := e.Group("lender/api/v1")
	loanAPI.Use(echoMiddleware.Logger())

	delivery.NewLoanHTTP(loanAPI, h.usecase, h.validate)
	deliInvestor.NewInvestorHTTP(loanAPI, h.usecase, h.validate)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(fmt.Sprintf(":%v", 3000)); err != nil {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	return nil	
}

func ServeAPI(usecases usecases.Usecases, v *validator.Validator, cf *config.Cache) []*cli.Command {
	h := &HTTP{usecase: usecases, validate: v, cacheConf: cf}
	return []*cli.Command{
		{
			Name: CmdServeHTTP,
			Usage: "Serve Document Service",
			Action: h.ServeAPI,
		},
	}
}