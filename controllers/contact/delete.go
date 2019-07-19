package contact

import (
	"github.com/gin-gonic/gin"
)

// Delete will delete a contact
func Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": "done",
	})
}
