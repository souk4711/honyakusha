package deeplapi

import (
	"strings"

	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/lang"
)

var (
	data = map[string]string{
		"zh-CN": "ZH",
		"zh-TW": "ZH",
	}
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
		return conf.URI + "/v2/translate"
	}
}

func buildReqBody(text string, source string, target string, conf conf.ConfTranslator) Req {
	source = buildReqBodySource(source)
	target = buildReqBodyTarget(target)
	if source == "" {
		return Req{
			Text:   []string{text},
			Target: target,
		}
	} else {
		return Req{
			Text:   []string{text},
			Source: source,
			Target: target,
		}
	}
}

func buildReqBodySource(source string) string {
	l := lang.Query(source)
	if l.Code == "" {
		return ""
	} else {
		return strings.ToUpper(l.Code_639_1())
	}
}

func buildReqBodyTarget(target string) string {
	l := lang.Query(target)
	if data[target] != "" {
		return data[target]
	} else {
		return strings.ToUpper(l.Code)
	}

}
