package service

import (
	"pos/internal/core/domain"
	"pos/internal/core/port"

	"github.com/gin-gonic/gin"
)

type MarketService struct {
	repo port.MarketRepository
}

func NewMarketService(repo port.MarketRepository) *MarketService {
	return &MarketService{
		repo,
	}
}

func (ms *MarketService) Get(c *gin.Context) (domain.Market, error) {
	market, err := ms.repo.Get(c)
	if err != nil {
		return market, err
	}

	return market, nil
}

func (ms *MarketService) Send() {}
