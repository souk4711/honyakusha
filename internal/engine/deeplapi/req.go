package deeplapi

import (
	"strings"

	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/lang"
)

type Req struct {
	Text   []string `json:"text"`
	Source string   `json:"source_lang"`
	Target string   `json:"target_lang"`
}

func buildReqURL(conf conf.ConfTranslator) string {
	if conf.URI == "" {
		return "https://api-free.deepl.com/v2/translate"
	} else {
		return conf.URI
	}
}

func buildReqBody(text string, from string, to string) Req {
	from = buildReqBodySource(from)
	to = buildReqBodyTarget(to)
	if from == "" {
		return Req{
			Text:   []string{text},
			Target: to,
		}
	} else {
		return Req{
			Text:   []string{text},
			Source: from,
			Target: to,
		}
	}
}

func buildReqBodySource(from string) string {
	frLang := lang.Query(from)
	code := frLang.Code
	if code == "" {
		return ""
	} else {
		return strings.ToUpper(frLang.Macro())
	}
}

func buildReqBodyTarget(to string) string {
	toLang := lang.Query(to)
	code := toLang.Code
	return strings.ToUpper(code)
}
