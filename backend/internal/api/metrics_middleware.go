package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func MetricsMiddleware(m *Metrics) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()

		if c.Path() == "/metrics" || c.Path() == "/health" {
			return err
		}

		route := "unknown"
		if r := c.Route(); r != nil && r.Path != "" {
			route = r.Path
		}

		// status code
		statusCode := c.Response().StatusCode()
		if err != nil {
			if e, ok := err.(*fiber.Error); ok {
				statusCode = e.Code
			} else {
				statusCode = fiber.StatusInternalServerError
			}
		}
		if statusCode == 0 {
			statusCode = fiber.StatusOK
		}
		// map → 2xx / 4xx / 5xx
		statusLabel := httpStatusLabel(statusCode)

		// ===== Increment Metrics =====
		m.HttpRequestsTotal.
			WithLabelValues(c.Method(), route, statusLabel).
			Inc()
		m.HttpRequestDuration.
			WithLabelValues(c.Method(), route, statusLabel).
			Observe(time.Since(start).Seconds())

		return err
	}
}

func httpStatusLabel(code int) string {
	switch {
	case code >= 500:
		return "5xx"
	case code >= 400:
		return "4xx"
	case code >= 300:
		return "3xx"
	default:
		return "2xx"
	}
}
