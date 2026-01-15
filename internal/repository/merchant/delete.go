package merchant

import (
	"context"
)

func (r *repo) Delete(ctx context.Context, id int32) error {
    _, err := r.querrier.DeleteStoreMerchant(ctx, id)
    return err
}