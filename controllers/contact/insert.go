package contact

import (
	"fmt"

	"../../models"
	"github.com/gin-gonic/gin"
)

// Insert contact data
func Insert(c *gin.Context) {
	db := models.DBConn()
	var json models.Contact
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}
	insert, err := db.Prepare("INSERT INTO `main` (`email`, `name`, `content`) VALUES (?, ?, ?);")
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"message3": "error 1",
		})
		return
	}
	_, err = insert.Exec(json.Email, json.Name, json.Content)
	if err != nil {
		panic(err)
	}
	res1 := models.Response{
		Success: true,
		Message: "inserted",
	}

	fmt.Println(res1)
	c.JSON(200, res1)
}
