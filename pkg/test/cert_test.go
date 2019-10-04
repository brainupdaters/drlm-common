/*
 * Copyright (C) 2019 DRLM Project
 * Authors: DRLM Common authors
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
