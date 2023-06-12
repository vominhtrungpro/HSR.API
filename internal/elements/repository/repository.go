package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/vominhtrungpro/internal/elements"
	"github.com/vominhtrungpro/internal/elements/generator"
	"github.com/vominhtrungpro/internal/model/model"
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

// Create element
func (r *eleRepo) Create(ctx context.Context, element model.Element) error {
	id, err := generator.ProductSNF.Generate()
	if err != nil {
		return err
	}

	element.ID = int32(id)
	if err = r.db.Model(&model.Element{}).Create(element).Error; err != nil {
		return err
	}

	return nil
}

func (r *eleRepo) CreatePath(ctx context.Context, path model.Path) error {
	id, err := generator.ProductSNF.Generate()
	if err != nil {
		return err
	}
	path.ID = int32(id)
	if err = r.db.Model(&model.Path{}).Create(path).Error; err != nil {
		return err
	}
	return nil
}

// Get element by en name
func (r *eleRepo) GetElementByEnname(ctx context.Context, name string) (model.Element, error) {
	var element model.Element
	err := r.db.Where("enname = ?", name).First(&element).Error
	if err != nil {
		return element, err
	}
	return element, nil
}

// Get path by en name
func (r *eleRepo) GetPathByEnname(ctx context.Context, name string) (model.Path, error) {
	var path model.Path
	err := r.db.Where("enname = ?", name).First(&path).Error
	if err != nil {
		return path, err
	}
	return path, nil
}

// Update image for element
func (r *eleRepo) UpdateElementImage(ctx context.Context, element model.Element, image []byte) error {
	element.Picture = image
	err := r.db.Save(&element).Error
	if err != nil {
		return err
	}
	assetsurl := os.Getenv("ASSETs_URL")
	fileurl := fmt.Sprintf(assetsurl + element.Enname + ".png")
	err = os.WriteFile(fileurl, image, 0644)
	if err != nil {
		panic(fileurl)
	}
	return nil
}

// Update image for path
func (r *eleRepo) UpdatePathImage(ctx context.Context, path model.Path, image []byte) error {
	path.Picture = image
	err := r.db.Save(&path).Error
	if err != nil {
		return err
	}
	assetsurl := os.Getenv("ASSETs_URL")
	fileurl := fmt.Sprintf(assetsurl + path.Enname + ".png")
	err = os.WriteFile(fileurl, image, 0644)
	if err != nil {
		panic(fileurl)
	}
	return nil
}
