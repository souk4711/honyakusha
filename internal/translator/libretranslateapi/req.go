package libretranslateapi

import (
	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/lang"
)

func buildReqURL(conf conf.ConfTranslator) string {
	if conf.URI == "" {
		return "https://libretranslate.com/translate"
	} else {
		return conf.URI + "/translate"
	}
}

func buildReqBody(text string, source string, target string, conf conf.ConfTranslator) map[string]string {
	return map[string]string{
		"q":       text,
		"source":  buildReqBodySource(source),
		"target":  buildReqBodyTarget(target),
		"format":  "text",
		"api_key": conf.ApiKey,
	}
}

func buildReqBodySource(source string) string {
	l := lang.Query(source)
	if l.Code == "" {
		return "auto"
	} else {
		return l.Code_639_1()
	}
}

func buildReqBodyTarget(target string) string {
	l := lang.Query(target)
	return l.Code_639_1()
}
