package external_product

import (
	"context"

	"github.com/tekam03/panierquebec-backend/internal/model"
)

type Service interface {
	GetAll(ctx context.Context) ([]*model.ExternalProduct, error)
	Create(ctx context.Context, ep *model.ExternalProduct) error
	Delete(ctx context.Context, id int) error
}
