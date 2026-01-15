package externalproduct

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	dbgen "github.com/tekam03/panierquebec-backend/internal/db/gen"
	"github.com/tekam03/panierquebec-backend/internal/model"
)

func (repo *repo) Create(ctx context.Context, ep *model.ExternalProduct) error {
	_, err := repo.querrier.GetExternalProductByExternalID(ctx, ep.ExternalId)
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
	_, err = repo.querrier.CreateExternalProduct(ctx, params)
	return err
}