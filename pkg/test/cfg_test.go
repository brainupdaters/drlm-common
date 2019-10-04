package test_test

import (
	"testing"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	"github.com/brainupdaters/drlm-common/pkg/test"
	"github.com/spf13/afero"

	"github.com/stretchr/testify/suite"
)

type TestCfgSuite struct {
	test.Test
}

func TestCfg(t *testing.T) {
	suite.Run(t, new(TestCfgSuite))
}

func (s *TestCfgSuite) TestGenerate() {
	var executed bool = false
	s.GenerateCfg(test.CfgPathCore, func() { executed = true })

	s.True(afero.Exists(fs.FS, test.CfgPathCore))
	s.True(executed)
}
