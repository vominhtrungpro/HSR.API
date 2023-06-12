package server

import (
	characterController "github.com/vominhtrungpro/internal/characters/controller"
	characterHandler "github.com/vominhtrungpro/internal/characters/handler"
	characterRepository "github.com/vominhtrungpro/internal/characters/repository"
	elementController "github.com/vominhtrungpro/internal/elements/controller"
	elementHandler "github.com/vominhtrungpro/internal/elements/handler"
	elementRepository "github.com/vominhtrungpro/internal/elements/repository"
)

// Map Server Handlers
func (s *Server) MapHandlers() error {
	// Init repositories
	cRepo := characterRepository.NewCharsRepository(s.db)
	eRepo := elementRepository.NewElementRepository(s.db)
	// Init controller
	cCtrl := characterController.NewCharacterController(cRepo)
	eCtrl := elementController.NewElementController(eRepo)

	// Init handlers
	cHdler := characterHandler.NewCharsHandlers(cCtrl)

	eHdler := elementHandler.NewElementHandlers(eCtrl)
	v1 := s.gin.Group("/api/v1")
	charactersGroup := v1.Group("/characters")
	elementsGroup := v1.Group("/elements")
	characterHandler.MapNewsRoutes(charactersGroup, cHdler)
	elementHandler.MapNewsRoutes(elementsGroup, eHdler)

	return nil
}
