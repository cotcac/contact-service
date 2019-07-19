package contact

import (
	"github.com/gin-gonic/gin"
)

// List will return list of all contact
func List(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": "work",
	})
}
