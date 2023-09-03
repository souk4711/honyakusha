package libretranslateapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	engine "github.com/souk4711/honyakusha/internal/engine/libretranslateapi"
)

func TestBuildReqBodySource(t *testing.T) {
	assert.Equal(t, "auto", engine.BuildReqBodySource(""))
	assert.Equal(t, "en", engine.BuildReqBodySource("en-US"))
	assert.Equal(t, "ja", engine.BuildReqBodySource("ja"))
}

func TestBuildReqBodyTarget(t *testing.T) {
	assert.Equal(t, "en", engine.BuildReqBodyTarget("en-US"))
	assert.Equal(t, "ja", engine.BuildReqBodySource("ja"))
}
