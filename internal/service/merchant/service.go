package merchant

import (
	"context"

	"github.com/tekam03/panierquebec-backend/internal/model"
	"github.com/tekam03/panierquebec-backend/internal/repository/merchant"
)

type service struct {
	repo merchant.Repo
}

func NewService(repo merchant.Repo) Service {
	return &service{repo: repo}
}

func (s *service) GetAll(ctx context.Context) ([]*model.StoreMerchant, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) GetByID(ctx context.Context, id int32) (*model.StoreMerchant, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) Create(ctx context.Context, m *model.MerchantCreate) (*model.StoreMerchant, error) {
	return s.repo.Create(ctx, m)
}

func (s *service) Patch(ctx context.Context, id int32, patch *model.MerchantPatch) (*model.StoreMerchant, error) {
	return s.repo.Patch(ctx, id, patch)
}

func (s *service) Delete(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}
