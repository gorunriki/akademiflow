package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorunriki/akademiflow/pkg/config"
	"github.com/gorunriki/akademiflow/pkg/database"
)

func main() {
	// init .env
	cfg := config.Load()

	// init DB
	db := database.Connect(cfg)
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// DB migrate
	database.Migrate(db)

	//init Gin
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
