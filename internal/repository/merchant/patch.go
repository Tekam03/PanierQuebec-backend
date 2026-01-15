package merchant

import (
	"context"

	"github.com/tekam03/panierquebec-backend/internal/db/gen"
	"github.com/tekam03/panierquebec-backend/internal/model"
	"github.com/tekam03/panierquebec-backend/internal/util"
)

func (r *repo) Patch(ctx context.Context, id int32, m *model.MerchantPatch) (*model.StoreMerchant, error) {
    params := dbgen.PatchStoreMerchantParams{
        ID:   id,
        Name: m.Name,
        Url:  m.Url,
    }
    updatedMerchant, err := r.querrier.PatchStoreMerchant(ctx, params)
    if err != nil {
        return nil, err
    }

    return &model.StoreMerchant{
        ID:   updatedMerchant.ID,
        Name: updatedMerchant.Name,
        Url:  util.FromPtr(updatedMerchant.Url),
    }, nil
}
