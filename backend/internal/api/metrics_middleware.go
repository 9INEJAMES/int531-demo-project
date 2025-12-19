package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// MetricsMiddleware records HTTP metrics
func MetricsMiddleware(m *Metrics) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()

		// skip internal endpoints
		if c.Path() == "/metrics" || c.Path() == "/health" {
			return err
		}

		// normalize route
		route := normalizePath(c)

		// normalize method
		method := c.Method()

		// determine status code
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
		statusLabel := httpStatusLabel(statusCode)

		// ===== record metrics =====
		m.HttpRequestsTotal.
			WithLabelValues(method, route, statusLabel).
			Inc()
		m.HttpRequestDuration.
			WithLabelValues(method, route, statusLabel).
			Observe(time.Since(start).Seconds())

		return err
	}
}

// httpStatusLabel maps status code to a label
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

func normalizePath(c *fiber.Ctx) string {
	route := "unknown"
	if r := c.Route(); r != nil && r.Path != "" {
		route = r.Path
	}
	return route
}
