package test

import (
	"github.com/brainupdaters/drlm-common/pkg/fs"

	"github.com/spf13/afero"
)

const (
	// CfgPathCtl is the configuration path of drlmctl
	CfgPathCtl = "/etc/drlm/drlmctl.toml"
	// CfgPathCore is the configuration path of DRLM Core
	CfgPathCore = "/etc/drlm/core.toml"
	// CfgPathAgent is the configuration path of DRLM Agent
	CfgPathAgent = "/etc/drlm/agent.toml"
)

// GenerateCfg generates the configuration required for the tests
func (t *Test) GenerateCfg(path string, init func()) {
	fs.FS = afero.NewMemMapFs()

	t.Require().Nil(afero.WriteFile(fs.FS, path, []byte(""), 0644))

	init()
}
