package routes

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gin-gonic/gin"
)

func AddExpenseRoutes(router *gin.RouterGroup) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "List of expenses"})
	})

	router.GET("/recent", func(c *gin.Context) {

		type ExpenseData struct {
			ExpenseLabel string  `json:"expenseLabel"`
			ExpenseValue float64 `json:"expenseValue"`
			Fill         string  `json:"fill"`
		}

		var recentExpenses []ExpenseData

		itemLength := gofakeit.Number(1, 30)

		for i := 0; i < itemLength; i++ {
			dummyData := ExpenseData{
				ExpenseLabel: gofakeit.ProductName(),
				ExpenseValue: float64(gofakeit.Number(100, 500)),
				Fill:         gofakeit.HexColor(),
			}
			recentExpenses = append(recentExpenses, dummyData)
		}

		c.JSON(200, gin.H{"data": recentExpenses})
	})

}
