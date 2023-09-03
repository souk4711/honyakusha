package trans

import (
	"sync"

	"github.com/souk4711/honyakusha/internal/conf"
	"github.com/souk4711/honyakusha/internal/res"
)

func TranslateText(text string, c conf.Conf) res.Res {
	translators := availableTranslators(c.Translators)

	resChannel := make(chan res.ResTranslator)
	defer close(resChannel)

	wg := sync.WaitGroup{}
	wg.Add(len(translators))

	for _, translator := range translators {
		go func(translator Translator) {
			r := translator.translateText(text, c.Translate.Source, c.Translate.Target)
			resChannel <- r
			wg.Done()
		}(translator)
	}

	res := res.NewResSuccess()
	for range translators {
		r := <-resChannel
		res.Translators = append(res.Translators, r)
	}

	wg.Wait()
	return res
}
