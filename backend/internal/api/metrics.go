package api

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	HttpRequestsTotal    *prometheus.CounterVec
	HttpRequestDuration *prometheus.HistogramVec
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		HttpRequestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "app",
				Subsystem: "http",
				Name:      "requests_total",
				Help:      "Total HTTP requests",
			},
			[]string{"method", "route", "status"},
		),
		HttpRequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: "app",
				Subsystem: "http",
				Name:      "request_duration_seconds",
				Help:      "HTTP request latency",
				Buckets:   prometheus.DefBuckets,
			},
			[]string{"method", "route"},
		),
	}

	reg.MustRegister(
		m.HttpRequestsTotal,
		m.HttpRequestDuration,
	)

	return m
}
