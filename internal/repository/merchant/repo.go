package merchant

import (
	"context"

	dbgen "github.com/tekam03/panierquebec-backend/internal/db/gen"
	"github.com/tekam03/panierquebec-backend/internal/model"
)

type Repo interface {
	GetAll(ctx context.Context) ([]*model.Merchant, error)
	GetByID(ctx context.Context, id int32) (*model.Merchant, error)
	Create(ctx context.Context, m *model.MerchantCreate) (*model.Merchant, error)
	Patch(ctx context.Context, id int32, m *model.MerchantPatch) (*model.Merchant, error)
	Delete(ctx context.Context, id int32) error
}

type repo struct {
	querrier *dbgen.Queries
}

func NewRepo(querrier *dbgen.Queries) Repo {
	return &repo{querrier: querrier}
}
