package api

import (
	"io"
	"net/http"

	"deeplx-load-balancer/internal/cache"
	"deeplx-load-balancer/internal/loadbalancer"
	"deeplx-load-balancer/internal/metrics"
	"deeplx-load-balancer/pkg/utils"

	"go.uber.org/zap"
)

type APIGateway struct {
	loadBalancer *loadbalancer.LoadBalancer
	cacheManager *cache.CacheManager
	logger       *zap.Logger
	metrics      *metrics.Metrics
}

func NewAPIGateway(logger *zap.Logger, metrics *metrics.Metrics, lb *loadbalancer.LoadBalancer, cm *cache.CacheManager) *APIGateway {
	return &APIGateway{
		loadBalancer: lb,
		cacheManager: cm,
		logger:       logger,
		metrics:      metrics,
	}
}

func (ag *APIGateway) HandleTranslate(w http.ResponseWriter, r *http.Request) {
	apiKey := r.URL.Path[1:] // Extract API key from URL path
	if !ag.validateAPIKey(apiKey) {
		http.Error(w, "Invalid API key", http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	cacheKey := ag.generateCacheKey(apiKey, body)
	if cachedResponse, err := ag.cacheManager.Get(cacheKey); err == nil {
		w.Write(cachedResponse)
		return
	}

	server, err := ag.loadBalancer.NextServer()
	if err != nil {
		http.Error(w, "No servers available", http.StatusServiceUnavailable)
		return
	}

	resp, err := utils.ForwardRequest(server, body)
	if err != nil {
		ag.logger.Error("Failed to forward request", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	ag.cacheManager.Set(cacheKey, resp)
	w.Write(resp)
}

func (ag *APIGateway) validateAPIKey(apiKey string) bool {
	// Implement API key validation logic
	return true // Placeholder
}

func (ag *APIGateway) generateCacheKey(apiKey string, body []byte) string {
	// Implement cache key generation logic
	return string(body) // Placeholder
}
