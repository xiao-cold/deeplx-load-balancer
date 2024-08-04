package cache

import (
	"time"

	"github.com/allegro/bigcache"
	"github.com/xiao-cold/deeplx-load-balancer/internal/metrics"
	"go.uber.org/zap"
)

type CacheManager struct {
	cache   *bigcache.BigCache
	logger  *zap.Logger
	metrics *metrics.Metrics
}

func NewCacheManager(logger *zap.Logger, metrics *metrics.Metrics) *CacheManager {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	return &CacheManager{
		cache:   cache,
		logger:  logger,
		metrics: metrics,
	}
}

func (cm *CacheManager) Get(key string) ([]byte, error) {
	value, err := cm.cache.Get(key)
	if err == nil {
		cm.metrics.CacheHits.Inc()
	} else {
		cm.metrics.CacheMisses.Inc()
	}
	return value, err
}

func (cm *CacheManager) Set(key string, value []byte) error {
	err := cm.cache.Set(key, value)
	if err != nil {
		cm.logger.Error("Failed to set cache", zap.Error(err))
	}
	return err
}
