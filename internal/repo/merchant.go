package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tekam03/panierquebec-backend/internal/db"
	"github.com/tekam03/panierquebec-backend/internal/model"
)

type merchantRepo struct {
    db *pgxpool.Pool
}

func NewMerchantRepo() MerchantRepo {
    return &merchantRepo{
        db: db.Pool, // reuse the pool from your db package
    }
}
func (r *merchantRepo) GetMerchant(ctx context.Context, id int) (*model.StoreMerchant, error) {
	query := `SELECT id, name, url FROM store_merchants WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	m := &model.StoreMerchant{}
	if err := row.Scan(&m.ID, &m.Name, &m.Url); err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("merchant with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to scan merchant: %w", err)
	}

	return m, nil
}

func (r *merchantRepo) GetAllMerchants(ctx context.Context) ([]*model.StoreMerchant, error) {
	query := `SELECT id, name, url FROM store_merchants ORDER BY id ASC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query merchants: %w", err)
	}
	defer rows.Close()

	var merchants []*model.StoreMerchant
    for rows.Next() {
    	m := &model.StoreMerchant{}
        if err := rows.Scan(&m.ID, &m.Name, &m.Url); err != nil {
            return nil, fmt.Errorf("failed to scan merchant: %w", err)
        }
        merchants = append(merchants, m)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("rows iteration error: %w", err)
    }

    return merchants, nil
}
