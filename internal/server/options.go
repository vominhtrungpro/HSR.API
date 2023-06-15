package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vominhtrungpro/pkg/cache/redis"
)

// Option -.
type Option func(*Server)

func FiberEngine(gin *gin.Engine) Option {
	return func(s *Server) {
		s.gin = gin
	}
}

func Redis(rdb redis.Client) Option {
	return func(s *Server) {
		s.redis = rdb
	}
}
