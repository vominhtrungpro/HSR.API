package characters

import (
	"context"

	"github.com/vominhtrungpro/internal/characters/charactermodel"
	"github.com/vominhtrungpro/internal/model/model"
)

type Repository interface {
	Create(ctx context.Context, news model.Character) error
	GetCharById(ctx context.Context, id string) (model.Character, error)
	UpdateCharImage(ctx context.Context, char model.Character, image []byte) error
	GetImageById(ctx context.Context, id string) ([]byte, string, error)
	GetAll(ctx context.Context) ([]charactermodel.SearchResult, error)
}
