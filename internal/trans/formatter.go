package trans

import (
	"fmt"

	"github.com/souk4711/honyakusha/internal/formatter/json"
	"github.com/souk4711/honyakusha/internal/formatter/plain"
	"github.com/souk4711/honyakusha/internal/res"
)

const (
	FORMATTER_JSON  = "json"
	FORMATTER_PLAIN = "plain"
)

type Formatter struct {
	Code string
}

func (f *Formatter) Format(res res.Res) string {
	switch f.Code {
	case FORMATTER_JSON:
		return json.Format(res)
	case FORMATTER_PLAIN:
		return plain.Format(res)
	default:
		return fmt.Sprintf("Unsupported Formatter `%s'", f.Code)
	}
}
