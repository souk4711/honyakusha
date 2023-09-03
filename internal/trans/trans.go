package trans

import (
	"sync"

	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/res"
)

func TranslateText(text string) {
	c := conf.Load()
	translateText(text, c)
}

func translateText(text string, c conf.Conf) res.Res {
	translators := availableTranslators(c.Translators)

	resChannel := make(chan res.ResTranslator)
	defer close(resChannel)

	wg := sync.WaitGroup{}
	wg.Add(len(translators))
	for _, translator := range translators {
		go func(translator Translator) {
			defer wg.Done()
			r := translator.translateText(text, c.Translate.Source, c.Translate.Target)
			resChannel <- r
		}(translator)
	}
	wg.Wait()

	return res.Res{}
}
