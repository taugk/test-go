package handler

import (
	"strconv"
	"test-oldo/internal/model"
	"test-oldo/internal/service"
	"test-oldo/pkg"

	"github.com/gofiber/fiber/v2"
)

type TransaksiHandler struct {
	service *service.TransaksiService
}

func NewTransaksiHandler(service *service.TransaksiService) *TransaksiHandler {
	return &TransaksiHandler{service: service}
}

func (h *TransaksiHandler) Create(c *fiber.Ctx) error {
	req := new(model.Transaksi)

	if err := c.BodyParser(req); err != nil {
		return pkg.ResponseBadRequest(c, err.Error())
	}

	if req.UserID == 0 || req.PaketID == 0 {
		return pkg.ResponseBadRequest(c, "user_id dan paket_id wajib diisi")
	}

	trx, err := h.service.Create(req.UserID, req.PaketID)
	if err != nil {
		return pkg.ResponseError(c, 400, err.Error())
	}

	return pkg.ResponseCreate(c, trx)
}

func (h *TransaksiHandler) GetAll(c *fiber.Ctx) error {
	data, err := h.service.GetAll()
	if err != nil {
		return pkg.ResponseError(c, 500, err.Error())
	}

	return pkg.ResponseOK(c, data)
}

func (h *TransaksiHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return pkg.ResponseBadRequest(c, "invalid id")
	}

	data, err := h.service.GetByID(uint(id))
	if err != nil {
		return pkg.ResponseNotFound(c, "transaksi tidak ditemukan")
	}

	return pkg.ResponseOK(c, data)
}
