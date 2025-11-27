package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {

	// initialize expense routes
	expensesRoutes := router.Group("/expenses")
	cashFlowRoutes := router.Group("/cashflow")
	goalRoutes := router.Group("/goal")

	// add goal routes
	AddGoalRoutes(goalRoutes)

	// add expense routes
	AddExpenseRoutes(expensesRoutes)
	AddCashFlowRoutes(cashFlowRoutes)
}
