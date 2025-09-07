package tasks

import (
	"github.com/wardvisual/go-layered/internal/app/config"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/wardvisual/go-layered/internal/features/tasks/internal/router"
)

func Provide(app *fiber.App, db *sqlx.DB, validator *config.Validator) {
	router.Route(app, db, validator)
}
