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

package tests_test

import (
	"testing"

	"github.com/brainupdaters/drlm-common/utils/tests"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestGenerateCert(t *testing.T) {
	assert := assert.New(t)

	t.Run("should work as expected", func(t *testing.T) {
		fs := afero.NewMemMapFs()

		tests.GenerateCert(t, fs, "testcert", "/etc/certs")

		const keyFile string = "/etc/certs/testcert.key"
		const crtFile string = "/etc/certs/testcert.crt"

		assert.True(afero.Exists(fs, keyFile))
		assert.True(afero.Exists(fs, crtFile))

		key, err := afero.ReadFile(fs, keyFile)
		assert.Nil(err)
		assert.NotEmpty(key)

		crt, err := afero.ReadFile(fs, crtFile)
		assert.Nil(err)
		assert.NotEmpty(crt)
	})
}
