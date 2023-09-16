package helpers

import (
	resty "github.com/go-resty/resty/v2"

	"github.com/souk4711/honyakusha/internal"
	"github.com/souk4711/honyakusha/internal/conf"
)

func BuildClient(conf conf.ConfTranslator) *resty.Client {
	client := resty.New()
	client.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36 Edg/104.0.1293.54")

	if internal.IsDebugMode() {
		client.SetDebug(true)
	}

	if conf.Proxy != "" {
		client.SetProxy(conf.Proxy)
	}

	return client
}
