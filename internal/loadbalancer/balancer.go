package loadbalancer

import (
	"errors"
	"sync"
	"sync/atomic"

	"deeplx-load-balancer/internal/metrics"
	"deeplx-load-balancer/internal/models"

	"go.uber.org/zap"
)

type LoadBalancer struct {
	servers      []*models.Server
	mu           sync.RWMutex
	currentIndex int32
	logger       *zap.Logger
	metrics      *metrics.Metrics
}

func NewLoadBalancer(logger *zap.Logger, metrics *metrics.Metrics) *LoadBalancer {
	return &LoadBalancer{
		logger:  logger,
		metrics: metrics,
	}
}

func (lb *LoadBalancer) NextServer() (*models.Server, error) {
	lb.mu.RLock()
	defer lb.mu.RUnlock()

	if len(lb.servers) == 0 {
		return nil, errors.New("no servers available")
	}

	// Implement weighted round-robin or other advanced algorithm here
	server := lb.servers[atomic.AddInt32(&lb.currentIndex, 1)%int32(len(lb.servers))]

	if !server.IsHealthy() {
		return lb.NextServer()
	}

	return server, nil
}

func (lb *LoadBalancer) UpdateServerList(servers []*models.Server) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	lb.servers = servers
}
