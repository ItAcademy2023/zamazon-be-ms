// File: auth/cmd/main.go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sharat789/zamazon-be-ms/auth/internal/api"
	"github.com/sharat789/zamazon-be-ms/auth/internal/api/handlers"
	"github.com/sharat789/zamazon-be-ms/auth/internal/config"
	"github.com/sharat789/zamazon-be-ms/auth/internal/service"
	"github.com/sharat789/zamazon-be-ms/metrics"
	"log"
	"os"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize auth service
	authService := service.NewAuthService(cfg.JWTSecret)

	// Create Fiber app
	app := fiber.New()

	// Middleware
	app.Use(metrics.PrometheusMiddleware())
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)

	// Setup routes
	api.SetupRoutes(app, authHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082" // Default auth service port
	}

	log.Printf("Auth service starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
