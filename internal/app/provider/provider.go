package provider

import (
	"github.com/wardvisual/go-layered/internal/app/config"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/wardvisual/go-layered/internal/features/tasks"
)

type Provider struct {
	App       *fiber.App
	Config    *config.Config
	DB        *sqlx.DB
	Validator *config.Validator
}

func (p *Provider) Provide() {
	// Register application modules here
	tasks.Provide(p.App, p.DB, p.Validator)
}
