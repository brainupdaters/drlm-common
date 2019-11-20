// SPDX-License-Identifier: AGPL-3.0-only

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

package fs_test

import (
	"testing"

	"github.com/brainupdaters/drlm-common/pkg/fs"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)

	fs.Init()

	assert.IsType(&afero.OsFs{}, fs.FS)
}
