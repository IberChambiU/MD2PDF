package env

import (
	"os"
	"strconv"
	"sync"
)

var (
	MuApp = sync.Mutex{}
)

// environmentApp contains the environment variables for the application
type environmentApp struct {
	port                          string
	portOK                        bool
	url                           string
	urlOK                         bool
	secure                        string
	secureOK                      bool
	UrlServidor                   string
	UrlServidorOK                 bool
	ChromedpPath                  string
	ChromedpPathOK                bool
	PdfCleanupIntervalMinutes     string
	PdfCleanupIntervalMinutesOK   bool
	PdfCleanupFileMaxAgeMinutes   string
	PdfCleanupFileMaxAgeMinutesOK bool
	SimetricKey                   string
	SimetricKeyOK                 bool
	// ...existing code...
}

// InitApp initializes the environment variables for the application
//
// if the variables are not set default values will be used
func (e *environment) initApp(wg *sync.WaitGroup) {
	MuApp.Lock()
	e.port, e.portOK = os.LookupEnv("PORT")
	e.url, e.urlOK = os.LookupEnv("URL")
	e.secure, e.secureOK = os.LookupEnv("SECURE")
	e.UrlServidor, e.UrlServidorOK = os.LookupEnv("URL_SERVIDOR")
	e.ChromedpPath, e.ChromedpPathOK = os.LookupEnv("CHROMEDP_PATH")
	e.PdfCleanupIntervalMinutes, e.PdfCleanupIntervalMinutesOK = os.LookupEnv("PDF_CLEANUP_INTERVAL_MINUTES")
	e.PdfCleanupFileMaxAgeMinutes, e.PdfCleanupFileMaxAgeMinutesOK = os.LookupEnv("PDF_CLEANUP_FILE_MAX_AGE_MINUTES")
	e.SimetricKey, e.SimetricKeyOK = os.LookupEnv("SIMETRIC_KEY")
	MuApp.Unlock()
	wg.Done()

}

// GetPdfCleanupFileMaxAgeMinutes retorna el valor de la variable PDF_CLEANUP_FILE_MAX_AGE_MINUTES como int, default 60
func (e *environment) GetPdfCleanupFileMaxAgeMinutes() int {
	MuApp.Lock()
	defer MuApp.Unlock()
	if !e.PdfCleanupFileMaxAgeMinutesOK {
		return 60
	}
	val, err := strconv.Atoi(e.PdfCleanupFileMaxAgeMinutes)
	if err != nil || val <= 0 {
		return 60
	}
	return val
}

// GetPdfCleanupIntervalMinutes retorna el valor de la variable PDF_CLEANUP_INTERVAL_MINUTES como int, default 24
func (e *environment) GetPdfCleanupIntervalMinutes() int {
	MuApp.Lock()
	defer MuApp.Unlock()
	if !e.PdfCleanupIntervalMinutesOK {
		return 24
	}
	val, err := strconv.Atoi(e.PdfCleanupIntervalMinutes)
	if err != nil || val <= 0 {
		return 24
	}
	return val
}

// `PORT`: Puerto en el que se ejecutará el servidor.
//
// Por defecto es `3000`.
func (e *environment) GetPort() (val string) {
	MuApp.Lock()
	defer MuApp.Unlock()
	if !e.portOK {
		val = "3030"
	} else {
		val = e.port
	}
	return
}

// `URL`: Dirección url base del servidor.
//
// Por defecto es `https://localhost`.
func (e *environment) GetUrl() (val string) {
	MuApp.Lock()
	defer MuApp.Unlock()
	if !e.urlOK {
		val = "http://localhost"
	} else {
		val = e.url
	}
	return
}

// `SECURE`: habilita la encriptación de los datos mediante TLS.
//
// Por defecto es `false`.
func (e *environment) GetSecure() (val bool) {
	MuApp.Lock()
	defer MuApp.Unlock()
	if e.secureOK {
		val = e.secure == "true"
	}
	return
}

// `URL_SERVIDOR`: Dirección url del servidor de archivos.
// / Por defecto es `http://localhost:4000`.
func (e *environment) GetUrlServidor() (val string) {
	MuApp.Lock()
	defer MuApp.Unlock()
	if !e.UrlServidorOK {
		val = "http://localhost:3030"
	} else {
		val = e.UrlServidor
	}
	return
}

// `CHROMEDP_PATH`: Ruta al ejecutable de Chromium/Chrome.
//
// Por defecto se buscará en las rutas comunes del sistema.
func (e *environment) GetChromedpPath() (val string) {
	MuApp.Lock()
	defer MuApp.Unlock()
	if !e.ChromedpPathOK {
		val = ""
		// val = "/usr/bin/chromium"
	} else {
		val = e.ChromedpPath
	}
	return
}

func (e *environment) GetSimetricKey() (val string) {
	MuApp.Lock()
	defer MuApp.Unlock()
	if !e.SimetricKeyOK {
		val = "ComprobemosSiAlguienPuedeLeerlos"
	} else {
		val = e.SimetricKey
	}
	return
}
