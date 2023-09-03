package trans

import (
	"fmt"

	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/res"
	"github.com/souk4711/honyakusha/internal/translator/deeplapi"
	"github.com/souk4711/honyakusha/internal/translator/google"
	"github.com/souk4711/honyakusha/internal/translator/libretranslateapi"
)

const (
	TRANSLATOR_GOOGLE             = "google"
	TRANSLATOR_DEEPL_API          = "deepl-api"
	TRANSLATOR_LIBRETRANSLATE_API = "libretranslate-api"
)

type Translator struct {
	Code string
	Conf conf.ConfTranslator
}

func (t *Translator) TranslateText(text string, source string, target string) res.ResTranslator {
	switch t.Code {
	case TRANSLATOR_GOOGLE:
		return google.TranslateText(text, source, target, t.Conf)
	case TRANSLATOR_DEEPL_API:
		return deeplapi.TranslateText(text, source, target, t.Conf)
	case TRANSLATOR_LIBRETRANSLATE_API:
		return libretranslateapi.TranslateText(text, source, target, t.Conf)
	default:
		return res.NewResTranslatorFailure(fmt.Sprintf("Unsupported Translator `%s'", t.Code))
	}
}

func availableTranslators(c conf.ConfTranslators) []Translator {
	translators := []Translator{}
	addTranslatorIfEnabled := func(ct conf.ConfTranslator, code string) {
		if ct.Enabled {
			translators = append(translators, Translator{Code: code, Conf: ct})
		}
	}

	addTranslatorIfEnabled(c.Google, TRANSLATOR_GOOGLE)
	addTranslatorIfEnabled(c.DeeplAPI, TRANSLATOR_DEEPL_API)
	addTranslatorIfEnabled(c.LibreTranslateAPI, TRANSLATOR_LIBRETRANSLATE_API)
	return translators
}
