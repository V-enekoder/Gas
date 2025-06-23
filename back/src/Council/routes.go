package council

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	councils := router.Group("/councils")
	{
		councils.POST("/", CreateCouncilController)
		councils.GET("/", GetAllCouncilsController)
		councils.GET("/:id", GetCouncilByIDController)
		councils.PUT("/:id", UpdateCouncilController)
		councils.DELETE("/:id", DeleteCouncilController)
	}
}
