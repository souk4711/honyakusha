package libretranslateapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/souk4711/honyakusha/internal/conf"
	translator "github.com/souk4711/honyakusha/internal/translator/libretranslateapi"
)

func TestTranslateText_Success(t *testing.T) {
	recorder, _ := recorder.New("fixtures/TestTranslateText_Success")
	defer func() { _ = recorder.Stop() }()

	conf := conf.ConfTranslator{URI: "https://translate.terraprint.co"}
	client := translator.BuildClient(conf)
	client.SetTransport(recorder)
	res := translator.MakeRequest(client,
		"Free and Open Source Machine Translation API. Self-hosted, offline capable and easy to setup.",
		"", "zh-CN", conf,
	)
	assert.Equal(t, 0, res.Code)
	assert.Equal(t, "免费和开放的源库。 自我主办的、能够和容易形成的线。", res.TranslatedText)
}

func TestTranslateText_ApiKeyRequired(t *testing.T) {
	recorder, _ := recorder.New("fixtures/TestTranslateText_ApiKeyRequired")
	defer func() { _ = recorder.Stop() }()

	conf := conf.ConfTranslator{}
	client := translator.BuildClient(conf)
	client.SetTransport(recorder)
	res := translator.MakeRequest(client,
		"Free and Open Source Machine Translation API. Self-hosted, offline capable and easy to setup.",
		"", "zh-CN", conf,
	)
	assert.Equal(t, 1, res.Code)
	assert.Contains(t, res.Error, "Visit https://portal.libretranslate.com to get an API key")
}

func TestTranslateText_SourceUnsupported(t *testing.T) {
	recorder, _ := recorder.New("fixtures/TestTranslateText_SourceUnsupported")
	defer func() { _ = recorder.Stop() }()

	conf := conf.ConfTranslator{URI: "https://translate.terraprint.co"}
	client := translator.BuildClient(conf)
	client.SetTransport(recorder)
	res := translator.MakeRequest(client,
		"驢不勝怒，蹄之",
		"lzh", "zh-CN", conf,
	)
	assert.Equal(t, 1, res.Code)
	assert.Contains(t, res.Error, "is not supported")
}

func TestTranslateText_TargetUnsupported(t *testing.T) {
	recorder, _ := recorder.New("fixtures/TestTranslateText_TargetUnsupported")
	defer func() { _ = recorder.Stop() }()

	conf := conf.ConfTranslator{URI: "https://translate.terraprint.co"}
	client := translator.BuildClient(conf)
	client.SetTransport(recorder)
	res := translator.MakeRequest(client,
		"Free and Open Source Machine Translation API. Self-hosted, offline capable and easy to setup.",
		"", "lzh", conf,
	)
	assert.Equal(t, 1, res.Code)
	assert.Contains(t, res.Error, "is not supported")
}
