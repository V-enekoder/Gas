package disabled

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	disabled := router.Group("/disabled")
	{
		disabled.POST("/", CreateDisabledController)
		disabled.GET("/", GetAllDisabledController)
		disabled.GET("/:id", GetDisabledByIDController)
		disabled.PUT("/:id", UpdateDisabledController)
		disabled.DELETE("/:id", DeleteDisabledController)
	}
}
