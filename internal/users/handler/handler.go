package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/mail"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vominhtrungpro/internal/users"
	"github.com/vominhtrungpro/internal/users/usermodel"
	"github.com/vominhtrungpro/pkg/cache/redis"
)

// Users handler
type userHandler struct {
	userController users.Controller
	rdb            redis.Client
}

// User cache
type CacheValue struct {
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
	LastLogin    string `json:"lastlogin"`
}

// Characters handlers constructor
func NewUserHandlers(userController users.Controller, rdb redis.Client) users.Handler {
	return &userHandler{
		userController: userController,
		rdb:            rdb,
	}
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
func (h userHandler) Login(ctx *gin.Context) {
	var request usermodel.LoginInput
	err := json.NewDecoder(ctx.Request.Body).Decode(&request)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := validatelogin(request); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result, err := h.userController.Login(ctx, request)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	rdbcontext := context.Background()
	cachevalue := CacheValue{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		LastLogin:    time.Now().Format("2006-01-02 15:04:05"),
	}
	b, err := json.Marshal(cachevalue)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.rdb.Set(rdbcontext, request.Username, string(b), 0)
	if err != nil {
		panic(err)
	}
	ctx.IndentedJSON(http.StatusOK, result)
}

func (h userHandler) GetCache(context *gin.Context) {
	charname := context.Param("name")
	var result CacheValue
	cache, err := h.rdb.Get(context, charname)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := json.Unmarshal([]byte(cache), &result); err != nil {
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

// Login
func (h userHandler) Transation(context *gin.Context) {
	err := h.userController.TestTransaction(context)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	context.IndentedJSON(http.StatusOK, "Success")
}
