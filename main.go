package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"softech.cloud/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})

	Album := new(controllers.Api)

	r.GET("/albums", Album.GetAlbums)
	r.GET("/albums/:id", Album.GetAlbumByID)
	r.POST("/albums", Album.CreateAlbum)

	// r.GET("/search", func(c *gin.Context) {
	// 	query := c.DefaultQuery("q", "none")
	// 	c.String(200, "Searched for: %s", query)
	// })

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
