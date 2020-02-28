// SPDX-License-Identifier: AGPL-3.0-only

package core_test

import (
	"path/filepath"
	"testing"

	"github.com/brainupdaters/drlm-common/pkg/core"
	"github.com/brainupdaters/drlm-common/pkg/test"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type TestClientSuite struct {
	test.Test
}

func TestClient(t *testing.T) {
	suite.Run(t, new(TestClientSuite))
}

func (s *TestClientSuite) TestNewClient() {
	const certPath = "cert/server.crt"

	s.Run("should work as expected with TLS", func() {
		fs := afero.NewMemMapFs()
		s.GenerateCert(fs, "server", filepath.Dir(certPath))

		_, conn := core.NewClient(fs, true, certPath, "host", 1312)
		s.NotEqual(&grpc.ClientConn{}, conn)
	})

	s.Run("should work as expected without TLS", func() {
		fs := afero.NewMemMapFs()

		_, conn := core.NewClient(fs, false, "", "host", 1312)
		s.NotEqual(&grpc.ClientConn{}, conn)
	})

	s.Run("should exit if there's an error loading the TLS certificate", func() {
		fs := afero.NewMemMapFs()

		s.Exits(func() { core.NewClient(fs, true, certPath, "host", 1312) })
	})
}
