package os

import (
	"errors"
	"fmt"
	stdOS "os"
	"strings"
)

// CmdFSHome returns the home directory of an user
// Heavily inspired (aka mostly copied) from github.com/mitchellh/go-homedir
func (os OS) CmdFSHome(c Client, usr string) (string, error) {
	switch {
	case os.IsUnix():
		if os == Darwin {
			out, err := c.Exec("sh", "-c", fmt.Sprintf(`dscl -q . -read /Users/%s NFSHomeDirectory | sed 's/^[^ ]*: //'`, usr))
			if err != nil {
				return "", fmt.Errorf("error getting the user home directory: %v", err)
			}

			home := strings.TrimSpace(string(out))
			if home == "" {
				return "", fmt.Errorf("error getting the user home directory: directory path is empty")
			}

			return home, nil
		}

		out, err := c.Exec("getent", "passwd", usr)
		if err != nil {
			return "", fmt.Errorf("error getting the user home directory: %v", err)
		}

		passwd := strings.TrimSpace(string(out))
		if passwd == "" {
			return "", errors.New("error getting the user home directory: empty passwd")
		}

		passwdItems := strings.Split(passwd, ":")
		if len(passwdItems) < 6 {
			return "", errors.New("error getting the user home directory: no home directory field in passwd")
		}

		return passwdItems[5], nil

		// TODO: Add windows support
	default:
		return "", ErrUnsupportedOS
	}
}

// CmdFSTempDir returns the absolute path of the tmp dir
func (os OS) CmdFSTempDir() string {
	switch {
	case os.IsUnix():
		return "/tmp"

		// TODO: Add windows support
	default:
		return ""
	}
}

// CmdFSCheckDir checks if a directory exists
func (os OS) CmdFSCheckDir(c Client, dir string) (bool, error) {
	switch {
	case os.IsUnix():
		out, err := c.Exec("sh", "-c", fmt.Sprintf(`test -d "%q" && echo "dir exists" || echo "dir not found"`, dir))
		if err != nil {
			return false, fmt.Errorf("error checking the directory existence: %v", err)
		}

		if strings.TrimSpace(string(out)) == "dir exists" {
			return true, nil
		}

		return false, nil

		// TODO: Windows support
	default:
		return false, ErrUnsupportedOS
	}
}

// CmdFSMkdir creates a new directory
func (os OS) CmdFSMkdir(c Client, dir string) error {
	switch {
	case os.IsUnix():
		_, err := c.ExecAdmin(os, "mkdir", dir)
		if err != nil {
			return fmt.Errorf("error creating the directory: %v", err)
		}

		return nil

		// TODO: Windows support
	default:
		return ErrUnsupportedOS
	}
}

// CmdFSCopy copies an item from one path to another. It keeps the permissions and it's recursive
func (os OS) CmdFSCopy(c Client, src, dst string) error {
	switch {
	case os.IsUnix():
		if _, err := c.ExecAdmin(os, "cp", "-rp", src, dst); err != nil {
			return fmt.Errorf("error copying the item: %v", err)
		}

		return nil

		// TODO: Windows support
	default:
		return ErrUnsupportedOS
	}
}

// CmdFSMove moves an item from one path to another
func (os OS) CmdFSMove(c Client, src, dst string) error {
	switch {
	case os.IsUnix():
		if _, err := c.ExecAdmin(os, "mv", src, dst); err != nil {
			return fmt.Errorf("error moving the item: %v", err)
		}

		return nil

		// TODO: Windows support
	default:
		return ErrUnsupportedOS
	}
}

// CmdFSChown changes the owner of a file or a directory
func (os OS) CmdFSChown(c Client, path, usr, grp string) error {
	switch {
	case os.IsUnix():
		var owner string
		if grp == "" {
			owner = usr
		} else {
			owner = fmt.Sprintf("%s:%s", usr, grp)
		}

		_, err := c.ExecAdmin(os, "chown", owner, path)
		if err != nil {
			return fmt.Errorf("error changing the owner: %s", err)
		}

		return nil

		// TODO: Windows support
	default:
		return ErrUnsupportedOS
	}
}

// CmdFSChmod changes the permissions of a file or a directory
func (os OS) CmdFSChmod(c Client, path string, permissions stdOS.FileMode) error {
	switch {
	case os.IsUnix():
		_, err := c.ExecAdmin(os, "chmod", permissions.String(), path)
		if err != nil {
			return fmt.Errorf("error changing the permissions: %v", err)
		}

		return nil

		// TODO: Windows support
	default:
		return ErrUnsupportedOS
	}
}

// CmdFSAppendToFile adds contents to a file. If it doesn't exist, it creates it
func (os OS) CmdFSAppendToFile(c Client, path string, b []byte) error {
	return cmdFSAppendToFile(os, c, false, path, b)
}

// CmdFSAppendToFileAdmin adds contents to a file. If it doesn't exist, it creates it. It's done as admin
func (os OS) CmdFSAppendToFileAdmin(c Client, path string, b []byte) error {
	return cmdFSAppendToFile(os, c, true, path, b)
}

func cmdFSAppendToFile(os OS, c Client, asAdmin bool, path string, b []byte) error {
	switch {
	case os.IsUnix():
		cmd := []string{"echo", fmt.Sprintf(`"%q"`, b), ">>", path}
		if asAdmin {
			_, err := c.ExecAdmin(os, cmd[0], cmd[1:]...)
			if err != nil {
				return fmt.Errorf("error appending the content to the file: %v", err)
			}
		} else {
			_, err := c.Exec(cmd[0], cmd[1:]...)
			if err != nil {
				return fmt.Errorf("error appending the content to the file: %v", err)
			}
		}

		return nil

	default:
		return ErrUnsupportedOS
	}
}
