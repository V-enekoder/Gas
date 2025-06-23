package delivery

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	deliveries := router.Group("/deliveries")
	{
		deliveries.POST("/", CreateDeliveryController)
		deliveries.GET("/", GetAllDeliveriesController)
		deliveries.GET("/:id", GetDeliveryByIDController)
	}
}
