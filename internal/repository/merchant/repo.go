package merchant

import (
	"context"

	"github.com/tekam03/panierquebec-backend/internal/db/gen"
	"github.com/tekam03/panierquebec-backend/internal/model"
)

type Repo interface {
    GetAll(ctx context.Context) ([]*model.StoreMerchant, error)
    GetByID(ctx context.Context, id int32) (*model.StoreMerchant, error)
	Create(ctx context.Context, m *MerchantCreate) (*model.StoreMerchant, error)
	Patch(ctx context.Context, id int32, m *MerchantPatch) (*model.StoreMerchant, error)
	Delete(ctx context.Context, id int32) error
}

type repo struct {
	querrier *dbgen.Queries
}

func NewRepo(querrier *dbgen.Queries) Repo {
	return &repo{querrier: querrier}
}
