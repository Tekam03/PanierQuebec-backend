package merchant

import (
	"context"

	"github.com/tekam03/panierquebec-backend/internal/model"
)

func (r *repo) GetByID(ctx context.Context, id int32) (*model.StoreMerchant, error) {
    merchant, err := r.querrier.GetStoreMerchantByID(ctx, id)
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

func (r *repo) GetAll(ctx context.Context) ([]*model.StoreMerchant, error) {
    // Use the querrier to get all merchants
    // Also convert the result to the model.StoreMerchant type
    merchants, err := r.querrier.GetStoreMerchants(ctx)
    if err != nil {
        return nil, err
    }
    var result []*model.StoreMerchant
    for _, m := range merchants { 
        var url string
        if m.Url != nil {
            url = *m.Url
        }
        result = append(result, &model.StoreMerchant{
            ID:   m.ID,
            Name: m.Name,
            Url:  url,
        })
    }
    return result, nil
}