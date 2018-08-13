package env

import (
	"github.com/sutd-statnlp/service-ladrawex/constant"
)

// EnableProdMode enables production mode.
func EnableProdMode() bool {
	return Bool(constant.EnableProdMode, false)
}

// ServerAddress is the server address for running web server.
func ServerAddress() string {
	return String(constant.ServerAddress, ":9000")
}
