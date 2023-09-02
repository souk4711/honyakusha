package conf

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

type Conf struct {
	Translate   ConfTranslate   `toml:"translate"`
	Translators ConfTranslators `toml:"translators"`
}

type ConfTranslate struct {
	From string `toml:"from"`
	To   string `toml:"to"`
}

type ConfTranslators struct {
	DeeplAPI ConfTranslator `toml:"deepl-api"`
}

type ConfTranslator struct {
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
