package contact

import (
	"fmt"

	"../../models"
	"github.com/gin-gonic/gin"
)

// Read a single row
func Read(c *gin.Context) {

	id := c.Param("id")
	var row models.Contact
	db := models.DBConn()
	stmt, err := db.Prepare("select * from main where id=?")
	if err != nil {
		panic(err)
	}
	err = stmt.QueryRow(id).Scan(&row.ID, &row.Name, &row.Email, &row.Content, &row.Date)
	if err != nil {
		panic(err)
	}
	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{
			"message": "Not Found!",
		})
		return
	}
	c.JSON(200, gin.H{
		"result": row,
	})
}
