package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vominhtrungpro/internal/characters"
)

// Map news routes
func MapNewsRoutes(charGroup *gin.RouterGroup, h characters.Handlers) {
	// newsGroup.Use(mw.AuthJWTMiddleware())
	charGroup.GET("/all", h.GetAll)
	charGroup.POST("/create", h.Create)
	charGroup.PUT("/update", h.Update)
	charGroup.POST("/image/:name", h.UpdateCharacterImage)
	charGroup.GET("/image/:id", h.GetImageById)
	charGroup.GET("", h.Test)
}
