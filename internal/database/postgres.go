package database

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/pay-card-api/internal/pkg/logger"
)

// NewPostgres returns DB
func NewPostgres(dsn, driver string) (*sqlx.DB, error) {
	ctx := context.Background()
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		logger.FatalKV(ctx, "failed to create database connection", "err", err)

		return nil, err
	}

	if err = db.Ping(); err != nil {
		logger.FatalKV(ctx, "Failed ping database", "err", err)

		return nil, err
	}

	return db, nil
}
