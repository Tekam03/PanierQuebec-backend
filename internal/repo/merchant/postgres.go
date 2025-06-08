package merchant

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tekam03/panierquebec-backend/internal/db"
	"github.com/tekam03/panierquebec-backend/internal/model"
)

type postgresMerchantRepo struct {
    db *pgxpool.Pool
}

func NewPostgresMerchantRepo() MerchantRepo {
    return &postgresMerchantRepo{
        db: db.Pool, // reuse the pool from your db package
    }
}
func (r *postgresMerchantRepo) GetMerchant(ctx context.Context, id int) (*model.StoreMerchant, error) {
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

func (r *postgresMerchantRepo) GetAllMerchants(ctx context.Context) ([]*model.StoreMerchant, error) {
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

func (r *postgresMerchantRepo) CreateMerchant(ctx context.Context, merchant *model.StoreMerchant) error {
	query := `INSERT INTO store_merchants (name, url) VALUES ($1, $2) RETURNING id`
	row := r.db.QueryRow(ctx, query, merchant.Name, merchant.Url)


	if err := row.Scan(&merchant.ID); err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("failed to create merchant: %w", err)
		}
		return fmt.Errorf("failed to scan new merchant ID: %w", err)
	}

	return nil
}

func (r *postgresMerchantRepo) UpdateMerchant(ctx context.Context, id int, merchant *model.StoreMerchant) (error) {
	query := `UPDATE store_merchants SET name = $1, url = $2 WHERE id = $3`
	_, err := r.db.Exec(ctx, query, merchant.Name, merchant.Url, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("merchant with id %d not found", id)
		}
		return fmt.Errorf("failed to update merchant: %w", err)
	}

	return nil
}

func (r *postgresMerchantRepo) DeleteMerchant(ctx context.Context, id int) error {
	query := `DELETE FROM store_merchants WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete merchant: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("merchant with id %d not found", id)
	}

	return nil
}
