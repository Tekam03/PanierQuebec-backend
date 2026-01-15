package externalproduct

import (
	"context"
	"log"

	"github.com/tekam03/panierquebec-backend/internal/model"
)

func (repo *repo) GetAll(ctx context.Context) ([]*model.ExternalProduct, error) {
	// Use the querrier to get all external products
	// Also convert the result to the model.ExternalProduct type
	external_products, err := repo.querrier.GetExternalProductsDetailed(ctx)
	if err != nil {
		return nil, err
	}
	var result []*model.ExternalProduct
	for _, product := range external_products {
		var price *float64
		if product.SpPrice.Valid {
			pricef64, err := product.SpPrice.Float64Value()
			if err != nil {
				log.Printf("could not parse price for external product ID %d: %v", product.EpID, err)
			} else {
				price = &pricef64.Float64
			}
		}

		var sp *model.StoreProduct
		if product.SpID != nil {
			sp = &model.StoreProduct{
				ID:                *product.SpID,
				StoreId:           *product.SpStoreID,
				SpecificProductId: *product.SpSpecificProductsID,
				Price:             price,
				LastUpdated:       product.SpLastUpdated.Time,
			}
		}

		result = append(result, &model.ExternalProduct{
			ID:           product.EpID,
			Source:       product.EpSource,
			ExternalId:   product.EpExternalID,
			Name:         product.EpName,
			Description:  product.EpDescription,
			Brand:        product.EpBrand,
			ScrapedAt:    product.EpScrapedAt.Time,
			StoreProduct: sp,
		})
	}
	return result, nil
}