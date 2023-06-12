package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vominhtrungpro/internal/elements"
)

// Map news routes
func MapNewsRoutes(elementGroup *gin.RouterGroup, h elements.Handlers) {
	elementGroup.POST("/create", h.Create)
	elementGroup.POST("/image/:name", h.UpdateElementImage)
	elementGroup.POST("/createpath", h.CreatePath)
	elementGroup.POST("/imagepath/:name", h.UpdatePathImage)
}
