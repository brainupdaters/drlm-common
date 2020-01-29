// SPDX-License-Identifier: AGPL-3.0-only

package minio_test

import (
	"testing"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	"github.com/brainupdaters/drlm-common/pkg/minio"
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

func (s *TestMinioSuite) TestNewSDK() {
	s.Run("should return the sdk correctly", func() {
		fs.FS = afero.NewMemMapFs()
		s.GenerateCert("minio", "cert")

		sdk, err := minio.NewSDK("localhost", 9443, "drlm3minio", "drlm3minio", true, "cert/minio.crt")

		s.NoError(err)
		s.NotNil(sdk)
	})

	s.Run("should return an error if there's an error creating the sdk", func() {
		_, err := minio.NewSDK("://", 9443, "drlm3minio", "drlm3minio", true, "cert/minio.crt")

		s.EqualError(err, "error creating the connection to minio: address ://:9443: too many colons in address")
	})

	s.Run("should return an error if there's an error setting the custom certificates", func() {
		fs.FS = afero.NewMemMapFs()

		_, err := minio.NewSDK("localhost", 9443, "drlm3minio", "drlm3minio", true, "cert/minio.crt")

		s.EqualError(err, "error creating the minio connection: error creating the minio http transport: error reading the certificate: open cert/minio.crt: file does not exist")
	})
}

func (s *TestMinioSuite) TestNewAdminClient() {
	s.Run("should return the client correctly", func() {
		fs.FS = afero.NewMemMapFs()
		s.GenerateCert("minio", "cert")

		cli, err := minio.NewAdminClient("localhost", 9443, "drlm3minio", "drlm3minio", true, "cert/minio.crt")

		s.NoError(err)
		s.NotNil(cli)
	})

	s.Run("should return an error if there's an error creating the client", func() {
		_, err := minio.NewAdminClient("://", 9443, "drlm3minio", "drlm3minio", true, "cert/minio.crt")

		s.EqualError(err, "error creating the minio admin connection: address ://:9443: too many colons in address")
	})

	s.Run("should return an error if there's an error setting the custom certificates", func() {
		fs.FS = afero.NewMemMapFs()

		_, err := minio.NewAdminClient("localhost", 9443, "drlm3minio", "drlm3minio", true, "cert/minio.crt")

		s.EqualError(err, "error creating the minio admin connection: error creating the minio http transport: error reading the certificate: open cert/minio.crt: file does not exist")
	})
}
