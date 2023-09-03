package lang_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/souk4711/honyakusha/internal/lang"
)

func TestMacro(t *testing.T) {
	r := lang.Query("fr")
	assert.Equal(t, "fr", r.Macro())

	r = lang.Query("fr-CA")
	assert.Equal(t, "fr", r.Macro())

	r = lang.Query("mhr")
	assert.Equal(t, "mhr", r.Macro())

	r = lang.Query("mni-Mtei")
	assert.Equal(t, "mni", r.Macro())

	r = lang.Query("")
	assert.Equal(t, "", r.Macro())
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
