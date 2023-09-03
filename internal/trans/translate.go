package trans

import (
	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/res"
)

func Translate(text string, c conf.Conf) res.Res {
	translators := availableTranslators(c.Translators)

	resChannel := make(chan res.ResTranslator)
	defer close(resChannel)

	for _, translator := range translators {
		go func(translator Translator) {
			r := translator.TranslateText(text, c.Translate.Source, c.Translate.Target)
			resChannel <- r
		}(translator)
	}

	res := res.NewResSuccess()
	for range translators {
		r := <-resChannel
		res.Translators = append(res.Translators, r)
	}

	return res
}
