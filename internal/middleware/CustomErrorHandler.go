package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/maodematos/hi-gofiber/pkg/util"
	"go.mongodb.org/mongo-driver/mongo"
)

func GeneralErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	errorType := fiber.ErrInternalServerError.Message

	var e *fiber.Error
	if errors.As(err, &e) {
		err = util.CustomErrorFromFiber(e)
	}

	if mongoError := mongoErrorsHandler(err); mongoError != nil {
		err = mongoError
	}

	var ce *util.CustomError
	if errors.As(err, &ce) {
		code = ce.Code
		errorType = ce.ErrorType
	}

	return ctx.Status(code).JSON(util.ApiResponse{
		"status":  code,
		"error":   errorType,
		"message": err.Error(),
	})
}

func mongoErrorsHandler(err error) error {
	switch err {
	case mongo.ErrNoDocuments:
		return util.NewCustomError(
			fiber.StatusNotFound,
			fiber.ErrNotFound.Message,
			err.Error(),
		)
	case mongo.ErrInvalidIndexValue:
		return util.NewCustomError(
			fiber.StatusBadRequest,
			fiber.ErrBadRequest.Message,
			err.Error(),
		)
	}

	return nil
}
