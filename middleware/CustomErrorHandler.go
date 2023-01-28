package middleware

import (
	"errors"
	"hi-gofiber/types"

	"github.com/gofiber/fiber/v2"
)

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return ctx.Status(code).JSON(types.ApiResponse{
		"message": err.Error(),
	})
}
