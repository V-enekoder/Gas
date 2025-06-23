package order

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	orders := router.Group("/orders")
	{
		orders.POST("/", CreateOrderController)
		orders.GET("/", GetAllOrdersController)
		orders.GET("/:id", GetOrderByIDController)
	}
}
