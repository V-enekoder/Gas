package orderstate

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOrderStatesController(c *gin.Context) {
	orderStates, err := GetAllOrderStatesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order states"})
		return
	}

	c.JSON(http.StatusOK, orderStates)
}
