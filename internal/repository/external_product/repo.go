package externalproduct

import (
	"context"

	"github.com/tekam03/panierquebec-backend/internal/db/gen"
	"github.com/tekam03/panierquebec-backend/internal/model"
)

type Repo interface {
	GetAll(ctx context.Context) ([]*model.ExternalProduct, error)
	Create(ctx context.Context, ep *model.ExternalProduct) error
	Delete(ctx context.Context, id int) error
}

type repo struct {
	querrier *dbgen.Queries
}

func NewRepo(querrier *dbgen.Queries) Repo {
	return &repo{querrier: querrier}
}
