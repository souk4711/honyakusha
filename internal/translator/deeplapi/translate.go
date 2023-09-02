package deeplapi

import (
	resty "github.com/go-resty/resty/v2"

	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/res"
)

func TranslateText(text string, from string, to string, conf conf.Translator) res.Translator {
	client := resty.New()

	if conf.Proxy != "" {
		client.SetProxy(conf.Proxy)
	}

	if conf.AuthKey != "" {
		client.
			SetAuthScheme("DeepL-Auth-Key").
			SetAuthToken(conf.AuthKey)
	}

	resp, err := client.R().
		SetBody(buildReqBody(text, from, to)).
		Post(buildReqURL(conf))
	if err != nil {
		return res.NewTranslatorFailure(err.Error())
	}

	return buildResultFromResp(resp)
}
