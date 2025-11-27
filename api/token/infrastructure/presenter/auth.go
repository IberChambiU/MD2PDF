package presenter

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorResponse is the General ErrorResponse that will be passed in the response by handler
func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"success": false,
		"error":   err.Error(),
	}
}
