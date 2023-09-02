package deeplapi

import (
	"strings"

	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/lang"
)

type Req struct {
	Text       []string `json:"text"`
	SourceLang string   `json:"source_lang"`
	TargetLang string   `json:"target_lang"`
}

func buildReqURL(conf conf.ConfTranslator) string {
	if conf.URI == "" {
		return "https://api-free.deepl.com/v2/translate"
	} else {
		return conf.URI
	}
}

func buildReqBody(text string, from string, to string) Req {
	from = buildReqBodySourceLang(from)
	to = buildReqBodyTargetLang(to)
	if from == "" { // autodetect
		return Req{Text: []string{text}, TargetLang: to}
	} else {
		return Req{Text: []string{text}, SourceLang: from, TargetLang: to}
	}
}

func buildReqBodySourceLang(from string) string {
	frLang := lang.Query(from)
	code := frLang.Code
	switch code {
	}
	if code == "" {
		return ""
	} else {
		return strings.ToUpper(code[0:2])
	}
}

func buildReqBodyTargetLang(to string) string {
	toLang := lang.Query(to)
	code := toLang.Code
	return strings.ToUpper(code)
}
