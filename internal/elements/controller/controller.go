package controller

import "github.com/vominhtrungpro/internal/elements"

// Characters UseCase
type eleController struct {
	eleRepo elements.Repository
}

func NewElementController(eleRepo elements.Repository) elements.Controller {
	return &eleController{eleRepo: eleRepo}
}
