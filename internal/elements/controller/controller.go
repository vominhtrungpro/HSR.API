package controller

import (
	"context"

	"github.com/vominhtrungpro/internal/elements"
	"github.com/vominhtrungpro/internal/elements/elementmodel"
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

// Get filter
func (u *eleController) GetFilter(ctx context.Context) ([]elementmodel.FilterResponse, error) {
	var filter []elementmodel.FilterResponse
	var fourstar = elementmodel.FilterResponse{
		Name:   "4 Sao",
		Enname: "4 Star",
		Type:   "Star",
	}
	filter = append(filter, fourstar)
	var fivestar = elementmodel.FilterResponse{
		Name:   "5 Sao",
		Enname: "5 Star",
		Type:   "Star",
	}
	filter = append(filter, fivestar)

	element, err := u.eleRepo.GetElement(ctx)
	if err != nil {
		return []elementmodel.FilterResponse{}, err
	}
	for _, item := range element {
		var element = elementmodel.FilterResponse{
			Name:   item.Name,
			Enname: item.Enname,
			Type:   "Element",
		}
		filter = append(filter, element)
	}

	path, err := u.eleRepo.GetPath(ctx)
	if err != nil {
		return []elementmodel.FilterResponse{}, err
	}
	for _, item := range path {
		var element = elementmodel.FilterResponse{
			Name:   item.Name,
			Enname: item.Enname,
			Type:   "Path",
		}
		filter = append(filter, element)
	}
	return filter, nil
}
