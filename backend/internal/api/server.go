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

func NewApp(db *sql.DB, reg *prometheus.Registry) *fiber.App {
	app := fiber.New()

	// สร้าง Metrics object
	m := NewMetrics()

	// register metrics เฉพาะตอน setup จริง
	if reg != nil {
		if err := reg.Register(m.HttpRequestsTotal); err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				m.HttpRequestsTotal = are.ExistingCollector.(*prometheus.CounterVec)
			}
		}
		if err := reg.Register(m.HttpRequestDuration); err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				m.HttpRequestDuration = are.ExistingCollector.(*prometheus.HistogramVec)
			}
		}

		app.Get("/metrics", adaptor.HTTPHandler(promhttp.HandlerFor(reg, promhttp.HandlerOpts{})))
	}

	// middleware & routes
	app.Use(middleware.RequestIDMiddleware)
	app.Use(middleware.LoggerMiddleware)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))

	app.Get("/health", HealthHandler(db, m))
	app.Use(MetricsMiddleware(m))

	api := app.Group("/api")
	api.Get("/users", UsersHandler(db))
	api.Post("/users", CreateUserHandler(db))
	api.Get("/users/:id", GetUserHandler(db))
	api.Put("/users/:id", UpdateUserHandler(db))
	api.Delete("/users/:id", DeleteUserHandler(db))

	return app
}
