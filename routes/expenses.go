package routes

import "github.com/gin-gonic/gin"

func AddExpenseRoutes(router *gin.RouterGroup) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "List of expenses"})
	})

	router.GET("/recent", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "List of expenses"})
	})

}
