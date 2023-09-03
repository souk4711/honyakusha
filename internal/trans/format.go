package trans

import (
	"github.com/souk4711/honyakusha/internal/res"
)

func Format(res res.Res, format string) string {
	formatter := Formatter{Code: format}
	return formatter.Format(res)
}
