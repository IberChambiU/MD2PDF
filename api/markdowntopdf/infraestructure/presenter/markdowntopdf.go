package presenter

import (
	"log"
	"time"

	d "github.com/IberChambiU/MD2PDF/api/markdowntopdf/domain"
	env "github.com/IberChambiU/MD2PDF/env"
)

// var e = env.Env()

func (m2p markdownToPDFPresenter) SuccessResponse(pdf *d.MarkDownToPDF) map[string]any {

	log.Print("url pdf: ", env.Env().GetUrlServidor()+"/api/v1/pdf/"+pdf.Id.String()+d.PdfExtension)

	MarkDownToPdfPresenter := d.MarkDownToPdfPresenter{
		Id:         pdf.Id,
		OldName:    pdf.Name,
		NewName:    pdf.Id.String() + d.PdfExtension,
		Url:        env.Env().GetUrlServidor() + "/api/v1/pdf/" + pdf.Id.String() + d.PdfExtension,
		Extension:  pdf.Extension,
		Expiration: time.Now().Add(time.Duration(env.Env().GetPdfCleanupFileMaxAgeMinutes()) * time.Minute),
	}

	return map[string]any{
		"data":    MarkDownToPdfPresenter,
		"success": true,
	}
}
func (m2p markdownToPDFPresenter) ErrorResponse(err error) map[string]any {
	return map[string]any{
		"error":   err.Error(),
		"success": false,
	}
}
