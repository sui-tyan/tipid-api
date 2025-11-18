package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {

	// initialize expense routes
	expensesRoutes := router.Group("/expenses")

	// add expense routes
	AddExpenseRoutes(expensesRoutes)
}
