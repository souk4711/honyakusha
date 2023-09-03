package libretranslateapi

import (
	resty "github.com/go-resty/resty/v2"

	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/res"
)

func TranslateText(text string, from string, to string, conf conf.ConfTranslator) res.ResTranslator {
	client := buildClient(conf)
	return makeRequest(client, text, from, to, conf)
}

func buildClient(conf conf.ConfTranslator) *resty.Client {
	client := resty.New()
	if conf.Proxy != "" {
		client.SetProxy(conf.Proxy)
	}
	return client
}

func makeRequest(client *resty.Client, text string, from string, to string, conf conf.ConfTranslator) res.ResTranslator {
	if resp, err := client.R().
		SetFormData(buildReqBody(text, from, to, conf)).
		Post(buildReqURL(conf)); err != nil {
		return res.NewResTranslatorFailure(err.Error())
	} else {
		return buildResultFromResp(resp)
	}
}
