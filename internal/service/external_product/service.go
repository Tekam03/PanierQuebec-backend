package external_product

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	dbgen "github.com/tekam03/panierquebec-backend/internal/db/gen"
	"github.com/tekam03/panierquebec-backend/internal/model"
)

type service struct {
	querrier dbgen.Querier
}

func NewService(q dbgen.Querier) Service {
	return &service{querrier: q}
}

func (s *service) GetAll(ctx context.Context) ([]*model.ExternalProduct, error) {
	// Use the querrier to get all external products
	// Also convert the result to the model.ExternalProduct type
	external_products, err := s.querrier.GetExternalProductsDetailed(ctx)
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

func (s *service) Create(ctx context.Context, ep *model.ExternalProduct) error {
	_, err := s.querrier.GetExternalProductByExternalID(ctx, ep.ExternalId)
	if err == nil {
		// Record EXISTS
		return fmt.Errorf("external product already exists")
	}

	if !errors.Is(err, pgx.ErrNoRows) {
		// Real error (DB down, syntax, etc.)
		log.Printf("could not check external product: %v", err)
		return err
	}

	var matchedID *int32
	if ep.StoreProduct != nil && ep.StoreProduct.ID != 0 {
		matchedID = &ep.StoreProduct.ID
	}

	params := dbgen.CreateExternalProductParams{
		Source:                ep.Source,
		ExternalID:            ep.ExternalId,
		Name:                  ep.Name,
		Description:           ep.Description,
		Brand:                 ep.Brand,
		MatchedStoreProductID: matchedID,
	}
	_, err = s.querrier.CreateExternalProduct(ctx, params)
	return err
}

func (s *service) Delete(ctx context.Context, id int) error {
	_, err := s.querrier.DeleteExternalProduct(ctx, int32(id))
	return err
}
