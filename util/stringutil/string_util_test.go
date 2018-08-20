package stringutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/util/stringutil"
)

func TestConcat(test *testing.T) {
	result := stringutil.Concat("AA", "BB")
	assert.NotNil(test, result)
	assert.True(test, len(result) > 0)
}
