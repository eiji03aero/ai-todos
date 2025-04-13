package main

import (
	"github.com/eijiosakabe/ai-todos/api/internal/presentation/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register routes
	handlers.RegisterRoutes(r)

	// Start the server
	r.Run(":4200")
}
