package json

import (
	"encoding/json"

	"github.com/souk4711/honyakusha/internal/res"
)

func Format(res res.Res) string {
	if data, err := json.MarshalIndent(&res, "", "  "); err != nil {
		return err.Error()
	} else {
		return string(data) + "\n"
	}
}
