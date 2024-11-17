package api

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const (
	TEMPLATE_PATH = "internal/wallet/api/templates"
	CMD_DIR       = "/cmd/"
	INTERNAL_DIR  = "/internal/"
)

var (
	templateDir string
	mutex       sync.Mutex
)

type WalletHandler interface {
	Index(c echo.Context) error
}

type walletHandler struct {
}

func NewWalletHandler() WalletHandler {
	templateDir = TemplateDir()
	return &walletHandler{}
}

func TemplateDir() string {
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
			templateDir = filepath.Join(rootDir, TEMPLATE_PATH)
		}
	}
	return templateDir
}

func (w *walletHandler) Index(c echo.Context) error {
	return c.File(path.Join(templateDir, "index.html"))
}
