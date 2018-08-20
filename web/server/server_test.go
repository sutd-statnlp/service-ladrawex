package server_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/web/server"
)

func TestDefault(test *testing.T) {
	webServer := server.Default()
	assert.NotNil(test, webServer)
}

func TestNew(test *testing.T) {
	webServer := server.New()
	assert.NotNil(test, webServer)
}
