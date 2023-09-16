package deeplapi

import (
	"encoding/json"

	resty "github.com/go-resty/resty/v2"

	"github.com/souk4711/honyakusha/internal/res"
)

type Resp struct {
	Translations []RespTranslation `json:"translations"`
}

type RespTranslation struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
  Text string `json:"text"`
}

func buildResultFromResp(resp *resty.Response) res.ResTranslator {
	r := new(Resp)
	if resp.StatusCode() != 200 {
		return res.NewResTranslatorFailure("ApiError: " + resp.Status())
	}

	if err := json.Unmarshal(resp.Body(), &r); err != nil {
		return res.NewResTranslatorFailure("ApiError: " + err.Error())
	}

	translations := r.Translations
	if len(translations) > 0 {
		return res.NewResTranslatorSuccess(translations[0].Text)
	} else {
		return res.NewResTranslatorFailure("ApiError: No result")
	}
}
