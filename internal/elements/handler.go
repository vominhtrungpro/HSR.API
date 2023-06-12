package elements

import "github.com/gin-gonic/gin"

// News HTTP Handlers interface
type Handlers interface {
	Create(context *gin.Context)
	CreatePath(context *gin.Context)
	UpdateElementImage(context *gin.Context)
	UpdatePathImage(context *gin.Context)
	GetFilter(context *gin.Context)
}
