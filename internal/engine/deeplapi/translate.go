package deeplapi

import (
	resty "github.com/go-resty/resty/v2"

	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/res"
)

func TranslateText(text string, from string, to string, conf conf.ConfTranslator) res.ResTranslator {
	client := resty.New()

	if conf.Proxy != "" {
		client.SetProxy(conf.Proxy)
	}

	if conf.ApiKey != "" {
		client.
			SetAuthScheme("DeepL-Auth-Key").
			SetAuthToken(conf.ApiKey)
	}

	resp, err := client.R().
		SetBody(buildReqBody(text, from, to)).
		Post(buildReqURL(conf))
	if err != nil {
		return res.NewResTranslatorFailure(err.Error())
	}

	return buildResultFromResp(resp)
}
