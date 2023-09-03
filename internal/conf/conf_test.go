package conf_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/souk4711/honyakusha/internal/conf"
)

func TestFilePath(t *testing.T) {
	home, _ := os.UserHomeDir()
	xdgConfPath, _ := filepath.Abs(filepath.Join(home, ".config/honyakusha.toml"))
	cwdConfPath, _ := filepath.Abs(filepath.Join("", "honyakusha.toml"))

	_ = exec.Command("touch", xdgConfPath).Run()
	_ = exec.Command("touch", cwdConfPath).Run()
  defer func () { _ = exec.Command("rm", cwdConfPath).Run() }()

	f := conf.FilePath()
	assert.Equal(t, cwdConfPath, f)

	_ = exec.Command("rm", cwdConfPath).Run()
	f = conf.FilePath()
	assert.Equal(t, xdgConfPath, f)

	if _, ok := os.LookupEnv("GITHUB_WORKSPACE"); ok {
		_ = exec.Command("rm", xdgConfPath).Run()
		f = conf.FilePath()
		assert.Equal(t, "", f)
	}
}
