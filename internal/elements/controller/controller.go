package controller

import (
	"context"

	"github.com/vominhtrungpro/internal/elements"
	"github.com/vominhtrungpro/internal/model/model"
)

// Element Controller
type eleController struct {
	eleRepo elements.Repository
}

func NewElementController(eleRepo elements.Repository) elements.Controller {
	return &eleController{eleRepo: eleRepo}
}

// Create element
func (u *eleController) Create(ctx context.Context, element model.Element) error {
	err := u.eleRepo.Create(ctx, element)
	if err != nil {
		return err
	}
	return nil
}

// Create path
func (u *eleController) CreatePath(ctx context.Context, path model.Path) error {
	err := u.eleRepo.CreatePath(ctx, path)
	if err != nil {
		return err
	}
	return nil
}

// Update element image
func (u *eleController) UpdateElementImage(ctx context.Context, name string, image []byte) error {
	element, err := u.eleRepo.GetElementByEnname(ctx, name)
	if err != nil {
		return err
	}
	updateErr := u.eleRepo.UpdateElementImage(ctx, element, image)
	if updateErr != nil {
		return updateErr
	}
	return nil
}

// Update path image
func (u *eleController) UpdatePathImage(ctx context.Context, name string, image []byte) error {
	path, err := u.eleRepo.GetPathByEnname(ctx, name)
	if err != nil {
		return err
	}
	updateErr := u.eleRepo.UpdatePathImage(ctx, path, image)
	if updateErr != nil {
		return updateErr
	}
	return nil
}
