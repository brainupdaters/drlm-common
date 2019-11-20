// SPDX-License-Identifier: AGPL-3.0-only

package test_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/brainupdaters/drlm-common/pkg/test"

	"github.com/stretchr/testify/suite"
)

type TestFreePortSuite struct {
	test.Test
}

func TestFreePort(t *testing.T) {
	suite.Run(t, new(TestFreePortSuite))
}

func (s *TestFreePortSuite) TestFreePort() {
	port := s.FreePort()

	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	defer l.Close()

	s.Nil(err)

}
