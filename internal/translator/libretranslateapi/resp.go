package libretranslateapi

import (
	"encoding/json"

	resty "github.com/go-resty/resty/v2"

	"github.com/souk4711/honyakusha/internal/res"
)

type Resp struct {
	DetectedLanguage struct {
		Confidence float32 `json:"confidence"`
		Language   string  `json:"language"`
	} `json:"detectedLanguage"`
	TranslatedText string `json:"translatedText"`

	Error string `json:"error"`
}

func buildResultFromResp(resp *resty.Response) res.ResTranslator {
	r := new(Resp)
  if resp.StatusCode() != 200 {
    if err := json.Unmarshal(resp.Body(), &r); err == nil && r.Error != "" {
      return res.NewResTranslatorFailure(r.Error)
    } else {
      return res.NewResTranslatorFailure(resp.Status())
    }
  }

  if err := json.Unmarshal(resp.Body(), &r); err != nil {
      return res.NewResTranslatorFailure(err.Error())
  } else {
    if r.Error != "" {
      return res.NewResTranslatorFailure(r.Error)
    } else {
      return res.NewResTranslatorSuccess(r.TranslatedText)
    }
  }
}
