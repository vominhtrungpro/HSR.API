package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vominhtrungpro/internal/users"
	"github.com/vominhtrungpro/internal/users/usermodel"
)

// Users handler
type userHandler struct {
	userController users.Controller
}

// Characters handlers constructor
func NewUserHandlers(userController users.Controller) users.Handler {
	return &userHandler{userController: userController}
}

// Register user handler
func (h userHandler) Register(context *gin.Context) {
	var request usermodel.CreateUserRequest
	err := json.NewDecoder(context.Request.Body).Decode(&request)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	result := h.userController.Register(context, request)
	if result != nil {
		http.Error(context.Writer, result.Error(), http.StatusBadRequest)
		return
	}
	context.IndentedJSON(http.StatusOK, "Success")
}
