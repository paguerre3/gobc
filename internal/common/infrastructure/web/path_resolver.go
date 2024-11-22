package web

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/labstack/gommon/log"
)

const (
	CMD_DIR      = "/cmd/"
	INTERNAL_DIR = "/internal/"
)

func NewPathResolver() func(string) string {
	var (
		path string
		once sync.Once
	)

	return func(pathOfDomain string) string {
		once.Do(func() {
			// Get the directory of the executable
			wdir, err := os.Getwd()
			if err != nil {
				log.Error(err)
				return
			}
			index := strings.Index(wdir, CMD_DIR)
			if index == -1 {
				index = strings.Index(wdir, INTERNAL_DIR)
				if index == -1 {
					log.Errorf("cannot find %s or %s in %s", CMD_DIR, INTERNAL_DIR, wdir)
					return
				}
			}
			rootDir := wdir[:index]
			path = filepath.Join(rootDir, pathOfDomain)
		})
		return path
	}
}
