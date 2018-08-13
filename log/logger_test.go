package log_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sutd-statnlp/service-ladrawex/constant"
	"github.com/sutd-statnlp/service-ladrawex/log"
)

func TestOpenFile(test *testing.T) {
	assert.True(test, log.OpenFile(constant.LogFilePath))
	assert.False(test, log.OpenFile(".."))
}

func TestCloseFile(test *testing.T) {
	assert.False(test, log.CloseFile())
	log.OpenFile(constant.LogFilePath)
	assert.True(test, log.CloseFile())
}
