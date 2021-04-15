package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatedID(c *gin.Context, id uint) {
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": message})
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": message})
}
