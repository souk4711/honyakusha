package helpers

import (
	"os"

	resty "github.com/go-resty/resty/v2"

	"github.com/souk4711/honyakusha/internal/conf"
)

func BuildClient(conf conf.ConfTranslator) *resty.Client {
	client := resty.New()

	if os.Getenv("DEBUG") != "" {
		client.SetDebug(true)
	}

	if conf.Proxy != "" {
		client.SetProxy(conf.Proxy)
	}

	return client
}
