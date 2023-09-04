package bing

import (
	resty "github.com/go-resty/resty/v2"

	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/res"
	"github.com/souk4711/honyakusha/internal/translator/helpers"
)

func TranslateText(text string, source string, target string, conf conf.ConfTranslator) res.ResTranslator {
	client := buildClient(conf)
	return makeRequest(client, text, source, target, conf)
}

func buildClient(conf conf.ConfTranslator) *resty.Client {
	client := helpers.BuildClient(conf)
	return client
}

func makeRequest(client *resty.Client, text string, source string, target string, conf conf.ConfTranslator) res.ResTranslator {
	if resp, err := client.R().
		SetBody(buildReqBody(text, source, target, conf)).
		Post(buildReqURL(conf)); err != nil {
		return res.NewResTranslatorFailure(err.Error())
	} else {
		return buildResultFromResp(resp)
	}
}
