package router

import (
	ma "github.com/IberChambiU/MD2PDF/api/markdowntopdf/application"
	mh "github.com/IberChambiU/MD2PDF/api/markdowntopdf/infraestructure/handler"
	mp "github.com/IberChambiU/MD2PDF/api/markdowntopdf/infraestructure/presenter"
	r "github.com/IberChambiU/MD2PDF/api/markdowntopdf/router"
	"github.com/gofiber/fiber/v2"
)

func ApiRouter(app *fiber.App) {

	var (
		markdowntopdfService ma.MarkdownToPDFService   = ma.NewMarkdownToPDFService()
		markdownPresenter    ma.MarkdownToPDFPresenter = mp.NewMarkdownToPDFPresenter()
		markdowntopdfHandler ma.MarkdownToPDFHandler   = mh.NewMarkdownToPDFHandler(markdowntopdfService, markdownPresenter)
	)

	r.MarkdownToPDFRouter(app, markdowntopdfHandler)

}
