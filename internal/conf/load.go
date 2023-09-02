package conf

import (
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

func Load() Conf {
	if f := FilePath(); f != "" {
		return loadFromFile(f)
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
