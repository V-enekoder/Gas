package orderstate

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	orderStates := router.Group("/order-states")
	{
		orderStates.GET("/", GetAllOrderStatesController)
	}
}
