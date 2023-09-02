package conf

import (
	"log"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/pelletier/go-toml/v2"
)

type Conf struct {
	Translation Translation `toml:"translation"`
	Translators Translators `toml:"translators"`
}

type Translation struct {
	From string `toml:"from"`
	To   string `toml:"to"`
}

type Translators struct {
	DeeplAPI string `toml:"deepl-api"`
}

type Translator struct {
	Enabled bool   `toml:"enabled"`
	Proxy   string `toml:"proxy"`
	AuthKey string `toml:"auth-key"`
}

func Load() Conf {
	if wd, err := os.Getwd(); err != nil {
		f := filepath.Join(wd, "honyakusha.toml")
		if _, err := os.Stat(f); err != nil && os.IsExist(err) {
			return loadFromFile(f)
		}
	}

	if f, err := xdg.ConfigFile("honyakusha.toml"); err != nil {
		if _, err := os.Stat(f); err != nil && os.IsExist(err) {
			return loadFromFile(f)
		}
	}

	return loadDefault()
}

func loadFromFile(filePath string) Conf {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	conf := Conf{}
	if err := toml.Unmarshal(data, &conf); err != nil {
		log.Fatal(err)
	}

	return conf
}

func loadDefault() Conf {
	return Conf{}
}
