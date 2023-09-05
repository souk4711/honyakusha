package conf

import (
	"embed"
	"log"
)

var (
	//go:embed templates/*.toml
	embeddedTemplates embed.FS
)

func Sample() []byte {
	data, err := embeddedTemplates.ReadFile("templates/honyakusha.toml")
	if err != nil {
		log.Fatal(err)
	}

	return data
}
