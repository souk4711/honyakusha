package trans

import (
	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/res"
)

func Translate(text string, source string, target string, specifiedTranslators []string, c conf.Conf) res.Res {
	if source == "" {
		source = c.Translate.Source
	}
	if target == "" {
		target = c.Translate.Target
	}

	resChannel := make(chan res.ResTranslator)
	defer close(resChannel)

	translators := availableTranslators(c.Translators, specifiedTranslators)
	for _, translator := range translators {
		go func(translator Translator) {
			r := translator.TranslateText(text, source, target)
			resChannel <- r
		}(translator)
	}

	res := res.NewResSuccess(text)
	for range translators {
		r := <-resChannel
		res.Translators = append(res.Translators, r)
	}

	return res
}
