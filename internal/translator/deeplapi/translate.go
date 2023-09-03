package deeplapi

import (
	resty "github.com/go-resty/resty/v2"

	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/res"
)

func TranslateText(text string, source string, target string, conf conf.ConfTranslator) res.ResTranslator {
	client := buildClient(conf)
	return makeRequest(client, text, source, target, conf)
}

func buildClient(conf conf.ConfTranslator) *resty.Client {
	client := resty.New()
	if conf.Proxy != "" {
		client.SetProxy(conf.Proxy)
	}
	if conf.ApiKey != "" {
		client.
			SetAuthScheme("DeepL-Auth-Key").
			SetAuthToken(conf.ApiKey)
	}
	return client
}

func makeRequest(client *resty.Client, text string, source string, target string, conf conf.ConfTranslator) res.ResTranslator {
	if resp, err := client.R().
		SetBody(buildReqBody(text, source, target)).
		Post(buildReqURL(conf)); err != nil {
		return res.NewResTranslatorFailure(err.Error())
	} else {
		return buildResultFromResp(resp)
	}
}