// metrics.go
package api

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	HttpRequestsTotal   *prometheus.CounterVec
	HttpRequestDuration *prometheus.HistogramVec
}

func NewMetrics(reg *prometheus.Registry) *Metrics {
	m := &Metrics{
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

	// register metric แบบไม่ panic ถ้า register ซ้ำ
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

	return m
}

