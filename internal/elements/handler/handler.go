package handler

import "github.com/vominhtrungpro/internal/elements"

type eleHandler struct {
	eleController elements.Controller
}

// Characters handlers constructor
func NewElementHandlers(eleController elements.Controller) elements.Handlers {
	return &eleHandler{eleController: eleController}
}
