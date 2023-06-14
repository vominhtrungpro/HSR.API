package controller

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vominhtrungpro/internal/model/model"
	"github.com/vominhtrungpro/internal/users"
	"github.com/vominhtrungpro/internal/users/generator"
	"github.com/vominhtrungpro/internal/users/usermodel"
)

// Characters UseCase
type userController struct {
	userRepo users.Repository
}

func NewUserController(userRepo users.Repository) users.Controller {
	return &userController{userRepo: userRepo}
}

// Register user
func (u *userController) Register(ctx context.Context, request usermodel.CreateUserRequest) error {
	id, err := generator.ProductSNF.Generate()
	if err != nil {
		return err
	}
	checkusername, err := u.userRepo.CheckUsernameIfExist(ctx, request.Username)
	if err != nil {
		return err
	}
	if checkusername {
		return errUsernameExist
	}
	checkemail, err := u.userRepo.CheckEmailIfExist(ctx, request.Email)
	if err != nil {
		return err
	}
	if checkemail {
		return errEmailExist
	}
	var user = model.User{
		ID:       int32(id),
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
	}
	result := u.userRepo.Create(ctx, user)
	if result != nil {
		return result
	}
	return nil
}

// Login
func (u *userController) Login(ctx context.Context, input usermodel.LoginInput) (usermodel.LoginOutput, error) {
	var token usermodel.LoginOutput
	user, err := u.userRepo.GetUserByUsername(ctx, input.Username)
	if err != nil {
		return usermodel.LoginOutput{}, err
	}

	if input.Password != user.Password {
		return usermodel.LoginOutput{}, errIncorrectPassword
	}

	accesstoken, err := CreateAccessToken(user)
	if err != nil {
		return usermodel.LoginOutput{}, err
	}

	refreshtoken, err := CreateRefreshToken()
	if err != nil {
		return usermodel.LoginOutput{}, err
	}
	expdate := time.Now()
	user.Refreshtoken = refreshtoken
	user.Refreshtokenexpiredate = &expdate
	err = u.userRepo.Update(ctx, user)
	if err != nil {
		return usermodel.LoginOutput{}, err
	}

	token.AccessToken = accesstoken
	token.RefreshToken = refreshtoken
	return token, nil
}

// Create access token
func CreateAccessToken(user model.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"id":       user.ID,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 240).Unix(),
		"role":     "None",
		"username": user.Username,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Create refreshtoken
func CreateRefreshToken() (string, error) {
	b := make([]byte, 64)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	token := hex.EncodeToString(b)
	return token, nil
}
