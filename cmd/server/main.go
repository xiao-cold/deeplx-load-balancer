package main

import (
	"net/http"

	"deeplx-load-balancer/internal/api"
	"deeplx-load-balancer/internal/cache"
	"deeplx-load-balancer/internal/health"
	"deeplx-load-balancer/internal/loadbalancer"
	"deeplx-load-balancer/internal/metrics"
	"deeplx-load-balancer/internal/performance"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	metricsCollector := metrics.NewMetrics()
	cacheManager := cache.NewCacheManager(logger, metricsCollector)
	healthChecker := health.NewHealthChecker(logger, metricsCollector)
	performanceTracker := performance.NewPerformanceTracker(logger, metricsCollector)
	loadBalancer := loadbalancer.NewLoadBalancer(logger, metricsCollector)
	apiGateway := api.NewAPIGateway(logger, metricsCollector, loadBalancer, cacheManager)

	// Start background processes
	go healthChecker.Start()
	go performanceTracker.Start()

	// Set up HTTP server
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/translate", http.HandlerFunc(apiGateway.HandleTranslate))

	logger.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
