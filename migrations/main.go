package main

import (
	"context"
	"embed"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ozonmp/pay-card-api/internal/config"
	"github.com/ozonmp/pay-card-api/internal/database"
	"github.com/ozonmp/pay-card-api/internal/pkg/logger"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	ctx := context.Background()
	if err := config.ReadConfigYML("config.yml"); err != nil {
		logger.FatalKV(ctx, "Failed init configuration", "err", err)
	}
	cfg := config.GetConfigInstance()

	conn, err := database.NewPostgres(cfg.Database.GetDSN(), cfg.Database.Driver)
	defer conn.Close()

	goose.SetBaseFS(embedMigrations)

	const cmd = "up"

	err = goose.Run(cmd, conn.DB, "migrations")
	if err != nil {
		logger.FatalKV(ctx, "goose.Status() error", "err", err)
	}

}
