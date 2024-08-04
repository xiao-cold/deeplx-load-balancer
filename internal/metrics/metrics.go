package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	RequestsTotal    prometheus.Counter
	RequestDuration  prometheus.Histogram
	CacheHits        prometheus.Counter
	CacheMisses      prometheus.Counter
	HealthyServers   prometheus.Gauge
	UnhealthyServers prometheus.Gauge
}

func NewMetrics() *Metrics {
	return &Metrics{
		RequestsTotal: promauto.NewCounter(prometheus.CounterOpts{
			Name: "requests_total",
			Help: "The total number of processed requests",
		}),
		RequestDuration: promauto.NewHistogram(prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "The duration of requests in seconds",
			Buckets: prometheus.DefBuckets,
		}),
		CacheHits: promauto.NewCounter(prometheus.CounterOpts{
			Name: "cache_hits_total",
			Help: "The total number of cache hits",
		}),
		CacheMisses: promauto.NewCounter(prometheus.CounterOpts{
			Name: "cache_misses_total",
			Help: "The total number of cache misses",
		}),
		HealthyServers: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "healthy_servers",
			Help: "The number of healthy servers",
		}),
		UnhealthyServers: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "unhealthy_servers",
			Help: "The number of unhealthy servers",
		}),
	}
}
