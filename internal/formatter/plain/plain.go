package plain

import (
	"strings"

	"github.com/souk4711/honyakusha/internal/helpers"
	"github.com/souk4711/honyakusha/internal/res"
)

func Format(res res.Res) string {
	if res.Code != 0 {
		return res.Error + "\n"
	}

	r := ""
	r = r + "Raw:\n\n"
	r = r + indent(res.Text) + "\n"

	for _, translator := range res.Translators {
		r = r + "\n"
		r = r + translator.Translator.Name + ":\n\n"
		if translator.Code != 0 {
			r = r + indent(translator.Error) + "\n"
		} else {
			r = r + indent(translator.TranslatedText) + "\n"
		}
	}
	return r
}

func indent(str string) string {
	var arr []string
	for _, ele := range helpers.SplitTextIntoArray(str) {
		arr = append(arr, "    "+ele)
	}
	return strings.Join(arr, "\n")
}
