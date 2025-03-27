package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "Hola Mundo",	
		})
	})

	router.Run("localhost:8080")
}