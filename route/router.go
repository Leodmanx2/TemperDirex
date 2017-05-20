package route

import (
	"bitbucket.org/leodmanx2/TemperDirex/config"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"xi2.org/x/httpgzip"
)

var templates *template.Template
var configuration *config.Configuration

// ReadTemplates parses all required template files from
// dir, storing the parsed data locally. Subsequent calls
// will overwrite this data.
func ReadTemplates(dir string) (err error) {
	templates, err = template.ParseGlob(dir + "/*")
	return
}

// New creates a new router with associated handlers
func New(conf *config.Configuration) http.Handler {
	configuration = conf
	router := httprouter.New()
	router.GET("/search", search)
	router.GET("/file/:file", serveFile)
	router.GET("/", index)
	gzipRouter, err := httpgzip.NewHandlerLevel(router, httpgzip.DefaultContentTypes, 5)
	if err != nil {
		log.Print(err)
		return router
	}
	return gzipRouter
}
