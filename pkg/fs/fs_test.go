// SPDX-License-Identifier: AGPL-3.0-only

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
