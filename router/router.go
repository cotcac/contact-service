package router

import (
	"../controllers/contact"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	router := r.Group("/api")
	{

		// CATEGORY ENDPOINT.

		router.POST("/contact/", contact.Insert)
		router.GET("/contact/", contact.List)
		router.GET("/contact/:id", contact.Read)
		router.DELETE("/contact/:id", contact.Delete)

	}
	return r
}

// Router is...
func Router() *gin.Engine {
	return setupRouter()
}
