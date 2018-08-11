package env

import "github.com/sutd-statnlp/service-ladrawex/constant"

var (
	// EnableProdMode enables production mode.
	EnableProdMode = Bool(constant.EnableProdMode, false)

	// ServerAddress is the server address for running web server.
	ServerAddress = String(constant.ServerAddress, ":9000")
)
