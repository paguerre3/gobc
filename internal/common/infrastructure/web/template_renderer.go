package web

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const (
	CMD_DIR      = "/cmd/"
	INTERNAL_DIR = "/internal/"
)

var (
	templateDir string
	mutex       sync.Mutex
)

type PageData struct {
	Year int
}

type TemplateRenderer interface {
	Render(w io.Writer, name string, data any, c echo.Context) error
}

// TemplateRenderer is a custom Echo template renderer
type templateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer(templatesPathOfDomain string) TemplateRenderer {
	templateDir = resolveTemplateDir(templatesPathOfDomain)
	templatesPath := fmt.Sprintf("%s/*.html", templateDir)
	return &templateRenderer{
		templates: template.Must(template.ParseGlob(templatesPath)),
	}
}

func (t *templateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func resolveTemplateDir(templatesPathOfDomain string) string {
	if len(templateDir) == 0 {
		mutex.Lock()
		defer mutex.Unlock()
		if len(templateDir) == 0 {
			// Get the directory of the executable
			wdir, err := os.Getwd()
			if err != nil {
				log.Error(err)
				return ""
			}
			index := strings.Index(wdir, CMD_DIR)
			if index == -1 {
				index = strings.Index(wdir, INTERNAL_DIR)
				if index == -1 {
					log.Errorf("cannot find %s or %s in %s", CMD_DIR, INTERNAL_DIR, wdir)
					return ""
				}
			}
			rootDir := wdir[:index]
			templateDir = filepath.Join(rootDir, templatesPathOfDomain)
		}
	}
	return templateDir
}
