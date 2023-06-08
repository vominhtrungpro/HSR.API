package server

import (
	characterController "github.com/vominhtrungpro/internal/characters/controller"
	characterHandler "github.com/vominhtrungpro/internal/characters/handler"
	characterRepository "github.com/vominhtrungpro/internal/characters/repository"
)

// Map Server Handlers
func (s *Server) MapHandlers() error {
	// Init repositories
	cRepo := characterRepository.NewCharsRepository(s.db)
	// Init controller
	cCtrl := characterController.NewCharacterController(cRepo)

	// Init handlers
	cHdler := characterHandler.NewCharsHandlers(cCtrl)
	v1 := s.gin.Group("/api/v1")
	newsGroup := v1.Group("/characters")
	characterHandler.MapNewsRoutes(newsGroup, cHdler)
	return nil
}
