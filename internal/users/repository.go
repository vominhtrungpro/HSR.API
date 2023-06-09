package users

import (
	"context"

	"github.com/vominhtrungpro/internal/model/model"
)

type Repository interface {
	Create(ctx context.Context, user model.User) error
	GetUserByUsername(ctx context.Context, name string) (model.User, error)
	Update(ctx context.Context, user model.User) error
	CheckUsernameIfExist(ctx context.Context, username string) (bool, error)
	CheckEmailIfExist(ctx context.Context, email string) (bool, error)
	TestTransaction(ctx context.Context) error
}
