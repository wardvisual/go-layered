package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	appconfig "github.com/wardvisual/go-layered/internal/app/config"
	apperr "github.com/wardvisual/go-layered/internal/app/exception"
	appmodel "github.com/wardvisual/go-layered/internal/app/model"
	"github.com/wardvisual/go-layered/internal/features/tasks/internal/entity"
	"github.com/wardvisual/go-layered/internal/features/tasks/internal/model"
	"github.com/wardvisual/go-layered/internal/features/tasks/internal/usecase"
)

type TaskController struct {
	UseCase   *usecase.TaskUseCase
	Validator *appconfig.Validator
}

func NewTaskController(useCase *usecase.TaskUseCase, validator *appconfig.Validator) *TaskController {
	return &TaskController{UseCase: useCase, Validator: validator}
}

func (c *TaskController) Create(ctx *fiber.Ctx) error {
	var req model.TaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		return apperr.PanicIfError(err)
	}

	if vErrs := c.Validator.Validate(req); len(vErrs) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(appmodel.Response{Code: 400, Status: "Bad Request", Data: c.Validator.Message(vErrs)})
	}

	t := &entity.Task{Title: req.Title}
	if req.Done != nil {
		t.Done = *req.Done
	}

	if err := c.UseCase.Create(t); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(appmodel.Response{Code: 500, Status: "Error", Data: err.Error()})
	}

	res := model.TaskResponse{Id: t.Id.String(), Title: t.Title, Done: t.Done}
	return ctx.Status(fiber.StatusCreated).JSON(appmodel.Response{Code: 201, Status: "Created", Data: res})
}

func (c *TaskController) List(ctx *fiber.Ctx) error {
	items, err := c.UseCase.List()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(appmodel.Response{Code: 500, Status: "Error", Data: err.Error()})
	}

	res := make([]model.TaskResponse, 0, len(items))
	for _, it := range items {
		res = append(res, model.TaskResponse{Id: it.Id.String(), Title: it.Title, Done: it.Done})
	}

	return ctx.JSON(appmodel.Response{Code: 200, Status: "OK", Data: res})
}

func (c *TaskController) Get(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(appmodel.Response{Code: 400, Status: "Bad Request", Data: "invalid id"})
	}

	item, err := c.UseCase.Get(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(appmodel.Response{Code: 404, Status: "Not Found", Data: err.Error()})
	}

	res := model.TaskResponse{Id: item.Id.String(), Title: item.Title, Done: item.Done}
	return ctx.JSON(appmodel.Response{Code: 200, Status: "OK", Data: res})
}

func (c *TaskController) Update(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(appmodel.Response{Code: 400, Status: "Bad Request", Data: "invalid id"})
	}

	var req model.TaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(appmodel.Response{Code: 400, Status: "Bad Request", Data: err.Error()})
	}

	if vErrs := c.Validator.Validate(req); len(vErrs) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(appmodel.Response{Code: 400, Status: "Bad Request", Data: c.Validator.Message(vErrs)})
	}

	item := entity.Task{Id: id, Title: req.Title}
	if req.Done != nil {
		item.Done = *req.Done
	}

	if err := c.UseCase.Update(&item); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(appmodel.Response{Code: 500, Status: "Error", Data: err.Error()})
	}

	res := model.TaskResponse{Id: item.Id.String(), Title: item.Title, Done: item.Done}
	return ctx.JSON(appmodel.Response{Code: 200, Status: "OK", Data: res})
}

func (c *TaskController) Delete(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(appmodel.Response{Code: 400, Status: "Bad Request", Data: "invalid id"})
	}

	if err := c.UseCase.Delete(id); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(appmodel.Response{Code: 404, Status: "Not Found", Data: err.Error()})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
