package bing

import (
	"encoding/json"

	resty "github.com/go-resty/resty/v2"

	"github.com/souk4711/honyakusha/internal/res"
)

type Resp struct {
	DetectedLanguage struct {
		Language string  `json:"language"`
		Score    float32 `json:"score"`
	} `json:"detectedLanguage"`
	Translations []RespTranslation `json:"translations"`
}

type RespTranslation struct {
	Text string `json:"text"`
}

func buildResultFromResp(resp *resty.Response) res.ResTranslator {
	var r []Resp
	if resp.StatusCode() != 200 {
		return res.NewResTranslatorFailure("ApiError: " + resp.Status())
	}

	if len(resp.Body()) == 0 {
		return res.NewResTranslatorFailure("ApiError: No result")
	}

	if resp.Body()[0] != '[' {
		return res.NewResTranslatorFailure("ApiError: " + string(resp.Body()))
	}

	if err := json.Unmarshal(resp.Body(), &r); err != nil {
		return res.NewResTranslatorFailure("ApiError: " + err.Error())
	}

	translations := r[0].Translations
	if len(translations) > 0 {
		var t string
		for _, s := range translations {
			t = t + s.Text + "\n"
		}
		return res.NewResTranslatorSuccess(t)
	} else {
		return res.NewResTranslatorFailure("ApiError: No result")
	}
}
