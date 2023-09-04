package google

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
	client.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36 Edg/104.0.1293.54")
	return client
}

func makeRequest(client *resty.Client, text string, source string, target string, conf conf.ConfTranslator) res.ResTranslator {
	if resp, err := client.R().
		SetQueryParams(buildReqQueryParams(text, source, target, conf)).
		Get(buildReqURL(conf)); err != nil {
		return res.NewResTranslatorFailure("HTTPError: " + err.Error())
	} else {
		return buildResultFromResp(resp)
	}
}
