package merchant

import (
	"context"

	"github.com/tekam03/panierquebec-backend/internal/db/gen"
	"github.com/tekam03/panierquebec-backend/internal/model"
)

func (r *repo) Create(ctx context.Context, m *model.MerchantCreate) (*model.StoreMerchant, error) {
    params := dbgen.CreateStoreMerchantParams{
        Name: m.Name,
        Url:  m.Url,
    }
    merchant, err := r.querrier.CreateStoreMerchant(ctx, params)
    if err != nil {
        return nil, err
    }
    var url string
    if merchant.Url != nil {
        url = *merchant.Url
    }
    return &model.StoreMerchant{
        ID:   merchant.ID,
        Name: merchant.Name,
        Url:  url,
    }, nil
}	
