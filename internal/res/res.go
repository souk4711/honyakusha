package res

type Res struct {
	Translators []ResTranslator `json:"translators"`
}

type ResTranslator struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewResTranslatorSuccess() ResTranslator {
	return ResTranslator{
		Code:    0,
		Message: "",
	}
}

func NewResTranslatorFailure(message string) ResTranslator {
	return ResTranslator{
		Code:    1,
		Message: message,
	}
}
