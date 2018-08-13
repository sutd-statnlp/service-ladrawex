package testutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/util/testutil"
)

func TestCreateFakeServer(test *testing.T) {
	assert.NotNil(test, testutil.CreateFakeServer())
}
