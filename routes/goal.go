package routes

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gin-gonic/gin"
)

func AddGoalRoutes(router *gin.RouterGroup) {

	router.GET("/", func(c *gin.Context) {

		type GoalData struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
			Fill  string  `json:"fill"`
		}

		goalAmount := float64(gofakeit.Number(1000, 10000))
		currentAmount := float64(gofakeit.Number(500, int(goalAmount)))

		goalItems := []GoalData{
			{Type: "Current", Value: currentAmount, Fill: "#90be6d"},
			{Type: "Goal", Value: goalAmount, Fill: "#d4d4d4ff"},
		}

		c.JSON(200, gin.H{"data": goalItems})
	})

}
