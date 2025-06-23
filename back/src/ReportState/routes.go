package reportstate

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	reportStates := router.Group("/report-states")
	{
		reportStates.GET("/", GetAllReportStatesController)
	}
}
