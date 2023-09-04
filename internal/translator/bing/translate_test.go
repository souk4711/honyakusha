package bing_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/souk4711/honyakusha/internal/conf"
	translator "github.com/souk4711/honyakusha/internal/translator/bing"
)

func TestTranslateText_Success(t *testing.T) {
	recorder, _ := recorder.New("fixtures/TestTranslateText_Success")
	defer func() { _ = recorder.Stop() }()

	conf := conf.ConfTranslator{}
	client := translator.BuildClient(conf)
	client.SetTransport(recorder)
	pdata, _ := translator.Preflight(client, conf)
	res := translator.MakeRequest(client,
		"Quickly translate words and phrases between English and over 100 languages.",
		"在英语和 100 多种语言之间快速翻译单词和短语。", "zh-CN", conf, pdata,
	)
	assert.Equal(t, 0, res.Code)
	assert.Contains(t, res.TranslatedText, "")
}

func TestTranslateText_MultipleLines(t *testing.T) {
	recorder, _ := recorder.New("fixtures/TestTranslateText_MultipleLines")
	defer func() { _ = recorder.Stop() }()

	conf := conf.ConfTranslator{}
	client := translator.BuildClient(conf)
	client.SetTransport(recorder)
	pdata, _ := translator.Preflight(client, conf)
	res := translator.MakeRequest(client,
		"なんとしてでも、地球を死の惑星にはしたくない。\n"+
			"未来に向かって、地球上のすべての生物との共存をめざし、むしろこれからが、人類のほんとうの“あけぼの”なのかもしれないとも思うのです。",
		"ja", "zh-CN", conf, pdata,
	)
	assert.Equal(t, 0, res.Code)
	assert.Contains(t, res.TranslatedText, "我不希望地球不惜一切代价成为一个死亡的星球。")
	assert.Contains(t, res.TranslatedText, "面向未来，我们的目标是与地球上的所有生物共存，我认为未来可能是人类真正的“黎明”")
}

func TestTranslateText_SourceUnsupported(t *testing.T) {
	recorder, _ := recorder.New("fixtures/TestTranslateText_SourceUnsupported")
	defer func() { _ = recorder.Stop() }()

	conf := conf.ConfTranslator{}
	client := translator.BuildClient(conf)
	client.SetTransport(recorder)
	pdata, _ := translator.Preflight(client, conf)
	res := translator.MakeRequest(client,
		"",
		"ii", "zh-CN", conf, pdata,
	)
	assert.Equal(t, 1, res.Code)
	assert.Contains(t, res.Error, "400")
}

func TestTranslateText_TargetUnsupported(t *testing.T) {
	recorder, _ := recorder.New("fixtures/TestTranslateText_TargetUnsupported")
	defer func() { _ = recorder.Stop() }()

	conf := conf.ConfTranslator{}
	client := translator.BuildClient(conf)
	client.SetTransport(recorder)
	pdata, _ := translator.Preflight(client, conf)
	res := translator.MakeRequest(client,
		"Quickly translate words and phrases between English and over 100 languages.",
		"", "ii", conf, pdata,
	)
	assert.Equal(t, 1, res.Code)
	assert.Contains(t, res.Error, "400")
}
