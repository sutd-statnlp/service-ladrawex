package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/web/api"
)

func TestGroup(test *testing.T) {
	group := api.Group()
	assert.NotNil(test, group)
}
