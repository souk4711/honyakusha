package html_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	formatter "github.com/souk4711/honyakusha/internal/formatter/html"
	"github.com/souk4711/honyakusha/internal/res"
)

func TestFormat(t *testing.T) {
	r := res.NewResSuccess()
	output := formatter.Format(r)
	assert.Contains(t, output, "<html>")

	r.Translators = append(r.Translators, res.NewResTranslatorSuccess("HELLO, WORLD"))
	output = formatter.Format(r)
	assert.Contains(t, output, "HELLO, WORLD")
}
