package commerce

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCommerceController(c *gin.Context) {
	var dto CommerceCreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := CreateCommerceService(dto)
	if err != nil {
		if errors.Is(err, ErrUserAlreadyHasCommerceRecord) || errors.Is(err, ErrRifExists) || errors.Is(err, ErrBossDocumentExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create commerce record"})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func GetAllCommercesController(c *gin.Context) {
	commerces, err := GetAllCommercesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve commerce records"})
		return
	}

	c.JSON(http.StatusOK, commerces)
}

func GetCommerceByIDController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	commerce, err := GetCommerceByIDService(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Commerce record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve commerce record"})
		return
	}

	c.JSON(http.StatusOK, commerce)
}

func UpdateCommerceController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var dto CommerceUpdateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCommerce, err := UpdateCommerceService(uint(id), dto)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Commerce record not found"})
			return
		}
		if errors.Is(err, ErrRifExists) || errors.Is(err, ErrBossDocumentExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update commerce record"})
		return
	}

	c.JSON(http.StatusOK, updatedCommerce)
}

func DeleteCommerceController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := DeleteCommerceService(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete commerce record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Commerce record deleted successfully"})
}
