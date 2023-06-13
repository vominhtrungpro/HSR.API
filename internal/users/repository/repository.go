package repository

import (
	"context"

	"github.com/vominhtrungpro/internal/model/model"
	"github.com/vominhtrungpro/internal/users"
	"gorm.io/gorm"
)

// Users repository
type userRepository struct {
	db *gorm.DB
}

// Users repository constructor
func NewUserRepository(db *gorm.DB) users.Repository {
	return &userRepository{db: db}
}

// Create character
func (r *userRepository) Create(ctx context.Context, user model.User) error {
	if err := r.db.Model(&model.User{}).Create(user).Error; err != nil {
		return err
	}
	return nil
}
