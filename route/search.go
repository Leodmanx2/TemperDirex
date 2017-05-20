package route

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"path/filepath"
	"strings"
)

func search(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Grab a  list of all files in the serving directory
	pattern := filepath.Join(configuration.ServeDirectory, "/*")
	files, err := filepath.Glob(pattern)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	// Find files that match all terms in the client's query
	query := req.FormValue("query")
	terms := strings.Fields(query)
	var fileMatches []string
	for _, file := range files {
		file = filepath.Base(file)
		var termMatches int
		for _, term := range terms {
			if strings.Contains(file, term) {
				termMatches++
			}
			if termMatches == len(terms) {
				fileMatches = append(fileMatches, file)
			}
		}
	}

	// Execute template using the results
	data := struct {
		Query string
		Files []string
	}{
		query,
		fileMatches,
	}
	err = templates.ExecuteTemplate(res, "index.html", &data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
