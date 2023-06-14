package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vominhtrungpro/internal/middleware"
	"github.com/vominhtrungpro/internal/users"
)

// Map news routes
func MapNewsRoutes(userGroup *gin.RouterGroup, h users.Handler) {
	userGroup.Use(middleware.CORSMiddleware())
	userGroup.POST("/register", h.Register)
	userGroup.POST("/login", h.Login)
	userGroup.GET("/transaction", h.Transation)
}
