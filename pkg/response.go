package pkg

import "github.com/gofiber/fiber/v2"

func Success(c *fiber.Ctx, data interface{}, message ...string) error {
	msg := "Success"
	if len(message) > 0 {
		msg = message[0]
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"data":    data,
		"message": msg,
	})
}

func Created(c *fiber.Ctx, data interface{}, message ...string) error {
	msg := "Created"
	if len(message) > 0 {
		msg = message[0]
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"data":    data,
		"message": msg,
	})
}

func Error(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(fiber.Map{
		"status":  "error",
		"message": message,
	})
}
