package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Configuration encompasses a set of variables that
// should be set by the user and loaded at run-time.
type Configuration struct {
	ServeDirectory string // ServeDirectory is the path to serve files from
	Cert           string // Cert is the path to an x509 certificate
	Key            string // Cert is the path to an x509 private key
	TLS            string // TLS is one of 'plain', 'both', or 'tls'
}

// New parses the input filename into a new Configuration.
func New(filename string) (conf *Configuration, err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = json.Unmarshal(file, &conf)
	if err != nil {
		return
	}

	// Unset directory will result in serving files from root.
	// We don't want that happening unintentionally.
	if conf.ServeDirectory == "" {
		err = errors.New("config option ServeDirectory must be set")
	}

	// Set default values
	if conf.Cert == "" {
		conf.Cert = "tls/cert.pem"
	}
	if conf.Key == "" {
		conf.Key = "tls/key.pem"
	}
	if conf.TLS == "" {
		conf.TLS = "plain"
	}
	return
}
