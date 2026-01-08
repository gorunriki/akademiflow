package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorunriki/akademiflow/internal/modules/auth"
	"github.com/gorunriki/akademiflow/internal/modules/users"
	"github.com/gorunriki/akademiflow/internal/shared/middleware"
	"github.com/gorunriki/akademiflow/pkg/config"
	"github.com/gorunriki/akademiflow/pkg/database"
)

func main() {
	// init .env
	cfg := config.Load()

	// init DB
	db := database.Connect(cfg)

	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo, cfg)
	authHandler := auth.NewHandler(authService)

	userRepo := users.NewRepository(db)
	userService := users.NewService(userRepo)
	userHandler := users.NewHandler(userService)

	// DB migrate
	database.Migrate(db)
	database.Seed(db)

	//init Gin
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"app":    cfg.AppName,
		})
	})

	r.POST("/register", userHandler.Register)
	r.POST("/login", authHandler.Login)

	api := r.Group("/api")
	api.Use(middleware.Auth(cfg))

	adminOnly := api.Group("/admin")
	adminOnly.Use(middleware.RequiredRole("admin"))
	adminOnly.GET("/dashboard", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome admin",
		})
	})
	adminOnly.GET("/users", userHandler.List)

	api.GET("/me", userHandler.Me)

	port := cfg.AppPort
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	r.Run(":" + port)
}
