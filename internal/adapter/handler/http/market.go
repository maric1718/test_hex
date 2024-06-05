package http

import (
	"pos/internal/core/port"

	"github.com/gin-gonic/gin"
)

type MarketHandler struct {
	svc port.MarketService
}

func NewMarketHandler(svc port.MarketService) *MarketHandler {
	return &MarketHandler{
		svc,
	}
}

func (mh *MarketHandler) GetMarkets(c *gin.Context) {
	// payload := []struct {
	// 	ID       string
	// 	Name     string
	// 	Status   int
	// 	Outcomes []struct {
	// 		ID     string
	// 		Name   string
	// 		Status int
	// 	}
	// }{}

	// data, err := mh.svc.Get(c, req.ID)
	// if err != nil {
	// 	// handleError(ctx, err)
	// 	return
	// }

	// rsp := newOrderResponse(order)

	// handleSuccess(c, rsp)
}
