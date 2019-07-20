package contact

import (
	"fmt"
	"strconv"

	"../../models"
	"github.com/gin-gonic/gin"
)

// List of all contact
func List(c *gin.Context) {
	db := models.DBConn()
	perPage := 5
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr) // the err is necessary
	if err != nil {
		panic(err)
	}
	var skip int
	if page > 0 {
		skip = (page - 1) * perPage
	} else {
		skip = 0
	}
	stmt, err := db.Prepare(`
	select *
	from main
	limit ?,?`)
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query(skip, perPage+1)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "server error",
		})
	}
	fmt.Print(rows)
	all := make([]models.Contact, 0)
	for rows.Next() {
		row := models.Contact{}
		err := rows.Scan(&row.ID, &row.Name, &row.Email, &row.Content, &row.Date)
		if err != nil {
			panic(err)
		}
		all = append(all, row)
	}
	fmt.Println(len(all))
	var nextPage bool
	theLen := len(all)
	if theLen > perPage {
		nextPage = true
		all = all[:len(all)-1] // remote the last item.
	} else {
		nextPage = false
	}

	if err = rows.Err(); err != nil {
		c.JSON(500, gin.H{"error:": "server error"})
		return
	}
	c.JSON(200, gin.H{
		"page":   page,
		"result": all,
		"next":   nextPage,
	})
}
