package main

import (
	"fmt"
	"log"

	"github.com/wardvisual/go-layered/internal/app/config"
	"github.com/wardvisual/go-layered/internal/app/provider"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg := config.NewConfig()
	app := config.NewFiber(cfg)
	db := config.NewPostgres(cfg)
	validator := config.NewValidator()
	bootstrap := provider.Provider{App: app, Config: cfg, DB: db, Validator: validator}

	defer func(d *sqlx.DB) {
		if err := d.Close(); err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}(db)

	bootstrap.Provide()

	if err := app.Listen(fmt.Sprintf(":%d", cfg.GetInt("APP_PORT", 8080))); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
