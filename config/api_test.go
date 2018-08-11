package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sutd-statnlp/service-ladrawex/config"
	"github.com/sutd-statnlp/service-ladrawex/util"
)

func TestSetAPIPaths(test *testing.T) {
	apiGroups := config.API(util.CreateFakeRouter())
	assert.NotNil(test, apiGroups)
	assert.True(test, len(apiGroups) > 0)
}
