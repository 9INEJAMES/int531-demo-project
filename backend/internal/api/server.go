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

var metrics *Metrics

func GetMetrics() *Metrics {
	if metrics == nil {
		metrics = NewMetrics(prometheus.DefaultRegisterer.(*prometheus.Registry))
	}
	return metrics
}

func NewApp(db *sql.DB) *fiber.App {
	// ===== metrics setup =====
	metrics := GetMetrics() // ใช้ instance เดียว
	app := fiber.New()

	// ===== base middleware =====
	app.Use(middleware.RequestIDMiddleware)
	app.Use(middleware.LoggerMiddleware)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))

	app.Use(MetricsMiddleware(metrics))

	app.Get("/metrics", adaptor.HTTPHandler(
		promhttp.Handler(),
	))

	app.Get("/health", HealthHandler(db, metrics))

	api := app.Group("/api")
	api.Get("/users", UsersHandler(db))
	api.Post("/users", CreateUserHandler(db))
	api.Get("/users/:id", GetUserHandler(db))
	api.Put("/users/:id", UpdateUserHandler(db))
	api.Delete("/users/:id", DeleteUserHandler(db))

	return app
}
