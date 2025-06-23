package disabled

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateDisabledController(c *gin.Context) {
	var dto DisabledCreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := CreateDisabledService(dto)
	if err != nil {
		if errors.Is(err, ErrUserAlreadyHasDisabilityRecord) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		// This can also fail if the UserID does not exist (foreign key violation)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create disability record"})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func GetAllDisabledController(c *gin.Context) {
	disabledList, err := GetAllDisabledService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve disability records"})
		return
	}

	c.JSON(http.StatusOK, disabledList)
}

func GetDisabledByIDController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	disabled, err := GetDisabledByIDService(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Disability record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve disability record"})
		return
	}

	c.JSON(http.StatusOK, disabled)
}

func UpdateDisabledController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var dto DisabledUpdateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedDisabled, err := UpdateDisabledService(uint(id), dto)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Disability record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update disability record"})
		return
	}

	c.JSON(http.StatusOK, updatedDisabled)
}

func DeleteDisabledController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := DeleteDisabledService(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete disability record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Disability record deleted successfully"})
}
