package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Profile struct {
	Id    int16  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var profiles []Profile

func init() {
	profiles = []Profile{
		{Id: 1, Title: "title 1", Body: "body 1"},
		{Id: 2, Title: "title 2", Body: "body 2"},
		{Id: 3, Title: "title 3", Body: "body 3"},
	}
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods: []string{
			"POST",
			"GET",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		AllowOrigins: []string{
			"http://localhost:3000",
		},
	}))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/events", func(c *gin.Context) {
			c.JSON(http.StatusOK, profiles)
		})

		v1.POST(("/new"), func(c *gin.Context) {
			profile := Profile{Id: 4, Title: "title 4", Body: "body 4"}
			profiles = append(profiles, profile)
			c.JSON(http.StatusOK, profile)
		})
	}

	r.Run(":3001")
}
