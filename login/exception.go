package main

import (
	"errors"

	"github.com/gofiber/fiber/v3"
)

var GlobalErrorHandler = func(c fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "internal server error"

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = err.Error()
	}

	return c.Status(code).JSON(HandlerResponse{
		Success: false,
		Message: message,
	})
}
