package payment

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	payments := router.Group("/payments")
	{
		payments.POST("/", CreatePaymentController)
		payments.GET("/", GetAllPaymentsController)
		payments.GET("/:id", GetPaymentByIDController)
	}
}
