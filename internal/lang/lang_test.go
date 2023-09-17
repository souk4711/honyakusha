package lang_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/souk4711/honyakusha/internal/lang"
)

func TestCode_639_1(t *testing.T) {
	r := lang.Query("fr")
	assert.Equal(t, "fr", r.Code_639_1())

	r = lang.Query("fr-CA")
	assert.Equal(t, "fr", r.Code_639_1())

	r = lang.Query("mhr")
	assert.Equal(t, "mhr", r.Code_639_1())

	r = lang.Query("mni-Mtei")
	assert.Equal(t, "mni", r.Code_639_1())

	r = lang.Query("")
	assert.Equal(t, "", r.Code_639_1())
}

func TestAutoDetect(t *testing.T) {
	os.Setenv("LANGUAGE", "en_US")
	r := lang.AutoDetect()
	assert.Equal(t, "en-US", r)

	os.Setenv("LANGUAGE", "zh_CN")
	r = lang.AutoDetect()
	assert.Equal(t, "zh-CN", r)

	os.Setenv("LANGUAGE", "ja_JP")
	r = lang.AutoDetect()
	assert.Equal(t, "ja", r)
}

func TestQuery(t *testing.T) {
	r := lang.Query("en-US")
	assert.Equal(t, "en-US", r.Code)
	assert.Equal(t, "English (American)", r.Name)
	assert.Equal(t, "English (American)", r.LocaleName)

	r = lang.Query("")
	assert.Equal(t, "", r.Code)
	assert.Equal(t, "", r.Name)
	assert.Equal(t, "", r.LocaleName)
}
