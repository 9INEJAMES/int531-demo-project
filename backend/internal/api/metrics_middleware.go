package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func MetricsMiddleware(m *Metrics) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()

		route := "unknown"
		if r := c.Route(); r != nil && r.Path != "" {
			route = r.Path
		}

		m.HttpRequestDuration.
			WithLabelValues(c.Method(), route).
			Observe(time.Since(start).Seconds())

		return err
	}
}
