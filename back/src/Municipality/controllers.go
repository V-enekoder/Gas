package municipality

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllMunicipalitiesController(c *gin.Context) {
	municipalities, err := GetAllMunicipalitiesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve municipalities"})
		return
	}

	c.JSON(http.StatusOK, municipalities)
}
