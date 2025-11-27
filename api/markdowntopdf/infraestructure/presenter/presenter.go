package presenter

import (
	a "github.com/IberChambiU/MD2PDF/api/markdowntopdf/application"
)

type markdownToPDFPresenter struct {
}

func NewMarkdownToPDFPresenter() a.MarkdownToPDFPresenter {
	return markdownToPDFPresenter{}
}
