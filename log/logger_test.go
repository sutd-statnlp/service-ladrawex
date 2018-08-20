package log_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sutd-statnlp/service-ladrawex/log"
)

func TestOpenFile(test *testing.T) {
	assert.True(test, log.OpenFile())
}

func TestCloseFile(test *testing.T) {
	log.OpenFile()
	assert.True(test, log.CloseFile())
}
