package paymentstate

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	paymentStates := router.Group("/payment-states")
	{
		paymentStates.GET("/", GetAllPaymentStatesController)
	}
}
