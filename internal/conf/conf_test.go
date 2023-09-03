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
	wd, _ := os.Getwd()
	xdgConfPath, _ := filepath.Abs(filepath.Join(xdg.ConfigHome, "honyakusha.toml"))
	cwdConfPath, _ := filepath.Abs(filepath.Join(wd, "honyakusha.toml"))

	if _, ok := os.LookupEnv("GITHUB_WORKSPACE"); ok {
		f := conf.FilePath()
		assert.Equal(t, "", f)
	}

	_, _ = os.Create(xdgConfPath)
	f := conf.FilePath()
	assert.Equal(t, xdgConfPath, f)

	_, _ = os.Create(cwdConfPath)
	defer func() { os.Remove(cwdConfPath) }()
	f = conf.FilePath()
	assert.Equal(t, cwdConfPath, f)
}
