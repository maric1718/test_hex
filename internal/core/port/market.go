package port

import (
	"pos/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type MarketService interface {
	Get(c *gin.Context) (domain.Market, error)
	Send()
}

type MarketRepository interface {
	Get(c *gin.Context) (domain.Market, error)
	Send()
}
