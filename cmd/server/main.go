package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorunriki/akademiflow/pkg/config"
)

func main() {
	cfg := config.Load()
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"app":    cfg.AppName,
		})
	})

	port := cfg.AppPort
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	r.Run(":" + port)
}
