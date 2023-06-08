package repository

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/vominhtrungpro/internal/characters"
	"github.com/vominhtrungpro/internal/characters/charactermodel"
	"github.com/vominhtrungpro/internal/model/model"
	"gorm.io/gorm"
)

// Characters Repository
type charsRepo struct {
	db *gorm.DB
}

// Characters repository constructor
func NewCharsRepository(db *gorm.DB) characters.Repository {
	return &charsRepo{db: db}
}

// Create character
func (r *charsRepo) Create(ctx context.Context, character model.Character) error {
	err := r.db.Model(&model.Character{}).Create(character).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *charsRepo) GetCharById(ctx context.Context, id string) (model.Character, error) {
	var chardb model.Character
	err := r.db.Where("id = ?", id).First(&chardb).Error
	if err != nil {
		return chardb, err
	}
	return chardb, nil
}

func (r *charsRepo) UpdateCharImage(ctx context.Context, char model.Character, image []byte) error {
	char.Picture = image
	err := r.db.Save(&char).Error
	if err != nil {
		return err
	}
	assetsurl := os.Getenv("ASSETs_URL")
	fileurl := fmt.Sprintf(assetsurl + char.Name + ".png")
	err = ioutil.WriteFile(fileurl, image, 0644)
	if err != nil {
		panic(fileurl)
	}
	return nil
}

func (r *charsRepo) GetImageById(ctx context.Context, id string) ([]byte, string, error) {
	var chardb model.Character
	err := r.db.Where("id = ?", id).First(&chardb).Error
	if err != nil {
		return nil, "", err
	}
	return chardb.Picture, chardb.Name, nil
}

func (r *charsRepo) GetAll(ctx context.Context) ([]charactermodel.SearchResult, error) {
	var characters []charactermodel.SearchResult
	var chardb []model.Character
	if err := r.db.Order("id").Find(&chardb).Error; err != nil {
		return characters, err
	}
	assetsurl := os.Getenv("ASSETs_URL")
	for _, element := range chardb {
		var character charactermodel.SearchResult
		character.ID = element.ID
		character.Name = element.Name
		character.Rarity = element.Rarity
		character.Element = element.Element
		character.Path = element.Path
		fileurl := fmt.Sprintf(assetsurl + element.Name + ".png")
		character.Url = fileurl
		characters = append(characters, character)
	}
	return characters, nil
}
