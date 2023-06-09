package controller

import (
	"context"

	"github.com/vominhtrungpro/internal/characters"
	"github.com/vominhtrungpro/internal/characters/charactermodel"
	"github.com/vominhtrungpro/internal/model/model"
)

// Characters UseCase
type charController struct {
	charsRepo characters.Repository
}

func NewCharacterController(charsRepo characters.Repository) characters.Controller {
	return &charController{charsRepo: charsRepo}
}

// Create character
func (u *charController) Create(ctx context.Context, news model.Character) error {
	err := u.charsRepo.Create(ctx, news)
	if err != nil {
		return err
	}
	return nil
}

// Update character image
func (u *charController) UpdateCharacterImage(ctx context.Context, name string, image []byte) error {
	character, err := u.charsRepo.GetCharByName(ctx, name)
	if err != nil {
		return err
	}
	updateErr := u.charsRepo.UpdateCharImage(ctx, character, image)
	if err != nil {
		return err
	}
	return updateErr
}

func (u *charController) GetImageById(ctx context.Context, id string) ([]byte, string, error) {
	image, name, err := u.charsRepo.GetImageById(ctx, id)
	if err != nil {
		return nil, "", err
	}
	return image, name, nil
}

func (u *charController) GetAll(ctx context.Context) ([]charactermodel.SearchResult, error) {
	characters, err := u.charsRepo.GetAll(ctx)
	if err != nil {
		return characters, err
	}
	return characters, nil
}

func (u *charController) Update(ctx context.Context, request charactermodel.UpdateRequest) error {
	character, err := u.charsRepo.GetCharById(ctx, request.Id)
	if err != nil {
		return err
	}
	checkexist := u.charsRepo.CheckNameExist(ctx, character)
	if checkexist != nil {
		return checkexist
	}
	updateResult := u.charsRepo.Update(ctx, character)
	if updateResult != nil {
		return updateResult
	}
	return nil
}
