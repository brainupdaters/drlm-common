// SPDX-License-Identifier: AGPL-3.0-only

package core

import (
	"crypto/x509"
	"path/filepath"
	"testing"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	"github.com/brainupdaters/drlm-common/pkg/test"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/suite"
)

type TestClientInternalSuite struct {
	test.Test
}

func TestClientInternal(t *testing.T) {
	suite.Run(t, new(TestClientInternalSuite))
}

func (s *TestClientInternalSuite) TestReadCert() {
	const certPath = "cert/server.crt"

	s.Run("should read the certificate correctly", func() {
		fs.FS = afero.NewMemMapFs()
		s.GenerateCert("server", filepath.Dir(certPath))

		cp, err := readCert(certPath)
		s.Nil(err)
		s.NotEqual(&x509.CertPool{}, cp)
	})

	s.Run("should return an error if there's an error reading the certificate file", func() {
		fs.FS = afero.NewMemMapFs()

		cp, err := readCert(certPath)
		s.EqualError(err, "error reading the certificate file: open cert/server.crt: file does not exist")
		s.Equal(&x509.CertPool{}, cp)
	})

	s.Run("should return an error if there's an error parsing the certificate", func() {
		fs.FS = afero.NewMemMapFs()
		afero.WriteFile(fs.FS, certPath, []byte("This isn't a cert!"), 0644)

		cp, err := readCert(certPath)
		s.EqualError(err, "error parsing the certificate: invalid certificate")
		s.Equal(&x509.CertPool{}, cp)
	})
}
