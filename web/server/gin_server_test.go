package server_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/util/testutil"
)

func TestConfig(test *testing.T) {
	webServer := testutil.CreateFakeServer()
	assert.NotNil(test, webServer)
	assert.True(test, webServer.Config())
}
