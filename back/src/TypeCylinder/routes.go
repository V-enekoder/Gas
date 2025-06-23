package typecylinder

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	typeCylinders := router.Group("/type-cylinders")
	{
		typeCylinders.GET("/", GetAllTypeCylindersController)
	}
}
