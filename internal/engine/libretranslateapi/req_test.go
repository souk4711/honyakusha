package libretranslateapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	api "github.com/souk4711/honyakusha/internal/engine/libretranslateapi"
)

func TestBuildReqBodySource(t *testing.T) {
	assert.Equal(t, "auto", api.BuildReqBodySource(""))
	assert.Equal(t, "en", api.BuildReqBodySource("en-US"))
	assert.Equal(t, "ja", api.BuildReqBodySource("ja"))
}

func TestBuildReqBodyTarget(t *testing.T) {
	assert.Equal(t, "en", api.BuildReqBodyTarget("en-US"))
	assert.Equal(t, "ja", api.BuildReqBodySource("ja"))
}
