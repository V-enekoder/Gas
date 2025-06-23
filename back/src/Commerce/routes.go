package commerce

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	commerces := router.Group("/commerces")
	{
		commerces.POST("/", CreateCommerceController)
		commerces.GET("/", GetAllCommercesController)
		commerces.GET("/:id", GetCommerceByIDController)
		commerces.PUT("/:id", UpdateCommerceController)
		commerces.DELETE("/:id", DeleteCommerceController)
	}
}
