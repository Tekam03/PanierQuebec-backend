package store

import (
	"context"

	"github.com/tekam03/panierquebec-backend/internal/model"
	"github.com/tekam03/panierquebec-backend/internal/util"
)

func (r *repo) GetByID(ctx context.Context, id int32) (*model.Store, error) {
	store, err := r.querrier.GetStoreByIDDetailed(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.Store{
		ID:       store.ID,
		Name:     getNameOrFallback(store.Name, store.MerchantName),
		Location: util.FromPtr(store.Location),
	}, nil
}

func (r *repo) GetAll(ctx context.Context) ([]*model.Store, error) {
	// Use the querrier to get all merchants
	// Also convert the result to the model.StoreMerchant type
	stores, err := r.querrier.GetStoresDetailed(ctx)
	if err != nil {
		return nil, err
	}
	var result []*model.Store
	for _, s := range stores {
		result = append(result, &model.Store{
			ID:       s.ID,
			Name:     getNameOrFallback(s.Name, s.MerchantName),
			Location: util.FromPtr(s.Location),
			Merchant: model.Merchant{
				ID:   s.MerchantID,
				Name: util.FromPtr(s.MerchantName),
				Url:  util.FromPtr(s.MerchantUrl),
			},
		})
	}
	return result, nil
}

func getNameOrFallback(name *string, merchantName *string) string {
	var storeName string
	if name == nil {
		if merchantName == nil {
			storeName = "Unknown Store"
		} else {
			storeName = *merchantName
		}
	} else {
		storeName = *name
	}
	return storeName
}
