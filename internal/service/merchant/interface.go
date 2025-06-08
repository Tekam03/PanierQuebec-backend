package merchant

import (
    "context"
    "github.com/tekam03/panierquebec-backend/internal/model"
)

type Service interface {
    GetAll(ctx context.Context) ([]*model.StoreMerchant, error)
    GetByID(ctx context.Context, id int) (*model.StoreMerchant, error)
    Create(ctx context.Context, m *model.StoreMerchant) error
    Update(ctx context.Context, id int, m *model.UpdateStoreMerchant) (*model.StoreMerchant, error)
    Delete(ctx context.Context, id int) error
}
