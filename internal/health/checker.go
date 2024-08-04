package health

import (
	"sync"
	"time"

	"deeplx-load-balancer/internal/metrics"
	"deeplx-load-balancer/internal/models"

	"go.uber.org/zap"
)

type HealthChecker struct {
	servers []*models.Server
	logger  *zap.Logger
	metrics *metrics.Metrics
}

func NewHealthChecker(logger *zap.Logger, metrics *metrics.Metrics) *HealthChecker {
	return &HealthChecker{
		logger:  logger,
		metrics: metrics,
	}
}

func (hc *HealthChecker) Start() {
	ticker := time.NewTicker(30 * time.Second)
	for range ticker.C {
		hc.checkHealth()
	}
}

func (hc *HealthChecker) checkHealth() {
	var wg sync.WaitGroup
	for _, server := range hc.servers {
		wg.Add(1)
		go func(s *models.Server) {
			defer wg.Done()
			healthy := hc.isServerHealthy(s.URL)
			s.SetHealth(healthy)
			if healthy {
				hc.metrics.HealthyServers.Inc()
			} else {
				hc.metrics.UnhealthyServers.Inc()
			}
		}(server)
	}
	wg.Wait()
}

func (hc *HealthChecker) isServerHealthy(url string) bool {
	// Perform a health check on the server
	return true
}
