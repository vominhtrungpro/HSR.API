package characters

import (
	"github.com/gin-gonic/gin"
)

// News HTTP Handlers interface
type Handlers interface {
	Create(c *gin.Context)
	Test(c *gin.Context)
	UpdateCharacterImage(context *gin.Context)
	GetImageById(context *gin.Context)
	GetAll(context *gin.Context)
	Update(context *gin.Context)
}
