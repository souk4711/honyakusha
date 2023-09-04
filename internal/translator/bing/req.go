package bing

import (
	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/lang"
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
	} else {
		return l.Code
	}
}

func buildFormDataTarget(target string) string {
  l := lang.Query(target)
  return l.Code
}
