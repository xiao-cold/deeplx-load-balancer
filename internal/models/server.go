package models

import "sync/atomic"

type Server struct {
	URL              string
	health           int32
	PerformanceScore float64
}

func NewServer(url string) *Server {
	return &Server{
		URL:    url,
		health: 1, // Initially set as healthy
	}
}

func (s *Server) IsHealthy() bool {
	return atomic.LoadInt32(&s.health) == 1
}

func (s *Server) SetHealth(healthy bool) {
	if healthy {
		atomic.StoreInt32(&s.health, 1)
	} else {
		atomic.StoreInt32(&s.health, 0)
	}
}
