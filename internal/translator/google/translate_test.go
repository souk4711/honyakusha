package google_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/souk4711/honyakusha/internal/conf"
	translator "github.com/souk4711/honyakusha/internal/translator/google"
)

func TestTranslateText_Success(t *testing.T) {
	recorder, _ := recorder.New("fixtures/TestTranslateText_Success")
	defer func() { _ = recorder.Stop() }()

	conf := conf.ConfTranslator{}
	client := translator.BuildClient(conf)
	client.SetTransport(recorder)
	res := translator.MakeRequest(client,
		"Google Translate is a multilingual neural machine translation service developed by Google"+
			"to translate text, documents and websites from one language into another.",
		"", "zh-CN", conf,
	)
	assert.Equal(t, 0, res.Code)
	assert.Contains(t, res.TranslatedText, "谷歌翻译是谷歌开发的多语言神经机器翻译服务")
}

func TestTranslateText_MultipleLines(t *testing.T) {
	recorder, _ := recorder.New("fixtures/TestTranslateText_MultipleLines")
	defer func() { _ = recorder.Stop() }()

	conf := conf.ConfTranslator{}
	client := translator.BuildClient(conf)
	client.SetTransport(recorder)
	res := translator.MakeRequest(client,
		"On your computer, open Chrome.\n"+
			"Go to a webpage written in another language.\n"+
			"On the right of the address bar, click Translate.\n"+
			"Click on your preferred language.\n"+
			"Chrome will translate your current webpage.\n",
		"", "zh-CN", conf,
	)
	assert.Equal(t, 0, res.Code)
	assert.Contains(t, res.TranslatedText, "在您的计算机上，打开 Chrome。")
	assert.Contains(t, res.TranslatedText, "转到用另一种语言编写的网页。")
	assert.Contains(t, res.TranslatedText, "在地址栏右侧，单击翻译。")
	assert.Contains(t, res.TranslatedText, "单击您的首选语言。")
	assert.Contains(t, res.TranslatedText, "Chrome 将翻译您当前的网页。")
}

func TestTranslateText_SourceUnsupported(t *testing.T) {
	recorder, _ := recorder.New("fixtures/TestTranslateText_SourceUnsupported")
	defer func() { _ = recorder.Stop() }()

	conf := conf.ConfTranslator{}
	client := translator.BuildClient(conf)
	client.SetTransport(recorder)
	res := translator.MakeRequest(client,
		"馬之千里者，一食或盡粟一石",
		"lzh", "zh-CN", conf,
	)
	assert.Equal(t, 1, res.Code)
	assert.Equal(t, res.Error, "400 Bad Request")
}

func TestTranslateText_TargetUnsupported(t *testing.T) {
	recorder, _ := recorder.New("fixtures/TestTranslateText_TargetUnsupported")
	defer func() { _ = recorder.Stop() }()

	conf := conf.ConfTranslator{}
	client := translator.BuildClient(conf)
	client.SetTransport(recorder)
	res := translator.MakeRequest(client,
		"Google Translate is a multilingual neural machine translation service developed by Google"+
			"to translate text, documents and websites from one language into another.",
		"", "lzh", conf,
	)
	assert.Equal(t, 1, res.Code)
	assert.Equal(t, res.Error, "400 Bad Request")
}
