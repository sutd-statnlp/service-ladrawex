package server

// WebServer is the http server interface.
type WebServer interface {
	Config() bool
	Run()
}
