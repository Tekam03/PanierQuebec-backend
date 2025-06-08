package merchant

import (
	"context"

	"github.com/tekam03/panierquebec-backend/internal/model"
)

type MerchantRepo interface {
	GetMerchant(ctx context.Context, id int) (*model.StoreMerchant, error)
	GetAllMerchants(ctx context.Context) ([]*model.StoreMerchant, error)
	CreateMerchant(ctx context.Context, merchant *model.StoreMerchant) error
	UpdateMerchant(ctx context.Context, id int, merchant *model.UpdateStoreMerchant) (*model.StoreMerchant, error)
	DeleteMerchant(ctx context.Context, id int) error
}
