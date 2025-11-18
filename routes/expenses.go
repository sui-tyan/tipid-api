package routes

import "github.com/gin-gonic/gin"

func AddExpenseRoutes(router *gin.RouterGroup) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "List of expenses"})
	})

	router.GET("/recent", func(c *gin.Context) {

		type ExpenseData struct {
			Label  string  `json:"label"`
			Amount float64 `json:"amount"`
		}

		recent_expenses := []ExpenseData{
			{Label: "Food", Amount: 150.75},
			{Label: "Travel", Amount: 80.50},
			{Label: "Guitar", Amount: 60.00},
			{Label: "Transaction Fees", Amount: 45.25},
			{Label: "Other", Amount: 120.00},
		}

		c.JSON(200, gin.H{"data": recent_expenses})
	})

}
