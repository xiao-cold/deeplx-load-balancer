package main

import (
	"net/http"

	"github.com/xiao-cold/deeplx-load-balancer/internal/api"
	"github.com/xiao-cold/deeplx-load-balancer/internal/cache"
	"github.com/xiao-cold/deeplx-load-balancer/internal/health"
	"github.com/xiao-cold/deeplx-load-balancer/internal/loadbalancer"
	"github.com/xiao-cold/deeplx-load-balancer/internal/metrics"
	"github.com/xiao-cold/deeplx-load-balancer/internal/performance"

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
	http.Handle("/translate", apiGateway.HandleTranslate)

	logger.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
