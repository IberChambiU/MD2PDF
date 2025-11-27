package router

import (
	mh "github.com/IberChambiU/MD2PDF/api/markdowntopdf/application"
	m "github.com/IberChambiU/MD2PDF/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func MarkdownToPDFRouter(app fiber.Router, h mh.MarkdownToPDFHandler) {
	mr := app.Group("/api/v1/markdowntopdf")
	mr.Post("", m.Admin(), h.ConvertMarkdownToPDF) // with token
	// mr.Post("", m.BasicAuth(), h.ConvertMarkdownToPDF) // with basic auth
}
