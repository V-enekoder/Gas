package reportstate

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllReportStatesController(c *gin.Context) {
	reportStates, err := GetAllReportStatesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve report states"})
		return
	}

	c.JSON(http.StatusOK, reportStates)
}
