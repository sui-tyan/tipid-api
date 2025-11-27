package routes

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gin-gonic/gin"
)

func AddCashFlowRoutes(router *gin.RouterGroup) {

	router.GET("/recent", func(c *gin.Context) {

		type CashFlowData struct {
			Inflow  float64 `json:"inflow"`
			Outflow float64 `json:"outflow"`
		}

		recentCashFlow := CashFlowData{Inflow: float64(gofakeit.Price(100, 50000)), Outflow: float64(gofakeit.Price(100, 50000))}

		c.JSON(200, gin.H{"data": recentCashFlow})
	})
}
