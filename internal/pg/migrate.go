package pg

import (
	"database/sql"
	"fmt"

	mfs "amartha-loan-system/db/migrations"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func MigrateUp(client *sql.DB, dbName string) error {
	dir, err := iofs.New(mfs.PGMigrationFS, "pg")
	if err != nil {
		return fmt.Errorf("failed to read embedded FS: %w", err)
	}

	driver, err := postgres.WithInstance(client, &postgres.Config{DatabaseName: dbName})
	if err != nil {
		return fmt.Errorf("failed to get migration driver: %w", err)
	}

	m, err := migrate.NewWithInstance("iofs", dir, dbName, driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

func MigrateDown(client *sql.DB, dbName string) error {
	dir, err := iofs.New(mfs.PGMigrationFS, "pg")
	if err != nil {
		return fmt.Errorf("failed to read embedded FS: %w", err)
	}

	driver, err := postgres.WithInstance(client, &postgres.Config{DatabaseName: dbName})
	if err != nil {
		return fmt.Errorf("failed to get migration driver: %w", err)
	}

	m, err := migrate.NewWithInstance("iofs", dir, dbName, driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	err = m.Down()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

func MigrateDrop(client *sql.DB, dbName string) error {
	dir, err := iofs.New(mfs.PGMigrationFS, "pg")
	if err != nil {
		return fmt.Errorf("failed to read embedded FS: %w", err)
	}

	driver, err := postgres.WithInstance(client, &postgres.Config{DatabaseName: dbName})
	if err != nil {
		return fmt.Errorf("failed to get migration driver: %w", err)
	}

	m, err := migrate.NewWithInstance("iofs", dir, dbName, driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	err = m.Drop()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
