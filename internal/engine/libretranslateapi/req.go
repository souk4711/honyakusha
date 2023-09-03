package libretranslateapi

import (
	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/lang"
)

type Req struct {
	Text   string `json:"q"`
	Source string `json:"source"`
	Target string `json:"target"`
	Format string `json:"format"`
	ApiKey string `json:"api_key"`
}

func buildReqURL(conf conf.ConfTranslator) string {
	if conf.URI == "" {
		return "https://libretranslate.com/translate"
	} else {
		return conf.URI
	}
}

func buildReqBody(text string, from string, to string, conf conf.ConfTranslator) Req {
	return Req{
		Text:   text,
		Source: buildReqBodySource(from),
		Target: buildReqBodyTarget(to),
		Format: "text",
		ApiKey: conf.ApiKey,
	}
}

func buildReqBodySource(from string) string {
	frLang := lang.Query(from)
	code := frLang.Code
	if code == "" {
		return "auto"
	} else {
		return frLang.Macro()
	}
}

func buildReqBodyTarget(to string) string {
	toLang := lang.Query(to)
	return toLang.Macro()
}
