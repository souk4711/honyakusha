package libretranslateapi

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
	return client
}

func makeRequest(client *resty.Client, text string, source string, target string, conf conf.ConfTranslator) res.ResTranslator {
	if resp, err := client.R().
		SetFormData(buildReqBody(text, source, target, conf)).
		Post(buildReqURL(conf)); err != nil {
		return res.NewResTranslatorSuccess(err.Error())
	} else {
		return buildResultFromResp(resp)
	}
}
