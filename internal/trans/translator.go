package trans

import (
	"fmt"

	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/res"
	"github.com/souk4711/honyakusha/internal/translator/deeplapi"
	"github.com/souk4711/honyakusha/internal/translator/google"
	"github.com/souk4711/honyakusha/internal/translator/libretranslateapi"
)

type Translator struct {
	Code string
	Conf conf.ConfTranslator
}

func (t *Translator) TranslateText(text string, source string, target string) res.ResTranslator {
	switch t.Code {
	case "google":
		return google.TranslateText(text, source, target, t.Conf)
	case "deepl-api":
		return deeplapi.TranslateText(text, source, target, t.Conf)
	case "libretranslate-api":
		return libretranslateapi.TranslateText(text, source, target, t.Conf)
	default:
		return res.NewResTranslatorFailure(fmt.Sprintf("Unsupported Translator `%s'", t.Code))
	}
}

func availableTranslators(c conf.ConfTranslators) []Translator {
	arr := []Translator{}
	if c.Google.Enabled {
		arr = append(arr, Translator{Code: "google", Conf: c.Google})
	}
	if c.DeeplAPI.Enabled {
		arr = append(arr, Translator{Code: "deepl-api", Conf: c.DeeplAPI})
	}
	if c.LibreTranslateAPI.Enabled {
		arr = append(arr, Translator{Code: "libretranslate-api", Conf: c.LibreTranslateAPI})
	}
	return arr
}
