package merchant

import (
	"context"

	"github.com/tekam03/panierquebec-backend/internal/model"
	"github.com/tekam03/panierquebec-backend/internal/repository/merchant"
)

type Service interface {
    GetAll(ctx context.Context) ([]*model.StoreMerchant, error)
    GetByID(ctx context.Context, id int32) (*model.StoreMerchant, error)
    Create(ctx context.Context, m *merchant.MerchantCreate) (*model.StoreMerchant, error)
    Patch(ctx context.Context, id int32, m *merchant.MerchantPatch) (*model.StoreMerchant, error)
    Delete(ctx context.Context, id int32) error
}
