package council

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCouncilController(c *gin.Context) {
	var dto CouncilCreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := CreateCouncilService(dto)
	if err != nil {
		if errors.Is(err, ErrUserAlreadyHasCouncilRecord) || errors.Is(err, ErrLeaderDocumentExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create council record"})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func GetAllCouncilsController(c *gin.Context) {
	councils, err := GetAllCouncilsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve council records"})
		return
	}

	c.JSON(http.StatusOK, councils)
}

func GetCouncilByIDController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	council, err := GetCouncilByIDService(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Council record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve council record"})
		return
	}

	c.JSON(http.StatusOK, council)
}

func UpdateCouncilController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var dto CouncilUpdateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCouncil, err := UpdateCouncilService(uint(id), dto)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Council record not found"})
			return
		}
		if errors.Is(err, ErrLeaderDocumentExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update council record"})
		return
	}

	c.JSON(http.StatusOK, updatedCouncil)
}

func DeleteCouncilController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := DeleteCouncilService(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete council record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Council record deleted successfully"})
}
