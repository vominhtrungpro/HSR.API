package repository

import (
	"context"

	"github.com/vominhtrungpro/internal/model/model"
	"github.com/vominhtrungpro/internal/users"
	"github.com/vominhtrungpro/internal/users/generator"
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
		if err == gorm.ErrRecordNotFound {
			return user, errUserNotFound
		}
		return user, err
	}
	return user, nil
}

// Update user
func (r *userRepository) Update(ctx context.Context, user model.User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) CheckUsernameIfExist(ctx context.Context, username string) (bool, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (r *userRepository) CheckEmailIfExist(ctx context.Context, email string) (bool, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (r *userRepository) TestTransaction(ctx context.Context) error {
	tx := r.db.Begin()
	var user model.User
	id, err := generator.ProductSNF.Generate()
	if err != nil {
		return err
	}
	user.ID = int32(id)
	user.Username = "Transaction1"
	user.Password = "123456789"
	user.Email = "Transaction1@gmail.com"
	if err := tx.Model(&model.User{}).Create(user).Error; err != nil {
		return err
	}
	var userdb model.User
	if err := r.db.Where("username = ?", "vominhtrung1").First(&userdb).Error; err != nil {
		tx.Rollback()
		tx.Commit()
		return err
	}
	tx.Commit()
	return nil
}
