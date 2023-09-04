package bing

import (
	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/lang"
)

func buildReqURL(conf conf.ConfTranslator) string {
	if conf.URI == "" {
		return "https://www.bing.com/translator"
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
	return l.Code_639_1()
}

func buildReqBodyTarget(target string) string {
	l := lang.Query(target)
	return l.Code_639_1()
}
