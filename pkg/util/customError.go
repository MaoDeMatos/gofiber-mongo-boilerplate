package util

import "github.com/gofiber/fiber/v2"

type CustomError struct {
	Code      int    `json:"code"`
	ErrorType string `json:"error"`
	Message   string `json:"message"`
}

func (ce CustomError) Error() string {
	return ce.Message
}

func CustomErrorFromFiber(err *fiber.Error) *CustomError {
	return &CustomError{
		Code:      err.Code,
		ErrorType: err.Message,
		Message:   err.Message,
	}
}

func NewCustomError(code int, errorType, message string) *CustomError {
	return &CustomError{
		Code:      code,
		ErrorType: errorType,
		Message:   message,
	}
}
