package application

import (
	"github.com/gofiber/fiber/v2"
)

type MarkdownToPDFHandler interface {
	ConvertMarkdownToPDF(c *fiber.Ctx) error
}
