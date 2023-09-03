package libretranslateapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	translator "github.com/souk4711/honyakusha/internal/translator/libretranslateapi"
)

func TestBuildReqBodySource(t *testing.T) {
	assert.Equal(t, "auto", translator.BuildReqBodySource(""))
	assert.Equal(t, "en", translator.BuildReqBodySource("en-US"))
	assert.Equal(t, "ja", translator.BuildReqBodySource("ja"))
}

func TestBuildReqBodyTarget(t *testing.T) {
	assert.Equal(t, "en", translator.BuildReqBodyTarget("en-US"))
	assert.Equal(t, "ja", translator.BuildReqBodySource("ja"))
}
