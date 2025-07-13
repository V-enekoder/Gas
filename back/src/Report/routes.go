package report

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	reports := router.Group("/reports")
	{
		reports.POST("/", CreateReportController)
		reports.GET("/", GetAllReportsController)
		reports.GET("/:id", GetReportByIDController)
		reports.GET("/:id/user", GetReportByUserIDController)
	}
}
