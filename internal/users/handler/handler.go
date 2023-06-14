package handler

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"strings"

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
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := validateregister(request); err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result := h.userController.Register(context, request)
	if result != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	context.IndentedJSON(http.StatusOK, "Success")
}

// Login
func (h userHandler) Login(context *gin.Context) {
	var request usermodel.LoginInput
	err := json.NewDecoder(context.Request.Body).Decode(&request)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := validatelogin(request); err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result, err := h.userController.Login(context, request)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	context.IndentedJSON(http.StatusOK, result)
}

// Validate login
func validatelogin(input usermodel.LoginInput) error {
	if strings.TrimSpace(input.Username) == "" {
		return errInvalidUsername
	}

	if strings.TrimSpace(input.Password) == "" {
		return errInvalidPassword
	}
	return nil
}

// Validate register
func validateregister(input usermodel.CreateUserRequest) error {
	if strings.TrimSpace(input.Username) == "" {
		return errInvalidUsername
	}

	if strings.TrimSpace(input.Password) == "" {
		return errInvalidPassword
	}
	valid := valid(input.Email)
	if !valid {
		return errInvalidEmail
	}
	return nil
}

// Validate email
func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
