package trans

import (
	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/engine/deeplapi"
	"github.com/souk4711/honyakusha/internal/res"
)

type Translator struct {
	Code string
	Conf conf.ConfTranslator
}

func (t *Translator) translateText(text string, from string, to string) res.ResTranslator {
	switch t.Code {
	case "deepl-api":
		return deeplapi.TranslateText(text, from, to, t.Conf)
	default:
		return res.NewResTranslatorFailure("")
	}
}

func availableTranslators(c conf.ConfTranslators) []Translator {
	arr := []Translator{}
	if c.DeeplAPI.Enabled {
		arr = append(arr, Translator{Code: "deepl-api", Conf: c.DeeplAPI})
	}
	return arr
}
