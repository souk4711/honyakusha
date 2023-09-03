package deeplapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	translator "github.com/souk4711/honyakusha/internal/translator/deeplapi"
)

func TestBuildReqBodySource(t *testing.T) {
	assert.Equal(t, "", translator.BuildReqBodySource(""))
	assert.Equal(t, "EN", translator.BuildReqBodySource("en-US"))
	assert.Equal(t, "JA", translator.BuildReqBodySource("ja"))
}

func TestBuildReqBodyTarget(t *testing.T) {
	assert.Equal(t, "EN-US", translator.BuildReqBodyTarget("en-US"))
	assert.Equal(t, "JA", translator.BuildReqBodySource("ja"))
}
