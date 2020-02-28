// SPDX-License-Identifier: AGPL-3.0-only

package test_test

import (
	"testing"

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
	fs := afero.NewMemMapFs()

	s.GenerateCert(fs, "testcert", "/etc/certs")

	const keyFile string = "/etc/certs/testcert.key"
	const crtFile string = "/etc/certs/testcert.crt"

	s.True(afero.Exists(fs, keyFile))
	s.True(afero.Exists(fs, crtFile))

	key, err := afero.ReadFile(fs, keyFile)
	s.Nil(err)
	s.NotEmpty(key)

	crt, err := afero.ReadFile(fs, crtFile)
	s.Nil(err)
	s.NotEmpty(crt)
}
