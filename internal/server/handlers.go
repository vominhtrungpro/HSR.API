package server

import (
	characterController "github.com/vominhtrungpro/internal/characters/controller"
	characterHandler "github.com/vominhtrungpro/internal/characters/handler"
	characterRepository "github.com/vominhtrungpro/internal/characters/repository"
	elementController "github.com/vominhtrungpro/internal/elements/controller"
	elementHandler "github.com/vominhtrungpro/internal/elements/handler"
	elementRepository "github.com/vominhtrungpro/internal/elements/repository"
	userController "github.com/vominhtrungpro/internal/users/controller"
	userHandler "github.com/vominhtrungpro/internal/users/handler"
	userRepository "github.com/vominhtrungpro/internal/users/repository"
)

// Map Server Handlers
func (s *Server) MapHandlers() error {
	// Init repositories
	cRepo := characterRepository.NewCharsRepository(s.db)
	eRepo := elementRepository.NewElementRepository(s.db)
	uRepo := userRepository.NewUserRepository(s.db)

	// Init controller
	cCtrl := characterController.NewCharacterController(cRepo)
	eCtrl := elementController.NewElementController(eRepo)
	uCtrl := userController.NewUserController(uRepo)

	// Init handlers
	cHdler := characterHandler.NewCharsHandlers(cCtrl)
	eHdler := elementHandler.NewElementHandlers(eCtrl)
	uHdler := userHandler.NewUserHandlers(uCtrl, s.redis)

	v1 := s.gin.Group("/api/v1")
	charactersGroup := v1.Group("/characters")
	elementsGroup := v1.Group("/elements")
	usersGroup := v1.Group("/users")
	characterHandler.MapNewsRoutes(charactersGroup, cHdler)
	elementHandler.MapNewsRoutes(elementsGroup, eHdler)
	userHandler.MapNewsRoutes(usersGroup, uHdler)
	return nil
}
