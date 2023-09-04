package bing_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	translator "github.com/souk4711/honyakusha/internal/translator/bing"
)

func TestBuildReqFormDataSource(t *testing.T) {
	assert.Equal(t, "auto-detect", translator.BuildReqFormDataSource(""))
	assert.Equal(t, "en", translator.BuildReqFormDataSource("en-US"))
	assert.Equal(t, "ja", translator.BuildReqFormDataSource("ja"))
	assert.Equal(t, "zh-Hans", translator.BuildReqFormDataSource("zh-CN"))
}

func TestBuildReqFormDataTarget(t *testing.T) {
	assert.Equal(t, "en", translator.BuildReqFormDataTarget("en-US"))
	assert.Equal(t, "ja", translator.BuildReqFormDataSource("ja"))
	assert.Equal(t, "zh-Hant", translator.BuildReqFormDataTarget("zh-TW"))
}
