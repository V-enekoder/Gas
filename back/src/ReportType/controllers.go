package reporttype

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllReportTypesController(c *gin.Context) {
	reportTypes, err := GetAllReportTypesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve report types"})
		return
	}

	c.JSON(http.StatusOK, reportTypes)
}
