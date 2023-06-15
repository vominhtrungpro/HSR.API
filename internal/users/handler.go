package users

import "github.com/gin-gonic/gin"

type Handler interface {
	Register(context *gin.Context)
	Login(context *gin.Context)
	Transation(context *gin.Context)
	GetCache(context *gin.Context)
}
