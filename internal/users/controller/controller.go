package controller

import (
	"context"

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
