package contact

import (
	"github.com/gin-gonic/gin"
)

// Read a single contact row
func Read(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": "work",
	})
}
