package app

import (
	"tipid-api/routes"

	"github.com/gin-gonic/gin"
)

func NewApp() *gin.Engine {
	// initialize gin engine
	r := gin.Default()

	return r
}

func RunApp() {
	// create new app
	r := NewApp()

	// setup routes
	routes.SetupRoutes(r)

	// run the app
	r.Run()
}
