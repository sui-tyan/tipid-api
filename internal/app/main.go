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

func RunApp(port ...string) {
	// create new app
	r := NewApp()

	// setup routes
	routes.SetupRoutes(r)

	customPort := ":8080"
	if len(port) > 0 {
		customPort = ":" + port[0]
	}

	// run the app
	r.Run(customPort)
}
