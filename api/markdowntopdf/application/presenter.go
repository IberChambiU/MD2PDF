package application

import (
	d "github.com/IberChambiU/MD2PDF/api/markdowntopdf/domain"
)

type MarkdownToPDFPresenter interface {
	SuccessResponse(pdf *d.MarkDownToPDF) map[string]any
	ErrorResponse(err error) map[string]any
}
