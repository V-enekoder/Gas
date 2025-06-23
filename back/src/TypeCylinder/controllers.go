package typecylinder

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTypeCylindersController(c *gin.Context) {
	typeCylinders, err := GetAllTypeCylindersService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve type cylinders"})
		return
	}

	c.JSON(http.StatusOK, typeCylinders)
}
