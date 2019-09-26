package os

import (
	"fmt"
	"path/filepath"
)

// CmdPkgInstallBinary installs a binary in a system
func (os OS) CmdPkgInstallBinary(c Client, name string, b []byte) error {
	switch os {
	case Linux:
		if err := os.CmdFSAppendToFileAdmin(c, filepath.Join("/usr/sbin", name), b); err != nil {
			return fmt.Errorf("error installing the binary: %v", err)
		}

		return nil

	default:
		return ErrUnsupportedOS
	}
}
