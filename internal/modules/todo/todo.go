package todo

import (
	"github.com/wardvisual/go-layered/internal/app/config"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/wardvisual/go-layered/internal/modules/todo/internal/router"
)

func Provide(app *fiber.App, db *sqlx.DB, validator *config.Validator) {
	router.Route(app, db, validator)
}
