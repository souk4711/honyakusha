package deeplapi

import (
	resty "github.com/go-resty/resty/v2"

	"github.com/souk4711/honyakusha/internal/res"
)

func buildResultFromResp(resp *resty.Response) res.Translator {
	println(resp)
	return res.NewTranslatorSuccess()
}
