package utils_test

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func TestMarkdownToPDF_FullFlow(t *testing.T) {
	// Rutas de archivos
	mdDir := "../md"
	pdfDir := "../pdf"
	mdFile := filepath.Join(mdDir, "test2.md")
	htmlFile := filepath.Join(pdfDir, "test2.html")
	pdfFile := filepath.Join(pdfDir, "test2.pdf")

	// Crear carpeta pdf si no existe
	if err := os.MkdirAll(pdfDir, 0o755); err != nil {
		t.Fatalf("no se pudo crear la carpeta pdf: %v", err)
	}

	// Leer Markdown
	md, err := os.ReadFile(mdFile)
	if err != nil {
		t.Fatalf("no se pudo leer el archivo Markdown: %v", err)
	}

	// Convertir Markdown a HTML
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
	buf.WriteString("<!doctype html><html><head><meta charset=\"utf-8\">")
	buf.WriteString("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">")
	buf.WriteString("<style>")
	buf.WriteString(BaseCSS())
	buf.WriteString("</style>")
	buf.WriteString("</head><body><main class=\"container\">")
	buf.Write(body)
	buf.WriteString("</main></body></html>")
	htmlContent := buf.String()

	// Guardar HTML
	if err := os.WriteFile(htmlFile, []byte(htmlContent), 0o644); err != nil {
		t.Fatalf("no se pudo escribir HTML: %v", err)
	}

	// Convertir HTML a PDF usando chromedp
	abs, err := filepath.Abs(htmlFile)
	if err != nil {
		t.Fatalf("no se pudo obtener ruta absoluta de HTML: %v", err)
	}
	url := "file://" + abs

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), chromedp.DefaultExecAllocatorOptions[:]...)
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
		t.Fatalf("error generando PDF: %v", err)
	}

	// Guardar PDF
	if err := os.WriteFile(pdfFile, pdfBuf, 0o644); err != nil {
		t.Fatalf("no se pudo escribir PDF: %v", err)
	}

	// Verificar que el PDF fue creado
	if _, err := os.Stat(pdfFile); err != nil {
		t.Fatalf("el PDF no fue creado: %v", err)
	}
}

func BaseCSS() string {
	return `
:root {
	--font: system-ui, -apple-system, Segoe UI, Roboto, Ubuntu, Cantarell, "Helvetica Neue", Arial, "Noto Sans", "Liberation Sans", sans-serif;
	max-width: 100dvh;
}
body {
	font-family: var(--font);
	color: #222;
	line-height: 1.4;
	background: #fafbfc;
}
.main, .container {
	margin: 1rem auto;
	padding: 0 1rem;
	max-width: 900px;
}
h1, h2, h3 {
	line-height: 1.25;
	margin-top: 1.2rem;
	margin-bottom: .2rem;
}
pre {
	padding: 1rem;
	overflow: auto;
	background: #f6f8fa;
	border-radius: 8px;
	border: 2px solid #e1e4e8;
	box-shadow: 0 2px 8px rgba(0,0,0,0.04);
	margin: 1.2em 0;
}
pre code {
	background: none;
	padding: 0;
	border-radius: 0;
	border: none;
	box-shadow: none;
}
code {
	background: #f6f8fa;
	padding: .2rem .4rem;
	border-radius: 4px;
	border: 1px solid #e1e4e8;
}
table {
	border-collapse: collapse;
	width: 100%;
	margin: 1.5rem 0;
	background: #fff;
	box-shadow: 0 2px 8px rgba(0,0,0,0.03);
}
th, td {
	border: 1px solid #bbb;
	padding: .7rem;
	text-align: left;
}
th {
	background: #f3f4f6;
	font-weight: 600;
}
tr:nth-child(even) {
	background: #f8f9fa;
}
tr:hover {
	background: #eef2fb;
}
blockquote {
	border-left: 4px solid #6c63ff;
	padding-left: 1rem;
	color: #555;
	background: #f6f8fa;
	margin: 1rem 0;
}
img {
	max-width: 100%;
	height: auto;
	display: block;
	margin: 1rem auto;
	box-shadow: 0 2px 8px rgba(0,0,0,0.05);
}
a {
	color: #0b5fff;
	text-decoration: none;
	border-bottom: 1px dotted #0b5fff;
}
a:hover {
	text-decoration: underline;
	background: #eaf4ff;
}
ul, ol {
	padding-left: 1.5rem;
	margin-bottom: 1rem;
}
li {
	margin-bottom: .4rem;
}
hr {
	border: none;
	border-top: 1px solid #e1e4e8;
	margin: 2rem 0;
}
/* Iconos check y cruz para tablas */
.check {
	color: #6c63ff;
	font-size: 1.2em;
}
.cross {
	color: #e74c3c;
	font-size: 1.2em;
}
`
}

func BaseCSS2() string {
	return `
:root { --font: system-ui, -apple-system, Segoe UI, Roboto, Ubuntu, Cantarell, "Helvetica Neue", Arial, "Noto Sans", "Liberation Sans", sans-serif; max-width: 800px; }
body { font-family: var(--font); color: #111; line-height: 1.6; }
.main, .container { margin: 2rem auto; padding: 0 1rem; max-width: 900px; }
h1,h2,h3 { line-height: 1.25; margin-top: 1.8rem; margin-bottom: .8rem; }
pre, code { font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, "Liberation Mono", monospace; }
pre { padding: 1rem; overflow:auto; background: #f6f8fa; border-radius: 8px; }
code { background: #f6f8fa; padding: .2rem .4rem; border-radius: 4px; }
table { border-collapse: collapse; width: 100%; }
th, td { border: 1px solid #ddd; padding: .6rem; }
blockquote { border-left: 4px solid #ddd; padding-left: 1rem; color: #555; }
img { max-width: 100%; height: auto; }
a { color: #0b5fff; text-decoration: none; }
a:hover { text-decoration: underline; }
ul, ol { padding-left: 1.2rem; }
`
}

func TestHtmlToPdfByUrl(t *testing.T) {
	// Implementar prueba para conversión de HTML a PDF a través de URL
	pdfDir := "../pdf"
	pdfFile := filepath.Join(pdfDir, "url_test.pdf")

	// Crear carpeta pdf si no existe
	if err := os.MkdirAll(pdfDir, 0o755); err != nil {
		t.Fatalf("no se pudo crear la carpeta pdf: %v", err)
	}

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	url := "https://www.canva.com/design/DAGxmJzDvvY/v2XaOr5NXTVL0gCdRZXcRw/view?utm_content=DAGxmJzDvvY&utm_campaign=designshare&utm_medium=link2&utm_source=uniquelinks&utlId=hfc20aa30df" // ejemplo: archivo HTML/texto público
	pdfBuf, err := HtmlToPdf(ctx, url)
	if err != nil {
		t.Fatalf("error generando PDF: %v", err)
	}

	// Guardar PDF
	if err := os.WriteFile(pdfFile, pdfBuf, 0o644); err != nil {
		t.Fatalf("no se pudo escribir PDF: %v", err)
	}

	// Verificar que el PDF fue creado
	if _, err := os.Stat(pdfFile); err != nil {
		t.Fatalf("el PDF no fue creado: %v", err)
	}
}

func HtmlToPdf(ctx context.Context, url string) ([]byte, error) {
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
	err := chromedp.Run(ctx, tasks)
	return pdfBuf, err
}

func TestHtmlToPdfByUrlHorizontal(t *testing.T) {
	// Implementar prueba para conversión de HTML a PDF a través de URL
	pdfDir := "../pdf"
	pdfFile := filepath.Join(pdfDir, "url_test.pdf")

	// Crear carpeta pdf si no existe
	if err := os.MkdirAll(pdfDir, 0o755); err != nil {
		t.Fatalf("no se pudo crear la carpeta pdf: %v", err)
	}

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	url := "https://online.fliphtml5.com/gnlcv/ccci/files/large/aa8743296e3ee58c31ea5d0492d6105b.webp?1679498399&1679498399" // ejemplo: archivo HTML/texto público
	pdfBuf, err := HtmlToPdfHorizontal(ctx, url)
	if err != nil {
		t.Fatalf("error generando PDF: %v", err)
	}

	// Guardar PDF
	if err := os.WriteFile(pdfFile, pdfBuf, 0o644); err != nil {
		t.Fatalf("no se pudo escribir PDF: %v", err)
	}

	// Verificar que el PDF fue creado
	if _, err := os.Stat(pdfFile); err != nil {
		t.Fatalf("el PDF no fue creado: %v", err)
	}
}

func HtmlToPdfHorizontal(ctx context.Context, url string) ([]byte, error) {
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
				WithPaperWidth(11.69).
				WithPaperHeight(8.27).
				Do(ctx)
			return err
		}),
	}
	err := chromedp.Run(ctx, tasks)
	return pdfBuf, err
}

func TestImageUrlToPdf(t *testing.T) {
	pdfDir := "../pdf"
	pdfFile := filepath.Join(pdfDir, "certificado_img3.pdf")
	htmlFile := filepath.Join(pdfDir, "certificado_img3.html")

	// Crear carpeta pdf si no existe
	if err := os.MkdirAll(pdfDir, 0o755); err != nil {
		t.Fatalf("no se pudo crear la carpeta pdf: %v", err)
	}

	// URL de la imagen del certificado
	imgUrl := "https://cdn.filestackcontent.com/snHvx1cTaCcT67o2LvGY?policy=eyJjYWxsIjpbInJlYWQiXSwiZXhwaXJ5IjoxNzU2NjA0MzYzLCJwYXRoIjoiLyJ9&signature=9af5b722e9388b3ade422cb39d2323d68c76c1ebbd16e015e66836cea808b6cf" // reemplaza por tu URL

	// Crear HTML con la imagen centrada y fondo blanco
	htmlContent := `<!doctype html>
<html><head><meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<style>
body { background: #fff; margin: 0; padding: 0; }
.img-cert { display: block; margin: 0 auto; max-width: 1000px; width: 100%; height: auto; box-shadow: 0 2px 8px rgba(0,0,0,0.08); }
</style>
</head><body>
<img src="` + imgUrl + `" class="img-cert" />
</body></html>`

	// Guardar HTML
	if err := os.WriteFile(htmlFile, []byte(htmlContent), 0o644); err != nil {
		t.Fatalf("no se pudo escribir HTML: %v", err)
	}

	// Convertir HTML a PDF usando chromedp
	abs, err := filepath.Abs(htmlFile)
	if err != nil {
		t.Fatalf("no se pudo obtener ruta absoluta de HTML: %v", err)
	}
	url := "file://" + abs

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var pdfBuf []byte
	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady("img", chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfBuf, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithMarginTop(0.2).
				WithMarginBottom(0.2).
				WithMarginLeft(0.2).
				WithMarginRight(0.2).
				WithPaperWidth(11.69).
				WithPaperHeight(8.27).
				Do(ctx)
			return err
		}),
	}
	if err := chromedp.Run(ctx, tasks); err != nil {
		t.Fatalf("error generando PDF: %v", err)
	}

	// Guardar PDF
	if err := os.WriteFile(pdfFile, pdfBuf, 0o644); err != nil {
		t.Fatalf("no se pudo escribir PDF: %v", err)
	}

	// Verificar que el PDF fue creado
	if _, err := os.Stat(pdfFile); err != nil {
		t.Fatalf("el PDF no fue creado: %v", err)
	}
}

func TestLocalImageToPdf(t *testing.T) {
	pdfDir := "../pdf"
	imgDir := "../img"
	imgFile := filepath.Join(imgDir, "certificate-391353386.jpg")
	pdfFile := filepath.Join(pdfDir, "certificate_local.pdf")
	htmlFile := filepath.Join(pdfDir, "certificate_local.html")

	// Crear carpeta pdf si no existe
	if err := os.MkdirAll(pdfDir, 0o755); err != nil {
		t.Fatalf("no se pudo crear la carpeta pdf: %v", err)
	}

	// Formato: "horizontal" o "vertical"
	formato := "horizontal" // Cambia a "vertical" si lo deseas
	var paperWidth, paperHeight float64
	if formato == "horizontal" {
		paperWidth = 11.69 // A4 horizontal
		paperHeight = 8.27
	} else {
		paperWidth = 8.27 // A4 vertical
		paperHeight = 11.69
	}

	// Crear HTML con la imagen centrada y fondo blanco
	htmlContent := `<!doctype html>
<html><head><meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<style>
body { background: #fff; margin: 0; padding: 0; }
.img-cert { display: block; margin: 0 auto; max-width: 1000px; width: 100%; height: auto; box-shadow: 0 2px 8px rgba(0,0,0,0.08); }
</style>
</head><body>
<img src="` + imgFile + `" class="img-cert" />
</body></html>`

	// Guardar HTML
	if err := os.WriteFile(htmlFile, []byte(htmlContent), 0o644); err != nil {
		t.Fatalf("no se pudo escribir HTML: %v", err)
	}

	// Convertir HTML a PDF usando chromedp
	abs, err := filepath.Abs(htmlFile)
	if err != nil {
		t.Fatalf("no se pudo obtener ruta absoluta de HTML: %v", err)
	}
	url := "file://" + abs

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var pdfBuf []byte
	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady("img", chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfBuf, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithMarginTop(0.2).
				WithMarginBottom(0.2).
				WithMarginLeft(0.2).
				WithMarginRight(0.2).
				WithPaperWidth(paperWidth).
				WithPaperHeight(paperHeight).
				Do(ctx)
			return err
		}),
	}
	if err := chromedp.Run(ctx, tasks); err != nil {
		t.Fatalf("error generando PDF: %v", err)
	}

	// Guardar PDF
	if err := os.WriteFile(pdfFile, pdfBuf, 0o644); err != nil {
		t.Fatalf("no se pudo escribir PDF: %v", err)
	}

	// Verificar que el PDF fue creado
	if _, err := os.Stat(pdfFile); err != nil {
		t.Fatalf("el PDF no fue creado: %v", err)
	}
}

func TestHtmlUrlTicketAreaToPdf(t *testing.T) {
	pdfDir := "../pdf"
	imgFile := filepath.Join(pdfDir, "ticket_area13.png")
	htmlFile := filepath.Join(pdfDir, "ticket_area13.html")
	pdfFile := filepath.Join(pdfDir, "ticket_area13.pdf")

	// Crear carpeta pdf si no existe
	if err := os.MkdirAll(pdfDir, 0o755); err != nil {
		t.Fatalf("no se pudo crear la carpeta pdf: %v", err)
	}

	url := "https://crt-trs-atlbqiwvnq-tl.a.run.app/v1/ticket/print/0892c546-db60-41c6-b5f2-fcc8ffec283a.html"

	// Coordenadas y dimensiones para el área del ticket (ajusta según lo necesites)
	x := 8.0        // posición X inicial
	y := 14.0       // posición Y inicial
	width := 358.0  // ancho del recorte (ejemplo: 350 px)
	height := 618.0 // alto del recorte (ejemplo: 600 px)

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx,
		// Ajustar el viewport antes de navegar
		emulation.SetDeviceMetricsOverride(int64(width), int64(height+5), 1.0, false).WithScreenOrientation(&emulation.ScreenOrientation{Type: emulation.OrientationTypePortraitPrimary, Angle: 0}),
		chromedp.Navigate(url),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, err = page.CaptureScreenshot().
				WithFormat(page.CaptureScreenshotFormatPng).
				WithClip(&page.Viewport{
					X:      x,
					Y:      y,
					Width:  width,
					Height: height,
					Scale:  1,
				}).
				Do(ctx)
			return err
		}),
	)
	if err != nil {
		t.Fatalf("error capturando imagen: %v", err)
	}

	// Guardar imagen
	if err := os.WriteFile(imgFile, buf, 0o644); err != nil {
		t.Fatalf("no se pudo escribir imagen: %v", err)
	}

	// Crear HTML con la imagen centrada y fondo blanco
	htmlContent := `<!doctype html>
<html><head><meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
</head><body>
<img src="` + imgFile + `" />
</body></html>`

	// Guardar HTML
	if err := os.WriteFile(htmlFile, []byte(htmlContent), 0o644); err != nil {
		t.Fatalf("no se pudo escribir HTML: %v", err)
	}

	// Convertir HTML a PDF usando chromedp
	abs, err := filepath.Abs(htmlFile)
	if err != nil {
		t.Fatalf("no se pudo obtener ruta absoluta de HTML: %v", err)
	}
	fileUrl := "file://" + abs

	var pdfBuf []byte
	ctx2, cancel2 := chromedp.NewContext(context.Background())
	defer cancel2()
	err = chromedp.Run(ctx2,
		chromedp.Navigate(fileUrl),
		chromedp.WaitReady("img", chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfBuf, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithMarginTop(0).
				WithMarginBottom(0).
				WithMarginLeft(0).
				WithMarginRight(0).
				WithPaperWidth(width / 96).
				WithPaperHeight(height / 96).
				Do(ctx)
			return err
		}),
	)
	if err != nil {
		t.Fatalf("error generando PDF: %v", err)
	}

	// Guardar PDF
	if err := os.WriteFile(pdfFile, pdfBuf, 0o644); err != nil {
		t.Fatalf("no se pudo escribir PDF: %v", err)
	}

	// Verificar que el PDF fue creado
	if _, err := os.Stat(pdfFile); err != nil {
		t.Fatalf("el PDF no fue creado: %v", err)
	}
}
