package application

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	d "github.com/IberChambiU/MD2PDF/api/markdowntopdf/domain"
)

type MarkdownToPDFService interface {
	MarkdownToPDF(file *multipart.FileHeader) (*d.MarkDownToPDF, error)
}

type markdownToPDFService struct {
}

func NewMarkdownToPDFService() MarkdownToPDFService {
	return &markdownToPDFService{}
}

func (s *markdownToPDFService) MarkdownToPDF(file *multipart.FileHeader) (*d.MarkDownToPDF, error) {

	log.Printf("Markdown to PDF conversion")

	md2pdf, err := d.MakeMarkDownToPDF(file)
	if err != nil {
		log.Printf("Error: No se pudo crear el objeto MarkDownToPDF: %v", err)
		return nil, d.ErrFileNotConverted
	}

	src, err := md2pdf.File.Open()
	if err != nil {
		log.Printf("Error: No se pudo abrir el archivo: %v", err)
		return nil, d.ErrFileNotConverted
	}
	defer src.Close()

	// crear si no existe el directorio pdfdir
	if err := os.MkdirAll(d.PdfDir, os.ModePerm); err != nil {
		log.Printf("Error: creating ./pdf directory: %v", err)
		return nil, d.ErrFileNotConverted
	}

	// crear si no existe el directorio mddir
	if err := os.MkdirAll(d.MdDir, os.ModePerm); err != nil {
		log.Printf("Error: creating ./md directory: %v", err)
		return nil, d.ErrFileNotConverted
	}

	// crear si no existe el directorio htmldir
	if err := os.MkdirAll(d.HtmlDir, os.ModePerm); err != nil {
		log.Printf("Error: creating ./html directory: %v", err)
		return nil, d.ErrFileNotConverted
	}

	// guardar el file en ./md cambiando el nombre por un uuid
	dst, err := os.Create(md2pdf.MdFilePath)
	if err != nil {
		log.Printf("Error: No se pudo crear el archivo Markdown: %v", err)
		return nil, d.ErrFileNotConverted
	}
	defer os.Remove(md2pdf.MdFilePath)
	defer os.Remove(md2pdf.HtmlFilePath)
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		log.Printf("Error: No se pudo guardar el archivo Markdown: %v", err)
		return nil, d.ErrFileNotConverted
	}

	if err := d.ConvertMarkDownToPDF(md2pdf); err != nil {
		log.Printf("Error: No se pudo convertir el archivo Markdown a PDF: %v", err)
		return nil, d.ErrFileNotConverted
	}

	return md2pdf, nil
}

// CleanOldFilesFromPDFDir elimina archivos en la carpeta pdf que tengan más de maxAge horas de antigüedad
func CleanOldFilesFromPDFDir(maxAge time.Duration) {
	pdfDir := d.PdfDir
	files, err := os.ReadDir(pdfDir)
	if err != nil {
		log.Printf("Error leyendo el directorio pdf: %v", err)
		return
	}
	now := time.Now()
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		path := filepath.Join(pdfDir, file.Name())
		info, err := os.Stat(path)
		if err != nil {
			log.Printf("No se pudo obtener info de %s: %v", path, err)
			continue
		}
		if now.Sub(info.ModTime()) > maxAge {
			err := os.Remove(path)
			if err != nil {
				log.Printf("No se pudo eliminar %s: %v", path, err)
			} else {
				log.Printf("Archivo eliminado por antigüedad: %s", path)
			}
		}
	}
}
