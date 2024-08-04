package performance

import (
	"sync"
	"time"

	"github.com/xiao-cold/deeplx-load-balancer/internal/metrics"
	"github.com/xiao-cold/deeplx-load-balancer/internal/models"
	"go.uber.org/zap"
)

type PerformanceTracker struct {
	servers []*models.Server
	logger  *zap.Logger
	metrics *metrics.Metrics
	mu      sync.RWMutex
}

func NewPerformanceTracker(logger *zap.Logger, metrics *metrics.Metrics) *PerformanceTracker {
	return &PerformanceTracker{
		logger:  logger,
		metrics: metrics,
	}
}

func (pt *PerformanceTracker) Start() {
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		pt.updatePerformanceScores()
	}
}

func (pt *PerformanceTracker) updatePerformanceScores() {
	pt.mu.RLock()
	defer pt.mu.RUnlock()

	for _, server := range pt.servers {
		score := pt.calculatePerformanceScore(server)
		server.PerformanceScore = score
	}
}

func (pt *PerformanceTracker) calculatePerformanceScore(server *models.Server) float64 {
	// Implement performance score calculation logic
	// This could include factors like response time, success rate, etc.
	return 1.0 // Placeholder
}

func (pt *PerformanceTracker) UpdateServerList(servers []*models.Server) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.servers = servers
}
