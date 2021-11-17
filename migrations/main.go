package main

import (
	"embed"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ozonmp/pay-card-api/internal/config"
	"github.com/ozonmp/pay-card-api/internal/database"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	conn, err := database.NewPostgres(cfg.Database.GetDSN(), cfg.Database.Driver)
	defer conn.Close()

	goose.SetBaseFS(embedMigrations)

	const cmd = "up"

	err = goose.Run(cmd, conn.DB, "migrations")
	if err != nil {
		log.Fatal().Err(err).Msg("goose.Status() error")
	}

}
