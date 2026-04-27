package pkg

import "github.com/gofiber/fiber/v2"


func ResponseOK(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "OK",
		"data":   data,
	})
}


func ResponseCreate(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "OK",
		"data":   data,
	})
}


func ResponseError(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(fiber.Map{
		"status":  "ERROR",
		"message": message,
	})
}


func ResponseNotFound(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"status":  "ERROR",
		"message": message,
	})
}


func ResponseBadRequest(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status":  "ERROR",
		"message": message,
	})
}