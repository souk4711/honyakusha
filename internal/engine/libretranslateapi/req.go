package libretranslateapi

import (
	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/lang"
)

func buildReqURL(conf conf.ConfTranslator) string {
	if conf.URI == "" {
		return "https://libretranslate.com/translate"
	} else {
		return conf.URI
	}
}

func buildReqBody(text string, from string, to string, conf conf.ConfTranslator) map[string]string {
	return map[string]string{
		"q":       text,
		"source":  buildReqBodySource(from),
		"target":  buildReqBodyTarget(to),
		"format":  "text",
		"api_key": conf.ApiKey,
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
