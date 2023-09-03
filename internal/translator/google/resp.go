package google

import (
	resty "github.com/go-resty/resty/v2"

	"github.com/souk4711/honyakusha/internal/res"
)

func buildResultFromResp(resp *resty.Response) res.ResTranslator {
	return res.NewResTranslatorSuccess("")
}
