package redis

import (
	"lender-service/internal/domain/loan/repository"
	"lender-service/package/connection/cache"
)

type LoanCache struct {
	cache cache.Cache
}



func NewLoanCache(cache cache.Cache) repository.LoanRedisRepository {
	return LoanCache{}
}