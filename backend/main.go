package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type WelcomeData struct {
	Message string `json:"message"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	{
		api.GET("/welcome", getWelcome)
	}

	r.Run(":8080")
}

func getWelcome(c *gin.Context) {
	data := WelcomeData{
		Message: "Welcome to your Go Vue Boilerplate!",
		Version: "1.0.0",
		Status:  "Server is running successfully",
	}
	c.JSON(http.StatusOK, data)
}
