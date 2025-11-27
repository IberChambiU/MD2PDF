package domain

import (
	"bytes"
	"context"
	"errors"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/IberChambiU/MD2PDF/env"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"github.com/google/uuid"
)

var (
	ErrFileNotExists      = errors.New("Sube un archivo en el campo 'file' (multipart/form-data)")
	ErrFileNotFound       = errors.New("Archivo no encontrado")
	ErrFileIsNotMarkdown  = errors.New("El archivo no es un archivo markdown")
	ErrFileIsNotExtencion = errors.New("El archivo no tiene la extensión .md")
	ErrFileNotConverted   = errors.New("El archivo no se pudo convertir")
	ErrGetPathHtml        = errors.New("No se pudo obtener la ruta del archivo HTML")
	ErrGenPdf             = errors.New("No se pudo generar el PDF")
	ErrDontWritePdf       = errors.New("No se pudo escribir el PDF")
	ErrFileNotCreated     = errors.New("El archivo no fue creado")
)

var (
	PdfDir        = "./pdf"
	MdDir         = "./md"
	HtmlDir       = "./html"
	PdfExtension  = ".pdf"
	MdExtension   = ".md"
	HtmlExtension = ".html"
)

var e = env.Env()

type MarkDownToPDF struct {
	File         *multipart.FileHeader
	Id           uuid.UUID
	Name         string
	Extension    string
	MdFilePath   string
	HtmlFilePath string
	PdfFilePath  string
}

type MarkDownToPdfPresenter struct {
	Id         uuid.UUID `json:"id"`
	OldName    string    `json:"old_name"`
	NewName    string    `json:"new_name"`
	Url        string    `json:"url"`
	Extension  string    `json:"extension"`
	Expiration time.Time `json:"expiration"`
}

func ValidFileMarkdown(file *multipart.FileHeader) error {
	if file == nil {
		return ErrFileNotFound
	}

	if file.Header.Get("Content-Type") != "text/markdown" {
		return ErrFileIsNotMarkdown
	}

	if filepath.Ext(file.Filename) != ".md" {
		return ErrFileIsNotExtencion
	}

	return nil
}

func MakeMarkDownToPDF(file *multipart.FileHeader) (*MarkDownToPDF, error) {

	id := uuid.New()

	return &MarkDownToPDF{
		File:         file,
		Id:           id,
		Name:         file.Filename,
		Extension:    filepath.Ext(file.Filename),
		MdFilePath:   filepath.Join(MdDir, id.String()+".md"),
		HtmlFilePath: filepath.Join(HtmlDir, id.String()+".html"),
		PdfFilePath:  filepath.Join(PdfDir, id.String()+".pdf"),
	}, nil
}

func ConvertMarkDownToPDF(md *MarkDownToPDF) error {

	// Leer Markdown
	mdContent, err := os.ReadFile(md.MdFilePath)
	if err != nil {
		log.Printf("Error: no se pudo leer el archivo Markdown: %v", err)
		return ErrFileNotConverted
	}

	_, err = ConvertMarkDownToHtml(md.HtmlFilePath, mdContent)
	if err != nil {
		log.Printf("Error: no se pudo convertir el archivo Markdown a HTML: %v", err)
		return ErrFileNotConverted
	}

	if err := ConvertHtmlToPDF(md.PdfFilePath, md.HtmlFilePath); err != nil {
		log.Printf("Error: no se pudo convertir el archivo HTML a PDF: %v", err)
		return ErrFileNotConverted
	}

	// Implement the conversion logic here
	return nil
}

func ConvertMarkDownToHtml(htmlFile string, md []byte) (*string, error) {

	ext := parser.CommonExtensions | parser.AutoHeadingIDs | parser.Tables | parser.Footnotes | parser.Strikethrough
	p := parser.NewWithExtensions(ext)
	opts := html.RendererOptions{
		Flags:           html.CommonFlags | html.HrefTargetBlank,
		Title:           "Markdown Document",
		HeadingIDPrefix: "h-",
	}
	renderer := html.NewRenderer(opts)
	body := markdown.ToHTML(md, p, renderer)

	var buf bytes.Buffer
	if _, err := buf.WriteString("<!doctype html><html><head><meta charset=\"utf-8\">"); err != nil {
		return nil, err
	}
	if _, err := buf.WriteString("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">"); err != nil {
		return nil, err
	}
	if _, err := buf.WriteString("<style>"); err != nil {
		return nil, err
	}
	if _, err := buf.WriteString(BaseCSS()); err != nil {
		return nil, err
	}
	if _, err := buf.WriteString("</style>"); err != nil {
		return nil, err
	}
	if _, err := buf.WriteString("</head><body><main class=\"container\">"); err != nil {
		return nil, err
	}
	if _, err := buf.Write(body); err != nil {
		return nil, err
	}
	if _, err := buf.WriteString("</main></body></html>"); err != nil {
		return nil, err
	}
	htmlContent := buf.String()

	// Guardar HTML
	if err := os.WriteFile(htmlFile, []byte(htmlContent), os.ModePerm); err != nil {
		log.Printf("no se pudo escribir HTML: %v", err)
		return nil, err
	}

	// Verificar que el html fue creado
	if _, err := os.Stat(htmlFile); err != nil {
		log.Printf("Error: no se pudo crear el archivo HTML: %v", err)
		return nil, err
	}

	return &htmlContent, nil
}

func ConvertHtmlToPDF(pdfFile string, htmlFile string) error {
	abs, err := filepath.Abs(htmlFile)
	if err != nil {
		return ErrGetPathHtml
	}
	url := "file://" + abs

	// Detectar ruta de Chromium por variable de entorno o usar por defecto
	chromePath := e.GetChromedpPath()
	allocatorOpts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("no-sandbox", true),
	)
	if chromePath == "" {
		// Docker Debian
		if _, err := os.Stat("/usr/bin/chromium"); err == nil {
			chromePath = "/usr/bin/chromium"
		} else if _, err := os.Stat("/usr/bin/chromium-browser"); err == nil {
			chromePath = "/usr/bin/chromium-browser"
		}
	}
	if chromePath != "" {
		allocatorOpts = append(allocatorOpts, chromedp.ExecPath(chromePath))
	}

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), allocatorOpts...)
	defer cancel()
	ctx, cancel2 := chromedp.NewContext(allocCtx)
	defer cancel2()
	ctx, cancel3 := context.WithTimeout(ctx, 60*time.Second)
	defer cancel3()

	var pdfBuf []byte
	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfBuf, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithMarginTop(0.4).
				WithMarginBottom(0.6).
				WithMarginLeft(0.5).
				WithMarginRight(0.5).
				WithPaperWidth(8.27).
				WithPaperHeight(11.69).
				Do(ctx)
			return err
		}),
	}
	if err := chromedp.Run(ctx, tasks); err != nil {
		return ErrGenPdf
	}

	// Guardar PDF
	if err := os.WriteFile(pdfFile, pdfBuf, 0o644); err != nil {
		return ErrDontWritePdf
	}

	// Verificar que el PDF fue creado
	if _, err := os.Stat(pdfFile); err != nil {
		return ErrFileNotCreated
	}

	return nil
}
