package conf_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/adrg/xdg"
	"github.com/stretchr/testify/assert"

	"github.com/souk4711/honyakusha/internal/conf"
)

func TestFilePath(t *testing.T) {
	xdgConfPath, _ := filepath.Abs(filepath.Join(xdg.ConfigHome, "honyakusha.toml"))
	cwdConfPath, _ := filepath.Abs("honyakusha.toml")

	_, _ = os.Create(xdgConfPath)
	_, _ = os.Create(cwdConfPath)
	defer func() { os.Remove(cwdConfPath) }()

	f := conf.FilePath()
	assert.Equal(t, cwdConfPath, f)

	_ = os.Remove(cwdConfPath)
	f = conf.FilePath()
	assert.Equal(t, xdgConfPath, f)

	if _, ok := os.LookupEnv("GITHUB_WORKSPACE"); ok {
		_ = os.Remove(xdgConfPath)
		f = conf.FilePath()
		assert.Equal(t, "", f)
	}
}
