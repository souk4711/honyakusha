package google

import (
	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/lang"
)

func buildReqURL(conf conf.ConfTranslator) string {
	if conf.URI == "" {
		return "https://translate.google.com"
	} else {
		return conf.URI
	}
}

func buildReqBody(text string, source string, target string, conf conf.ConfTranslator) map[string]string {
	return map[string]string{
		"source": buildReqBodySource(source),
		"target": buildReqBodyTarget(target),
	}
}

func buildReqBodySource(source string) string {
	l := lang.Query(source)
	return l.Macro()
}

func buildReqBodyTarget(target string) string {
	l := lang.Query(target)
	return l.Macro()
}
