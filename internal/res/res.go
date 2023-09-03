package res

type Res struct {
	Translators []ResTranslator `json:"translators"`
}

type ResTranslator struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Text string `json:"text"`
	} `json:"data"`
}

func NewResTranslatorSuccess(text string) ResTranslator {
	r := ResTranslator{Code: 0, Message: ""}
	r.Data.Text = text
	return r
}

func NewResTranslatorFailure(message string) ResTranslator {
	return ResTranslator{Code: 1, Message: message}
}
