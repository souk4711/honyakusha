package plain

import (
	"strings"

	"github.com/souk4711/honyakusha/internal/res"
)

func Format(res res.Res) string {
	if res.Code != 0 {
		return res.Error
	}

	r := ""
	for i, translator := range res.Translators {
		if i != 0 {
			r = r + "\n"
		}

		r = r + translator.Translator.Name + ":\n"
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
	for _, ele := range strings.Split(strings.TrimSpace(str), "\n") {
		arr = append(arr, "    "+ele)
	}
	return strings.Join(arr, "\n")
}
