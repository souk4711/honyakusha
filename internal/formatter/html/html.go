package html

import (
	"bytes"
	"embed"

	"github.com/unrolled/render"

	"github.com/souk4711/honyakusha/internal/res"
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
	var buff bytes.Buffer
	if err := r.HTML(&buff, 0, "index", res); err != nil {
		return err.Error()
	} else {
		return buff.String()
	}
}
