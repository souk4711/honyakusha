package google

import (
	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/lang"
)

func buildReqURL(conf conf.ConfTranslator) string {
	if conf.URI == "" {
		return "https://translate.google.com/translate_a/single"
	} else {
		return conf.URI + "/translate_a/single"
	}
}

func buildReqQueryParams(text string, source string, target string, conf conf.ConfTranslator) map[string]string {
	return map[string]string{
		"client": "gtx", "dt": "t", "dj": "1",
		"q":  text,
		"sl": buildReqQueryParamsSource(source),
		"tl": buildReqQueryParamsTarget(target),
	}
}

func buildReqQueryParamsSource(source string) string {
	l := lang.Query(source)
	code := l.Code
	if code == "" {
		return "auto"
	} else {
		return l.Code_639_1()
	}
}

func buildReqQueryParamsTarget(target string) string {
	l := lang.Query(target)
	return l.Code_639_1()
}
