package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.POST("/", CreateUserController)
		users.GET("/", GetAllUsersController)
		users.GET("/:id", GetUserByIDController)
		users.PUT("/:id", UpdateUserController)
		users.DELETE("/:id", DeleteUserController)
	}
}
