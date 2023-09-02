package res

type Res struct {
	Translators Translators
}

type Translators struct {
	DeeplAPI Translator
}

type Translator struct {
	Code    int
	Message string
}

func NewTranslatorSuccess() Translator {
	return Translator{
		Code:    0,
		Message: "",
	}
}

func NewTranslatorFailure(message string) Translator {
	return Translator{
		Code:    1,
		Message: message,
	}
}
