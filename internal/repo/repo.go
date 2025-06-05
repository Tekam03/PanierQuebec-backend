package repo

import (
	"context"

	"github.com/tekam03/panierquebec-backend/internal/model"
)

// type StoreRepo interface {

// }

type MerchantRepo interface {
	GetMerchant(ctx context.Context, id int) (*model.StoreMerchant, error)

	GetAllMerchants(ctx context.Context) ([]*model.StoreMerchant, error)

	// CreateMerchant(ctx context.Context, merchant *StoreMerchant) error

	// UpdateMerchant(ctx context.Context, merchant *StoreMerchant) error

	// DeleteMerchant(ctx context.Context, id int) error
}
