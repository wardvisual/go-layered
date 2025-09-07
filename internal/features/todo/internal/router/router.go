package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	appconfig "github.com/wardvisual/go-layered/internal/app/config"
	controller "github.com/wardvisual/go-layered/internal/modules/todo/internal/controller"
	inmemory "github.com/wardvisual/go-layered/internal/modules/todo/internal/repository/inmemory"
	"github.com/wardvisual/go-layered/internal/modules/todo/internal/usecase"
)

func Route(app *fiber.App, _ *sqlx.DB, validator *appconfig.Validator) {
	group := app.Group("/v1/todo")

	repo := inmemory.NewTaskRepository()
	uc := usecase.NewTaskUseCase(repo)
	ctrl := controller.NewTaskController(uc, validator)

	group.Post("/", ctrl.Create)
	group.Get("/", ctrl.List)
	group.Get("/:id", ctrl.Get)
	group.Put("/:id", ctrl.Update)
	group.Delete("/:id", ctrl.Delete)
}
