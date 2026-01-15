package external_product

import (
	"context"

	"github.com/tekam03/panierquebec-backend/internal/model"
	externalproduct "github.com/tekam03/panierquebec-backend/internal/repository/external_product"
)

type service struct {
	repo externalproduct.Repo
}

func NewService(repo externalproduct.Repo) Service {
	return &service{repo: repo}
}

func (s *service) GetAll(ctx context.Context) ([]*model.ExternalProduct, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) Create(ctx context.Context, ep *model.ExternalProduct) error {
	return s.repo.Create(ctx, ep)
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
