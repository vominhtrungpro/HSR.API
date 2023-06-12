package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vominhtrungpro/internal/characters"
	"github.com/vominhtrungpro/internal/middleware"
)

// Map news routes
func MapNewsRoutes(charGroup *gin.RouterGroup, h characters.Handlers) {
	charGroup.Use(middleware.CORSMiddleware())
	charGroup.GET("/all", h.GetAll)
	charGroup.POST("/create", h.Create)
	charGroup.PUT("/update", h.Update)
	charGroup.POST("/image/:name", h.UpdateCharacterImage)
	charGroup.GET("/image/:id", h.GetImageById)
	charGroup.GET("", h.Test)
}
