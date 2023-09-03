package deeplapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	api "github.com/souk4711/honyakusha/internal/engine/deeplapi"
)

func TestBuildReqBodySource(t *testing.T) {
	assert.Equal(t, "", api.BuildReqBodySource(""))
	assert.Equal(t, "EN", api.BuildReqBodySource("en-US"))
	assert.Equal(t, "JA", api.BuildReqBodySource("ja"))
}

func TestBuildReqBodyTarget(t *testing.T) {
	assert.Equal(t, "EN-US", api.BuildReqBodyTarget("en-US"))
	assert.Equal(t, "JA", api.BuildReqBodySource("ja"))
}
