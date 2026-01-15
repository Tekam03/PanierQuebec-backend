package externalproduct

import "context"

func (repo *repo) Delete(ctx context.Context, id int) error {
	_, err := repo.querrier.DeleteExternalProduct(ctx, int32(id))
	return err
}