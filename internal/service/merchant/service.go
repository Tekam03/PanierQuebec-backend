package merchant

import (
	"context"

	"github.com/tekam03/panierquebec-backend/internal/model"
    "github.com/tekam03/panierquebec-backend/internal/db/gen"
)

type service struct {
	querrier dbgen.Querier
}

func NewService(q dbgen.Querier) Service {
	return &service{querrier: q}
}

func (s *service) GetAll(ctx context.Context) ([]*model.StoreMerchant, error) {
    // Use the querrier to get all merchants
    // Also convert the result to the model.StoreMerchant type
    merchants, err := s.querrier.GetStoreMerchants(ctx)
    if err != nil {
        return nil, err
    }
    var result []*model.StoreMerchant
    for _, m := range merchants {
        var name string
        if m.Name != nil {
            name = *m.Name
        }
        var url string
        if m.Url != nil {
            url = *m.Url
        }
        result = append(result, &model.StoreMerchant{
            ID:   m.ID,
            Name: name,
            Url:  url,
        })
    }
    return result, nil
}

func (s *service) GetByID(ctx context.Context, id int) (*model.StoreMerchant, error) {
    merchant, err := s.querrier.GetStoreMerchantByID(ctx, int32(id))
    if err != nil {
        return nil, err
    }
    var name string
    if merchant.Name != nil {
        name = *merchant.Name
    }
    var url string
    if merchant.Url != nil {
        url = *merchant.Url
    }
    return &model.StoreMerchant{
        ID:   merchant.ID,
        Name: name,
        Url:  url,
    }, nil
}

func (s *service) Create(ctx context.Context, m *model.StoreMerchant) error {
    params := dbgen.CreateStoreMerchantParams{
        Name: &m.Name,
        Url:  &m.Url,
    }
    _, err := s.querrier.CreateStoreMerchant(ctx, params)
    return err
}

func (s *service) Update(ctx context.Context, id int, m *model.UpdateStoreMerchant) (*model.StoreMerchant, error) {
    params := dbgen.UpdateStoreMerchantParams{
        ID:   int32(id),
        Name: m.Name,
        Url:  m.Url,
    }
    updatedMerchant, err := s.querrier.UpdateStoreMerchant(ctx, params)
    if err != nil {
        return nil, err
    }
    var name string
    if updatedMerchant.Name != nil {
        name = *updatedMerchant.Name
    }
    var url string
    if updatedMerchant.Url != nil {
        url = *updatedMerchant.Url
    }
    return &model.StoreMerchant{
        ID:   updatedMerchant.ID,
        Name: name,
        Url:  url,
    }, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
    _, err := s.querrier.DeleteStoreMerchant(ctx, int32(id))
    return err
}
