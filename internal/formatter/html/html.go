package html

import (
	"bytes"
	"embed"
	"fmt"

	"github.com/unrolled/render"

	"github.com/souk4711/honyakusha/internal/res"
)

const (
	TMPL_ERROR = `<html>
  <body>
    %s
  </body>
</html>
`
)

var (
	//go:embed templates/*.tmpl
	embeddedTemplates embed.FS

	r = render.New(render.Options{
		Directory: "templates",
		FileSystem: &render.EmbedFileSystem{
			FS: embeddedTemplates,
		},
		Extensions: []string{".tmpl"},
	})
)

func Format(res res.Res) string {
	if res.Code != 0 {
		return fmt.Sprintf(TMPL_ERROR, res.Error)
	}

	var buff bytes.Buffer
	if err := r.HTML(&buff, 0, "index", res); err != nil {
		return fmt.Sprintf(TMPL_ERROR, res.Error)
	} else {
		return buff.String()
	}
}
