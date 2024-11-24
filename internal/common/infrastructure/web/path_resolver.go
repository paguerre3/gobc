package web

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/paguerre3/blockchain/configs"
)

var (
	config = configs.Instance()
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
			index := strings.Index(wdir, config.CmdDir())
			if index == -1 {
				index = strings.Index(wdir, config.InternalDir())
				if index == -1 {
					log.Errorf("cannot find %s or %s in %s", config.CmdDir(), config.InternalDir(), wdir)
					return
				}
			}
			rootDir := wdir[:index]
			path = filepath.Join(rootDir, pathOfDomain)
		})
		return path
	}
}
