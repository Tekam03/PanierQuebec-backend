package merchant

import (
	"context"

	"github.com/tekam03/panierquebec-backend/internal/model"
)

type Service interface {
	GetAll(ctx context.Context) ([]*model.StoreMerchant, error)
	GetByID(ctx context.Context, id int32) (*model.StoreMerchant, error)
	Create(ctx context.Context, m *model.MerchantCreate) (*model.StoreMerchant, error)
	Patch(ctx context.Context, id int32, m *model.MerchantPatch) (*model.StoreMerchant, error)
	Delete(ctx context.Context, id int32) error
}
