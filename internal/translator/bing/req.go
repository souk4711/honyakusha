package bing

import (
	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/lang"
)

var (
	data = map[string]string{
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

func buildFormData(text string, source string, target string, conf conf.ConfTranslator, pdata map[string]string) map[string]string {
	return map[string]string{
		"text":     text,
		"fromLang": buildFormDataSource(source),
		"to":       buildFormDataTarget(target),
		"token":    pdata["token"],
		"key":      pdata["key"],
	}
}

func buildFormDataSource(source string) string {
	l := lang.Query(source)
	if l.Code == "" {
		return "auto-detect"
	} else if data[source] != "" {
		return data[source]
	} else {
		return l.Code
	}
}

func buildFormDataTarget(target string) string {
	l := lang.Query(target)
	if data[target] != "" {
		return data[target]
	} else {
		return l.Code
	}
}
