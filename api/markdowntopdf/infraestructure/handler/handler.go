package handler

import (
	a "github.com/IberChambiU/MD2PDF/api/markdowntopdf/application"
)

type markdownToPDFHandler struct {
	service   a.MarkdownToPDFService
	presenter a.MarkdownToPDFPresenter
}

func NewMarkdownToPDFHandler(service a.MarkdownToPDFService, presenter a.MarkdownToPDFPresenter) a.MarkdownToPDFHandler {
	return &markdownToPDFHandler{
		service:   service,
		presenter: presenter,
	}
}
