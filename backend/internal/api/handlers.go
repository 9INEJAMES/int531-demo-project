package api

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func HealthHandler(db *sql.DB, metrics *Metrics) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 1*time.Second)
		defer cancel()

		if err := db.PingContext(ctx); err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"status": "unhealthy",
			})
		}

		return c.JSON(fiber.Map{"status": "ok"})
	}
}

func UsersHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
		defer cancel()

		rows, err := db.QueryContext(
			ctx,
			`SELECT id, name, created_at FROM users ORDER BY id`,
		)
		if err != nil {
			log.Printf("list users error: %v", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		defer rows.Close()

		users := []User{}

		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Name, &u.CreatedAt); err != nil {
				log.Printf("scan user error: %v", err)
				return c.SendStatus(fiber.StatusInternalServerError)
			}
			users = append(users, u)
		}

		return c.JSON(users)
	}
}

func CreateUserHandler(db *sql.DB) fiber.Handler {
	type request struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	return func(c *fiber.Ctx) error {
		var req request
		if err := c.BodyParser(&req); err != nil || req.ID == "" || req.Name == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}

		ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
		defer cancel()

		var id string
		err := db.QueryRowContext(
			ctx,
			`INSERT INTO users (id, name) VALUES ($1, $2) RETURNING id`,
			req.ID,
			req.Name,
		).Scan(&id)

		if err != nil {
			log.Printf("create user error: %v", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"id": id,
		})
	}
}

func GetUserHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
		defer cancel()

		var u User
		err := db.QueryRowContext(
			ctx,
			`SELECT id, name, created_at FROM users WHERE id = $1`,
			id,
		).Scan(&u.ID, &u.Name, &u.CreatedAt)

		if err == sql.ErrNoRows {
			return c.SendStatus(fiber.StatusNotFound)
		}
		if err != nil {
			log.Printf("get user error: %v", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(u)
	}
}

func UpdateUserHandler(db *sql.DB) fiber.Handler {
	type request struct {
		Name string `json:"name"`
	}

	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		var req request
		if err := c.BodyParser(&req); err != nil || req.Name == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
		defer cancel()

		res, err := db.ExecContext(
			ctx,
			`UPDATE users SET name = $1 WHERE id = $2`,
			req.Name,
			id,
		)
		if err != nil {
			log.Printf("update user error: %v", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		rows, _ := res.RowsAffected()
		if rows == 0 {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

func DeleteUserHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
		defer cancel()

		res, err := db.ExecContext(
			ctx,
			`DELETE FROM users WHERE id = $1`,
			id,
		)
		if err != nil {
			log.Printf("delete user error: %v", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		rows, _ := res.RowsAffected()
		if rows == 0 {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

