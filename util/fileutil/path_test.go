package fileutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/util/fileutil"
)

func TestAbsPath(test *testing.T) {
	path := *fileutil.AbsPath("./")
	assert.NotNil(test, path)
	assert.True(test, len(path) > 0)
}
