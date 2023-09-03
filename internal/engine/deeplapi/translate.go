package deeplapi

import (
	resty "github.com/go-resty/resty/v2"

	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/res"
)

func TranslateText(text string, from string, to string, conf conf.ConfTranslator) res.ResTranslator {
	client := buildClient(conf)
	if resp, err := makeRequest(client, text, from, to, conf); err != nil {
		return res.NewResTranslatorFailure(err.Error())
	} else {
		return buildResultFromResp(resp)
	}
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

func makeRequest(client *resty.Client, text string, from string, to string, conf conf.ConfTranslator) (*resty.Response, error) {
	return client.R().
		SetBody(buildReqBody(text, from, to)).
		Post(buildReqURL(conf))
}
