package reporttype

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	reportTypes := router.Group("/report-types")
	{
		reportTypes.GET("/", GetAllReportTypesController)
	}
}
