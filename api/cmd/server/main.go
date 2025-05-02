package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/eijiosakabe/ai-todos/api/internal/infrastructure/database"
	infraRepos "github.com/eijiosakabe/ai-todos/api/internal/infrastructure/repositories"
	"github.com/eijiosakabe/ai-todos/api/internal/presentation/handlers"
	"github.com/eijiosakabe/ai-todos/api/internal/services"
)

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Create Gin router
	r := gin.Default()

	// Configure CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Database configuration
	dbConfig := database.GetDatabaseConfigFromEnv()
	db, err := database.NewPostgresConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run database migrations
	if err := database.MigrateDatabase(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Setup repositories
	userRepo := infraRepos.NewUserRepository(db)
	userSessionRepo := infraRepos.NewPostgresUserSessionRepository(db)

	// Setup services
	authService := services.NewAuthService(userRepo, userSessionRepo)

	// API Group
	authGroup := r.Group("/api/auth")
	{
		handlers.SetupAuthRoutes(authGroup, authService)
	}

	// Get port from environment or default to 4200
	port := os.Getenv("PORT")
	if port == "" {
		port = "4200"
	}

	// Start server
	log.Printf("Starting server on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
