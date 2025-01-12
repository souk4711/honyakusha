package trans

import (
	"fmt"
	"slices"

	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/res"
	"github.com/souk4711/honyakusha/internal/translator/bing"
	"github.com/souk4711/honyakusha/internal/translator/deeplapi"
	"github.com/souk4711/honyakusha/internal/translator/google"
	"github.com/souk4711/honyakusha/internal/translator/libretranslateapi"
)

const (
	TRANSLATOR_BING               = "bing"
	TRANSLATOR_GOOGLE             = "google"
	TRANSLATOR_DEEPL_API          = "deepl-api"
	TRANSLATOR_LIBRETRANSLATE_API = "libretranslate-api"
)

type Translator struct {
	Code string
	Conf conf.ConfTranslator
}

func (t *Translator) Name() string {
	switch t.Code {
	case TRANSLATOR_BING:
		return "Bing Translator"
	case TRANSLATOR_GOOGLE:
		return "Google Translate"
	case TRANSLATOR_DEEPL_API:
		return "DeepL Translate"
	case TRANSLATOR_LIBRETRANSLATE_API:
		return "LibreTranslate"
	default:
		return t.Code
	}
}

func (t *Translator) TranslateText(text string, source string, target string) res.ResTranslator {
	var r res.ResTranslator

	switch t.Code {
	case TRANSLATOR_BING:
		r = bing.TranslateText(text, source, target, t.Conf)
	case TRANSLATOR_GOOGLE:
		r = google.TranslateText(text, source, target, t.Conf)
	case TRANSLATOR_DEEPL_API:
		r = deeplapi.TranslateText(text, source, target, t.Conf)
	case TRANSLATOR_LIBRETRANSLATE_API:
		r = libretranslateapi.TranslateText(text, source, target, t.Conf)
	default:
		r = res.NewResTranslatorFailure(fmt.Sprintf("Unsupported translator `%s'", t.Code))
	}

	r.Translator.Code = t.Code
	r.Translator.Name = t.Name()
	return r
}

func availableTranslators(c conf.ConfTranslators, specifiedTranslators []string) []Translator {
	translators := []Translator{}

	addTranslatorIfEnabled := func(ct conf.ConfTranslator, code string) {
		if len(specifiedTranslators) != 0 {
			if slices.Contains(specifiedTranslators, code) {
				translators = append(translators, Translator{Code: code, Conf: ct})
			}
			return
		}

		if ct.Enabled {
			translators = append(translators, Translator{Code: code, Conf: ct})
		}
	}

	addTranslatorIfEnabled(c.Bing, TRANSLATOR_BING)
	addTranslatorIfEnabled(c.Google, TRANSLATOR_GOOGLE)
	addTranslatorIfEnabled(c.DeeplAPI, TRANSLATOR_DEEPL_API)
	addTranslatorIfEnabled(c.LibreTranslateAPI, TRANSLATOR_LIBRETRANSLATE_API)
	return translators
}
