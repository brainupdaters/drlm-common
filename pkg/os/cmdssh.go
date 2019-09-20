package os

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

// CmdSSHGetHostKeys returns the public SSH keys of a host
func (os OS) CmdSSHGetHostKeys(c Client, host string, port int) ([]string, error) {
	switch {
	case os.IsUnix():
		out, err := c.Exec("ssh-keyscan", "-p", strconv.Itoa(port), host)
		if err != nil {
			return []string{}, fmt.Errorf("error getting the host SSH keys: %v", err)
		}

		var keys []string
		for _, l := range strings.Split(string(out), "\n") {
			if !strings.HasPrefix(l, "# ") {
				keys = append(keys, l)
			}
		}

		return keys, nil

	default:
		return []string{}, ErrUnsupportedOS
	}
}

// CmdSSHCopyID copies the key to the OS
func (os OS) CmdSSHCopyID(c Client, usr string, key []byte) error {
	switch {
	case os.IsUnix():
		home, err := os.CmdFSHome(c, usr)
		if err != nil {
			return fmt.Errorf("error copying the SSH key: %v", err)
		}

		grp, err := os.CmdUserGroup(c, usr)
		if err != nil {
			return err
		}

		sshDir := filepath.Join(home, ".ssh")
		exists, err := os.CmdFSCheckDir(c, sshDir)
		if err != nil {
			return fmt.Errorf("error checking the SSH directory: %v", err)
		}

		// If the SSH directory doesn't exist, create it
		if !exists {
			err := os.CmdFSMkdir(c, sshDir)
			if err != nil {
				return fmt.Errorf("error creating the SSH directory: %v", err)
			}

			err = os.CmdFSChown(c, sshDir, usr, grp)
			if err != nil {
				return fmt.Errorf("error changing the SSH directory owner: %v", err)
			}

			err = os.CmdFSChmod(c, sshDir, 0700)
			if err != nil {
				return fmt.Errorf("error changing the SSH directory permissions: %v", err)
			}
		}

		authKeys := filepath.Join(sshDir, "authorized_keys")
		tmpAuthKeys := filepath.Join(os.CmdFSTempDir(), "drlm_core_authorized_keys")
		err = os.CmdFSCopy(c, authKeys, tmpAuthKeys)
		if err != nil {
			return fmt.Errorf("error copying the authorized_keys file: %v", err)
		}

		err = os.CmdFSAppendToFile(c, tmpAuthKeys, key)
		if err != nil {
			return fmt.Errorf("error adding the key to the authorized_keys file: %v", err)
		}

		err = os.CmdFSChown(c, tmpAuthKeys, usr, grp)
		if err != nil {
			return fmt.Errorf("error changing the authorized_keys owner: %v", err)
		}

		err = os.CmdFSChmod(c, tmpAuthKeys, 0600)
		if err != nil {
			return fmt.Errorf("error changing the authorized_keys permissions: %v", err)
		}

		err = os.CmdFSMove(c, tmpAuthKeys, authKeys)
		if err != nil {
			return fmt.Errorf("error replacing the authorized_keys file: %v", err)
		}

		return nil

	default:
		return ErrUnsupportedOS
	}
}
