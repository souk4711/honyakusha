package google

import (
	"encoding/json"

	resty "github.com/go-resty/resty/v2"

	"github.com/souk4711/honyakusha/internal/res"
)

type Resp struct {
	Sentences []RespSentence `json:"sentences"`
	Source    string         `json:"src"`
}

type RespSentence struct {
	Trans   string `json:"trans"`
	Orig    string `json:"orig"`
	Backend int32  `json:"backend"`
}

func buildResultFromResp(resp *resty.Response) res.ResTranslator {
	r := new(Resp)
	if resp.StatusCode() != 200 {
		return res.NewResTranslatorFailure(resp.Status())
	}

	if err := json.Unmarshal(resp.Body(), &r); err != nil {
		return res.NewResTranslatorFailure(err.Error())
	}

	sentences := r.Sentences
	if len(sentences) > 0 {
		var t string
		for _, s := range sentences {
			t = t + s.Trans + "\n"
		}
		return res.NewResTranslatorSuccess(t)
	} else {
		return res.NewResTranslatorFailure("")
	}
}
