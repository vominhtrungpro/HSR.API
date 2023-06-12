package elements

import (
	"context"

	"github.com/vominhtrungpro/internal/model/model"
)

type Repository interface {
	Create(ctx context.Context, element model.Element) error
	CreatePath(ctx context.Context, path model.Path) error
	GetElementByEnname(ctx context.Context, name string) (model.Element, error)
	GetPathByEnname(ctx context.Context, name string) (model.Path, error)
	UpdateElementImage(ctx context.Context, element model.Element, image []byte) error
	UpdatePathImage(ctx context.Context, path model.Path, image []byte) error
}
