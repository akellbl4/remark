package templates

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/go-pkgz/lgr"
	"github.com/rakyll/statik/fs"
)

type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]os.FileInfo, error)
	Stat() (os.FileInfo, error)
}
type FileSystem interface {
	ReadFile(name string) ([]byte, error)
}

type osFS struct{}

var statikInstance http.FileSystem
var isStatik bool

func Init() FileSystem {
	var res FileSystem = osFS{}
	statikFS, err := fs.New()

	if err != nil {
		isStatik = false
		log.Printf("[INFO] templates will read from disk")
	} else {
		statikInstance = statikFS
		isStatik = true
		log.Printf("[INFO] templates will read from statik")
	}

	return res
}

func (osFS) ReadFile(path string) ([]byte, error) {
	if isStatik {
		return fs.ReadFile(statikInstance, path)
	}

	return ioutil.ReadFile(path)
}
