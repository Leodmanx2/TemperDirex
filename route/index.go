package route

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"path/filepath"
)

func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// Grab a  list of all files in the serving directory
	pattern := filepath.Join(configuration.ServeDirectory, "/*")
	files, err := filepath.Glob(pattern)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	// Strip away the parent path
	for i, file := range files {
		files[i] = filepath.Base(file)
	}

	// Execute template using the results
	data := struct {
		Query string
		Files []string
	}{
		"",
		files,
	}
	err = templates.ExecuteTemplate(res, "index.html", &data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
