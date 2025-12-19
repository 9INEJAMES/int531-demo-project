package api

import (
	"database/sql"

	"github.com/9inejames/int531-demo-project/internal/middleware"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewApp(db *sql.DB) *fiber.App {

	// ===== metrics setup =====
	reg := prometheus.NewRegistry()
	metrics := NewMetrics(reg)

	app := fiber.New(fiber.Config{
		ErrorHandler: MetricsErrorHandler(metrics),
	})
	// ===== base middleware =====
	app.Use(middleware.RequestIDMiddleware)
	app.Use(middleware.LoggerMiddleware)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))
	app.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})

	// ===== endpoints =====
	app.Get("/metrics", adaptor.HTTPHandler(
		promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
	))

	app.Get("/health", HealthHandler(db, metrics))

	api := app.Group("/api")
	api.Get("/users", UsersHandler(db))
	api.Post("/users", CreateUserHandler(db))
	api.Get("/users/:id", GetUserHandler(db))
	api.Put("/users/:id", UpdateUserHandler(db))
	api.Delete("/users/:id", DeleteUserHandler(db))
	// ===== metrics middleware =====
	app.Use(MetricsMiddleware(metrics))

	return app
}
