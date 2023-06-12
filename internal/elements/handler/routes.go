package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vominhtrungpro/internal/elements"
	"github.com/vominhtrungpro/internal/middleware"
)

// Map news routes
func MapNewsRoutes(elementGroup *gin.RouterGroup, h elements.Handlers) {
	elementGroup.Use(middleware.CORSMiddleware())
	elementGroup.POST("/create", h.Create)
	elementGroup.POST("/image/:name", h.UpdateElementImage)
	elementGroup.POST("/createpath", h.CreatePath)
	elementGroup.POST("/imagepath/:name", h.UpdatePathImage)
	elementGroup.GET("/filter", h.GetFilter)
}
