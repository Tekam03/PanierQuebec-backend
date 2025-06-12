package migrate

import (
	"fmt"

	"github.com/tekam03/panierquebec-backend/internal/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/stdlib"
)

func RunMigrations(migrationsPath string) error {
	db := stdlib.OpenDBFromPool(db.Pool)
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("create migrate driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver)
	if err != nil {
		return fmt.Errorf("create migrate instance: %w", err)
	}
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("No migrations to apply")
			return nil
		} else {
			return fmt.Errorf("run migrate up: %w", err)
		}
	}
	fmt.Println("Migrations applied successfully")
	return nil
}
