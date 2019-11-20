// SPDX-License-Identifier: AGPL-3.0-only

package test_test

import (
	"testing"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	"github.com/brainupdaters/drlm-common/pkg/test"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/suite"
)

type TestCertSuite struct {
	test.Test
}

func TestCert(t *testing.T) {
	suite.Run(t, new(TestCertSuite))
}

func (s *TestCertSuite) TestGenerate() {
	fs.FS = afero.NewMemMapFs()

	s.GenerateCert("testcert", "/etc/certs")

	const keyFile string = "/etc/certs/testcert.key"
	const crtFile string = "/etc/certs/testcert.crt"

	s.True(afero.Exists(fs.FS, keyFile))
	s.True(afero.Exists(fs.FS, crtFile))

	key, err := afero.ReadFile(fs.FS, keyFile)
	s.Nil(err)
	s.NotEmpty(key)

	crt, err := afero.ReadFile(fs.FS, crtFile)
	s.Nil(err)
	s.NotEmpty(crt)
}
