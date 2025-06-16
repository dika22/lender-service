package repositories

import (
	invRepo "lender-service/internal/domain/investor/repository"
	loanRepo "lender-service/internal/domain/loan/repository"
	"lender-service/package/config"
	"lender-service/package/connection/cache"

	"gorm.io/gorm"
)

type WrappedRepositories struct {
	DBrepo SQLRepositories
	RedisRepo RedisRepositories
}

type SQLRepositories struct {
	Loan loanRepo.LoanDatabase
	Investor invRepo.InvestDatabase
}


func NewWrappedRepositories(cr SQLRepositories, rr RedisRepositories) WrappedRepositories {
	return WrappedRepositories{
		DBrepo: cr,
		RedisRepo: rr,
	}
}

func NewDatabaseRepositories(g *gorm.DB) SQLRepositories {
	return SQLRepositories{
		Loan: loanRepo.NewLoanRepository(g),
		Investor: invRepo.NewInvestRepository(g),
	}
}

type RedisRepositories struct{
	loanCache loanRepo.LoanRedisRepository
}

func NewCacheRepositories(conf *config.Config, cache cache.Cache, repo SQLRepositories) RedisRepositories {
	return RedisRepositories{
		loanCache: cache,
	}
}