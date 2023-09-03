package google_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	translator "github.com/souk4711/honyakusha/internal/translator/google"
)

func TestBuildReqQueryParamsSource(t *testing.T) {
	assert.Equal(t, "auto", translator.BuildReqQueryParamsSource(""))
	assert.Equal(t, "en", translator.BuildReqQueryParamsSource("en-US"))
	assert.Equal(t, "ja", translator.BuildReqQueryParamsSource("ja"))
}

func TestBuildReqQueryParamsTarget(t *testing.T) {
	assert.Equal(t, "en", translator.BuildReqQueryParamsTarget("en-US"))
	assert.Equal(t, "ja", translator.BuildReqQueryParamsSource("ja"))
}
