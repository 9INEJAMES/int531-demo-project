// metrics.go
package api

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	HttpRequestsTotal   *prometheus.CounterVec
	HttpRequestDuration *prometheus.HistogramVec
}
func NewMetrics() *Metrics {
	return &Metrics{
		HttpRequestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "app",
				Subsystem: "http",
				Name:      "requests_total",
				Help:      "Total HTTP requests",
			},
			[]string{"method", "path", "status"},
		),
		HttpRequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: "app",
				Subsystem: "http",
				Name:      "request_duration_seconds",
				Help:      "HTTP request latency",
			},
			[]string{"method", "path", "status"},
		),
	}
}


