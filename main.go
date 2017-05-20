package main

import (
	"bitbucket.org/leodmanx2/TemperDirex/config"
	"bitbucket.org/leodmanx2/TemperDirex/route"
	"log"
	"net/http"
	"strings"
)

var conf *config.Configuration

// redirect is an http.Handler that redirects requests to their HTTPS equivalent
func redirect(res http.ResponseWriter, req *http.Request) {
	host := strings.Split(req.Host, ":")[0]
	target := "https://" + host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	http.Redirect(res, req, target, http.StatusPermanentRedirect)
}

func main() {
	// Utility function for top-level error handling
	var err error
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	conf, err = config.New("conf.json")
	check(err)

	err = route.ReadTemplates("template/")
	check(err)

	// Start serving HTTPS if enabled
	if conf.TLS == "tls" || conf.TLS == "both" {
		go func() {
			log.Println("Listening on port 443 (HTTPS).")
			err := http.ListenAndServeTLS(":https", conf.Cert, conf.Key, route.New(conf))
			check(err)
		}()
	}

	if conf.TLS == "tls" {
		// If only HTTPS is enabled, redirect HTTP to HTTPS
		log.Println("Redirecting port 80 (HTTP) to port 443 (HTTPS)")
		err = http.ListenAndServe(":http", http.HandlerFunc(redirect))
		check(err)
	} else {
		// Otherwise serve plain HTTP
		log.Println("Listening on port 80 (HTTP).")
		err = http.ListenAndServe(":http", route.New(conf))
		check(err)
	}
}
