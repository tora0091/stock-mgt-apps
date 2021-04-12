package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/v1/events", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello gin world.",
		})
	})
	r.Run(":3001")
}
