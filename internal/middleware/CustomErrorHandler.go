package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/maodematos/hi-gofiber/types"
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
