package bing

import (
	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/lang"
)

var (
	data = map[string]string{
		"en-GB": "en",
		"en-US": "en",
		"zh-CN": "zh-Hans",
		"zh-TW": "zh-Hant",
	}
)

func buildReqURL(conf conf.ConfTranslator) string {
	if conf.URI == "" {
		return "https://www.bing.com/ttranslatev3"
	} else {
		return conf.URI + "/ttranslatev3"
	}
}

func buildReqFormData(text string, source string, target string, conf conf.ConfTranslator, pdata map[string]string) map[string]string {
	return map[string]string{
		"text":     text,
		"fromLang": buildReqFormDataSource(source),
		"to":       buildReqFormDataTarget(target),
		"token":    pdata["token"],
		"key":      pdata["key"],
	}
}

func buildReqFormDataSource(source string) string {
	l := lang.Query(source)
	if l.Code == "" {
		return "auto-detect"
	} else if data[source] != "" {
		return data[source]
	} else {
		return l.Code
	}
}

func buildReqFormDataTarget(target string) string {
	l := lang.Query(target)
	if data[target] != "" {
		return data[target]
	} else {
		return l.Code
	}
}
