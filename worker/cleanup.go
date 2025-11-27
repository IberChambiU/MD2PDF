package worker

import (
	"time"

	"github.com/IberChambiU/MD2PDF/api/markdowntopdf/application"
	"github.com/IberChambiU/MD2PDF/env"
)

func StartCleanupWorker() {

	e := env.Env()

	workerInterval := time.Duration(e.GetPdfCleanupIntervalMinutes()) * time.Minute

	fileMaxAge := time.Duration(e.GetPdfCleanupFileMaxAgeMinutes()) * time.Minute

	go func() {
		ticker := time.NewTicker(workerInterval)
		defer ticker.Stop()
		for {
			application.CleanOldFilesFromPDFDir(fileMaxAge)
			<-ticker.C
		}
	}()
}
