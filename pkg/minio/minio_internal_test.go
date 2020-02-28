// SPDX-License-Identifier: AGPL-3.0-only

package minio

import (
	"net/http"
	"testing"

	"github.com/brainupdaters/drlm-common/pkg/test"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/suite"
)

type TestMinioSuite struct {
	test.Test
}

func TestMinio(t *testing.T) {
	suite.Run(t, new(TestMinioSuite))
}

func (s *TestMinioSuite) TestConn() {
	endpoint, aKey, sKey, ssl := conn("localhost", 9443, "drlm3minio", "drlm3minio", true)

	s.Equal("localhost:9443", endpoint)
	s.Equal("drlm3minio", aKey)
	s.Equal("drlm3minio", sKey)
	s.True(ssl)
}

func (s *TestMinioSuite) TestTransport() {
	s.Run("should return a correct transport", func() {
		fs := afero.NewMemMapFs()
		s.GenerateCert(fs, "minio", "cert")

		tr := http.DefaultTransport.(*http.Transport)
		tr.TLSClientConfig = nil
		s.Nil(transport(fs, tr, "cert/minio.crt"))
	})

	s.Run("should return an error if there's an error reading the cert", func() {
		fs := afero.NewMemMapFs()

		tr := http.DefaultTransport.(*http.Transport)
		s.EqualError(transport(fs, tr, "cert/minio.crt"), "error creating the minio http transport: error reading the certificate: open cert/minio.crt: file does not exist")
	})

	s.Run("should return an error if there's SOMETHING", func() {
		fs := afero.NewMemMapFs()

		s.Require().Nil(afero.WriteFile(fs, "cert/minio.crt", []byte(`invalid cert`), 0644))

		tr := http.DefaultTransport.(*http.Transport)
		s.EqualError(transport(fs, tr, "cert/minio.crt"), "error creating the minio http transport: error parsing the certificate")
	})
}
