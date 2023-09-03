package trans

import (
	"fmt"

	"github.com/souk4711/honyakusha/internal/formatter/json"
	"github.com/souk4711/honyakusha/internal/res"
)

type Formatter struct {
	Code string
}

func (f *Formatter) Format(res res.Res) string {
	switch f.Code {
	case "json":
		return json.Format(res)
	default:
		return fmt.Sprintf("Unsupported Formatter `%s'", f.Code)
	}
}
