package handler

import (
	"strconv"
	"test-oldo/internal/model"
	"test-oldo/internal/service"
	"test-oldo/pkg"

	"github.com/gofiber/fiber/v2"
)

type PaketHandler struct {
	service *service.PaketService
}

func NewPaketHandler(service *service.PaketService) *PaketHandler {
	return &PaketHandler{service: service}
}

func (h *PaketHandler) Create(c *fiber.Ctx) error {
	paket := new(model.Paket)

	if err := c.BodyParser(paket); err != nil {
		return pkg.ResponseBadRequest(c, err.Error())
	}

	if err := h.service.Create(paket); err != nil {
		return pkg.ResponseBadRequest(c, err.Error())
	}

	return pkg.ResponseCreate(c, paket)
}

func (h *PaketHandler) GetAll(c *fiber.Ctx) error {
	data, err := h.service.GetAll()
	if err != nil {
		return pkg.ResponseError(c, 500, err.Error())
	}

	return pkg.ResponseOK(c, data)
}

func (h *PaketHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return pkg.ResponseBadRequest(c, "invalid id")
	}

	data, err := h.service.GetByID(uint(id))
	if err != nil {
		return pkg.ResponseNotFound(c, "paket tidak ditemukan")
	}

	return pkg.ResponseOK(c, data)
}

func (h *PaketHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return pkg.ResponseBadRequest(c, "invalid id")
	}

	input := new(model.Paket)

	if err := c.BodyParser(input); err != nil {
		return pkg.ResponseBadRequest(c, err.Error())
	}

	data, err := h.service.Update(uint(id), input)
	if err != nil {
		return pkg.ResponseError(c, 400, err.Error())
	}

	return pkg.ResponseOK(c, data)
}

func (h *PaketHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return pkg.ResponseBadRequest(c, "invalid id")
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return pkg.ResponseError(c, 400, err.Error())
	}

	return pkg.ResponseOK(c, "deleted")
}
