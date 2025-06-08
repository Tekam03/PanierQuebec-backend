package merchant

import (
    "context"
    "github.com/tekam03/panierquebec-backend/internal/model"
    "github.com/tekam03/panierquebec-backend/internal/repo/merchant"
)

type service struct {
    repo merchant.MerchantRepo
}

func NewService(r merchant.MerchantRepo) Service {
    return &service{repo: r}
}

func (s *service) GetAll(ctx context.Context) ([]*model.StoreMerchant, error) {
    return s.repo.GetAllMerchants(ctx)
}

func (s *service) GetByID(ctx context.Context, id int) (*model.StoreMerchant, error) {
    return s.repo.GetMerchant(ctx, id)
}

func (s *service) Create(ctx context.Context, m *model.StoreMerchant) error {
    return s.repo.CreateMerchant(ctx, m)
}

func (s *service) Update(ctx context.Context, id int, m *model.UpdateStoreMerchant) (*model.StoreMerchant, error) {
    return s.repo.UpdateMerchant(ctx, id, m)
}

func (s *service) Delete(ctx context.Context, id int) error {
    return s.repo.DeleteMerchant(ctx, id)
}
