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

// Create user
func (r *userRepository) Create(ctx context.Context, user model.User) error {
	if err := r.db.Model(&model.User{}).Create(user).Error; err != nil {
		return err
	}
	return nil
}

// Get user by username
func (r *userRepository) GetUserByUsername(ctx context.Context, name string) (model.User, error) {
	var user model.User
	if err := r.db.Where("username = ?", name).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Update user
func (r *userRepository) Update(ctx context.Context, user model.User) error {
	return r.db.Save(&user).Error
}
