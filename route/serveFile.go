package route

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"path/filepath"
)

func serveFile(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	path := filepath.Join(configuration.ServeDirectory, params.ByName("file"))
	http.ServeFile(res, req, path)
}
