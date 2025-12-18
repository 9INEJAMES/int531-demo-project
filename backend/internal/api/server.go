package api

import (
	"database/sql"

	"github.com/9inejames/int531-demo-project/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewApp(db *sql.DB) *fiber.App {
	app := fiber.New()

	// middleware
	app.Use(middleware.RequestIDMiddleware)
	app.Use(middleware.LoggerMiddleware)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))

	// health
	app.Get("/health", HealthHandler(db))

	// api group
	api := app.Group("/api")
	api.Get("/users", UsersHandler(db))
	api.Post("/users", CreateUserHandler(db))
	api.Get("/users/:id", GetUserHandler(db))
	api.Put("/users/:id", UpdateUserHandler(db))
	api.Delete("/users/:id", DeleteUserHandler(db))

	return app
}
