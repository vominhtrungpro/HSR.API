package characters

import (
	"context"

	"github.com/vominhtrungpro/internal/characters/charactermodel"
	"github.com/vominhtrungpro/internal/model/model"
)

// News use case
type Controller interface {
	Create(ctx context.Context, char model.Character) error
	UpdateCharacterImage(ctx context.Context, id string, image []byte) error
	GetImageById(ctx context.Context, id string) ([]byte, string, error)
	GetAll(ctx context.Context) ([]charactermodel.SearchResult, error)
}
