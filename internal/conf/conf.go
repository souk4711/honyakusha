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
	Source string `toml:"source"`
	Target string `toml:"target"`
}

type ConfTranslators struct {
	Google            ConfTranslator `toml:"google"`
	DeeplAPI          ConfTranslator `toml:"deepl-api"`
	LibreTranslateAPI ConfTranslator `toml:"libretranslate-api"`
}

type ConfTranslator struct {
	Enabled bool   `toml:"enabled"`
	Proxy   string `toml:"proxy"`
	URI     string `toml:"uri"`
	ApiKey  string `toml:"api-key"`
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
