package handler

import (
	"strconv"
	"test-oldo/internal/model"
	"test-oldo/internal/service"
	"test-oldo/pkg"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// CREATE
func (h *UserHandler) Create(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return pkg.ResponseError(c, 400, err.Error())
	}

	if err := h.service.Create(user); err != nil {
		return pkg.ResponseError(c, 400, err.Error())
	}

	return pkg.ResponseCreate(c, user)
}

// GET ALL
func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.service.GetAll()
	if err != nil {
		return pkg.ResponseError(c, 500, err.Error())
	}

	return pkg.ResponseOK(c, users)
}

// GET BY ID
func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return pkg.ResponseError(c, 400, "invalid id")
	}

	user, err := h.service.GetByID(uint(id))
	if err != nil {
		return pkg.ResponseNotFound(c, "user tidak ditemukan")
	}

	return pkg.ResponseOK(c, user)
}

// UPDATE
func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return pkg.ResponseError(c, 400, "invalid id")
	}

	input := new(model.User)

	if err := c.BodyParser(input); err != nil {
		return pkg.ResponseError(c, 400, err.Error())
	}

	user, err := h.service.Update(uint(id), input)
	if err != nil {
		return pkg.ResponseError(c, 400, err.Error())
	}

	return pkg.ResponseOK(c, user)
}

// DELETE
func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return pkg.ResponseError(c, 400, "invalid id")
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return pkg.ResponseError(c, 400, err.Error())
	}

	return pkg.ResponseOK(c, "deleted")
}
