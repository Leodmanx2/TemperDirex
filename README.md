TemperDirex
===========

TemperDirex is a small search-enabled file server written in Go.

To build it simply run `go get bitbucket.org/leodmanx2/TemperDirex`.
The binary will be in $GOPATH/bin.

The program requires `template/` and `conf.json` to be in the working
directory. You can run it from
`$GOPATH/src/bitbucket.org/leodmanx2/TemperDirex/` or copy these out.

Configuration
-------------

Options may be set in `conf.json`. Each option is case-sensitive.

*   ServeDirectory
    *   Mandatory
    *   The directory to serve files from
*   Cert
    *   Optional
    *   Default: tls/cert.pem
    *   The certificate TemperDirex will use if TLS is enabled
*   Key
    *   Optional
    *   Default: tls/key.pem
    *   The key TemperDirex will use if TLS is enabled
*   TLS
    *   Optional
    *   Default: plain
    *   Determines whether to serve over HTTP, HTTPS, or both.
    *   Possible values: `plain`, `tls`, `both`
