package json

import (
	"encoding/json"
	"fmt"

	"github.com/souk4711/honyakusha/internal/res"
)

const (
	TMPL = `{
    "code": 1,
    "error": "%s",
    "text": "",
    "translators": []
}`
)

func Format(res res.Res) string {
	if data, err := json.MarshalIndent(&res, "", "    "); err != nil {
		return fmt.Sprintf(TMPL, res.Error) + "\n"
	} else {
		return string(data) + "\n"
	}
}
