package paymentstate

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPaymentStatesController(c *gin.Context) {
	paymentStates, err := GetAllPaymentStatesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve payment states"})
		return
	}

	c.JSON(http.StatusOK, paymentStates)
}
