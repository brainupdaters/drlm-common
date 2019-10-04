/*
 * Copyright (C) 2019 IsardVDI
 * Authors: IsardVDI Authors
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

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
