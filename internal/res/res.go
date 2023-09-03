package res

type Res struct {
	Code        int             `json:"code"`
	Error       string          `json:"error"`
	Translators []ResTranslator `json:"translators"`
}

type ResTranslator struct {
	Code           int    `json:"code"`
	Error          string `json:"error"`
	TranslatedText string `json:"translatedText"`
}

func NewResSuccess() Res {
	return Res{Code: 0}
}

func NewResFailure(message string) Res {
	return Res{Code: 1, Error: message}
}

func NewResTranslatorSuccess(text string) ResTranslator {
	return ResTranslator{Code: 0, TranslatedText: text}
}

func NewResTranslatorFailure(message string) ResTranslator {
	return ResTranslator{Code: 1, Error: message}
}
