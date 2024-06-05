package repository

import (
	"encoding/json"
	"os"
	"pos/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type MarketAdapter struct{}

func NewMarketRepository() *MarketAdapter {
	return &MarketAdapter{}
}

func (mr *MarketAdapter) Get(c *gin.Context) (domain.Market, error) {
	market := domain.Market{}

	path := "internal/adapter/storage/file_system/json_data/initial/events.json"

	content, err := os.ReadFile(path)
	if err != nil {
		return market, err
	}

	if err := json.Unmarshal(content, &market); err != nil {
		return market, err
	}

	return market, nil

}

func (mr *MarketAdapter) Send() {

}
