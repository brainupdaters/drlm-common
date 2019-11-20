// SPDX-License-Identifier: AGPL-3.0-only

package fs

import (
	"github.com/spf13/afero"
)

// FS is the filesystem that the whole program is going to use
var FS afero.Fs

// Init sets the FS to the OS FS
func Init() {
	FS = afero.NewOsFs()
}
