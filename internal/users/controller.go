package users

import (
	"context"

	"github.com/vominhtrungpro/internal/users/usermodel"
)

type Controller interface {
	Register(ctx context.Context, request usermodel.CreateUserRequest) error
	Login(ctx context.Context, input usermodel.LoginInput) (usermodel.LoginOutput, error)
}
