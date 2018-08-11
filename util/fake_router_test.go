package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/util"
)

func TestCreateFakeRouter(test *testing.T) {
	assert.NotNil(test, util.CreateFakeRouter())
}
