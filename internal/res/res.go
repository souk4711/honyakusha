package res

type Res struct {
	Code        int             `json:"code"`
	Error       string          `json:"error"`
	Text        string          `json:"text"`
	Translators []ResTranslator `json:"translators"`
}

type ResTranslator struct {
	Translator struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"translator"`
	Code           int    `json:"code"`
	Error          string `json:"error"`
	TranslatedText string `json:"translatedText"`
}

func NewResSuccess(text string) Res {
	return Res{Code: 0, Text: text, Translators: []ResTranslator{}}
}

func NewResFailure(message string) Res {
	return Res{Code: 1, Error: message, Translators: []ResTranslator{}}
}

func NewResTranslatorSuccess(text string) ResTranslator {
	return ResTranslator{Code: 0, TranslatedText: text}
}

func NewResTranslatorFailure(message string) ResTranslator {
	return ResTranslator{Code: 1, Error: message}
}
