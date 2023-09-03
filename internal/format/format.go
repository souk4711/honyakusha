package format

import (
	"log"

	"github.com/souk4711/honyakusha/internal/formatter/json"
	"github.com/souk4711/honyakusha/internal/res"
)

func Format(res res.Res, formatter string) string {
	switch formatter {
	case "":
		return json.Format(res)
	case "json":
		return json.Format(res)
	default:
		log.Fatalf("Unknown formmatter `%s'", formatter)
		return ""
	}
}
