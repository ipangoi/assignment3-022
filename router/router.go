package routers

import (
	"assignment3/api"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	router := gin.Default()

	// router.GET("/weather", api.Get)
	router.PUT("/weather", api.Update)

	return router
}
