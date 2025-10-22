package rest

import "github.com/gofiber/fiber/v2"

// ErrorMessage for general message
func ErrorMessage(ctx *fiber.Ctx, status int, err error) error {
	return ctx.Status(status).JSON(&fiber.Map{
		"error": err.Error(),
	})
}

func InternalError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
		"error": err.Error(),
	})
}

func BadRequestError(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
		"message": msg,
	})
}

func ResponseMessage(ctx *fiber.Ctx, message string, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": message,
		"data":    data,
	})
}
