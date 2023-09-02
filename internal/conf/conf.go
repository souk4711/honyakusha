package conf

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

type Conf struct {
	Translate   Translate   `toml:"translate"`
	Translators Translators `toml:"translators"`
}

type Translate struct {
	From string `toml:"from"`
	To   string `toml:"to"`
}

type Translators struct {
	DeeplAPI Translator `toml:"deepl-api"`
}

type Translator struct {
	Enabled bool   `toml:"enabled"`
	Proxy   string `toml:"proxy"`
	URI     string `toml:"uri"`
	AuthKey string `toml:"auth-key"`
}

func FilePath() string {
	if wd, err := os.Getwd(); err == nil {
		f := filepath.Join(wd, "honyakusha.toml")
		if _, err := os.Stat(f); err == nil {
			return f
		}
	}

	if f, err := xdg.ConfigFile("honyakusha.toml"); err == nil {
		if _, err := os.Stat(f); err == nil {
			return f
		}
	}

	return ""
}
