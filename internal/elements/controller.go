package elements

import (
	"context"

	"github.com/vominhtrungpro/internal/model/model"
)

// News use case
type Controller interface {
	Create(ctx context.Context, element model.Element) error
	CreatePath(ctx context.Context, path model.Path) error
	UpdateElementImage(ctx context.Context, name string, image []byte) error
	UpdatePathImage(ctx context.Context, name string, image []byte) error
}
