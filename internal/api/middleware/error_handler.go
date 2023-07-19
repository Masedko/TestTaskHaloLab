package middleware

import (
	"TestTaskHaloLab/internal/core/dtos"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}

	err = c.Status(code).JSON(dtos.MessageResponse{Message: message})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Cannot send error JSON message")
	}

	return nil
}
