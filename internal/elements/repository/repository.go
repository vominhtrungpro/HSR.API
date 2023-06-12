package repository

import (
	"github.com/vominhtrungpro/internal/elements"
	"gorm.io/gorm"
)

// Characters Repository
type eleRepo struct {
	db *gorm.DB
}

// Characters repository constructor
func NewElementRepository(db *gorm.DB) elements.Repository {
	return &eleRepo{db: db}
}
