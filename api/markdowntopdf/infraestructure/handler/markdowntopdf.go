package handler

import (
	d "github.com/IberChambiU/MD2PDF/api/markdowntopdf/domain"
	"github.com/gofiber/fiber/v2"
)

func (h *markdownToPDFHandler) ConvertMarkdownToPDF(c *fiber.Ctx) error {

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.presenter.ErrorResponse(d.ErrFileNotExists))
	}

	if err := d.ValidFileMarkdown(file); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.presenter.ErrorResponse(err))
	}

	// Convertir el archivo markdown a PDF
	md2pdf, err := h.service.MarkdownToPDF(file)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(h.presenter.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(h.presenter.SuccessResponse(md2pdf))
}
