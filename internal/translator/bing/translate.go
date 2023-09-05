package bing

import (
	resty "github.com/go-resty/resty/v2"

	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/res"
	"github.com/souk4711/honyakusha/internal/translator/helpers"
)

func TranslateText(text string, source string, target string, conf conf.ConfTranslator) res.ResTranslator {
	client := buildClient(conf)

	pdata, err := preflight(client, conf)
	if err != nil {
		return res.NewResTranslatorFailure("PreflightError: " + err.Error())
	}

	return makeRequest(client, text, source, target, conf, pdata)
}

func buildClient(conf conf.ConfTranslator) *resty.Client {
	client := helpers.BuildClient(conf)
	client.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36 Edg/104.0.1293.54")
	return client
}

func makeRequest(client *resty.Client, text string, source string, target string, conf conf.ConfTranslator, pdata map[string]string) res.ResTranslator {
	if resp, err := client.R().
		SetQueryParam("IG", pdata["ig"]).
		SetQueryParam("IID", pdata["iid"]).
		SetFormData(buildReqFormData(text, source, target, conf, pdata)).
		Post(buildReqURL(conf)); err != nil {
		return res.NewResTranslatorFailure("HTTPError: " + err.Error())
	} else {
		return buildResultFromResp(resp)
	}
}
