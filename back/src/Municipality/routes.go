package municipality

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	municipalities := router.Group("/municipalities")
	{
		municipalities.GET("/", GetAllMunicipalitiesController)
	}
}
