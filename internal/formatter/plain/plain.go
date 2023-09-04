package plain

import (
	"github.com/souk4711/honyakusha/internal/res"
)

func Format(res res.Res) string {
  if res.Code != 0 {
    return res.Error
  }

  r := ""
	for _, translator := range res.Translators {
    r = r + translator.Translator + ":\n"
    if translator.Code != 0 {
      r = r + "\n\t" + translator.Error + "\n\n"
    } else {
      r = r + "\n\t" + translator.TranslatedText + "\n\n"
    }
  }
	return r
}
