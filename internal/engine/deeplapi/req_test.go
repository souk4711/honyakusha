package deeplapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	engine "github.com/souk4711/honyakusha/internal/engine/deeplapi"
)

func TestBuildReqBodySource(t *testing.T) {
	assert.Equal(t, "", engine.BuildReqBodySource(""))
	assert.Equal(t, "EN", engine.BuildReqBodySource("en-US"))
	assert.Equal(t, "JA", engine.BuildReqBodySource("ja"))
}

func TestBuildReqBodyTarget(t *testing.T) {
	assert.Equal(t, "EN-US", engine.BuildReqBodyTarget("en-US"))
	assert.Equal(t, "JA", engine.BuildReqBodySource("ja"))
}
