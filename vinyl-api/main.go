package main

import (
	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json: "id"`
	Title string `json: "title"`
	Artist string `json: "artist"`
	Price float64 `json: "price"`
}

var albums = []album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "John Coltrane", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 90.99},
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "Hola Mundo",	
		})
	})

	router.Run("localhost:8080")
}