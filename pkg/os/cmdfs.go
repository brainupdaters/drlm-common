package os

import (
	"errors"
	"fmt"
	"strings"

	"github.com/brainupdaters/drlm-common/pkg/os/client"
)

// CmdFSHome returns the home directory of an user
// Heavily inspired (aka mostly copied) from github.com/mitchellh/go-homedir
func (os OS) CmdFSHome(c client.Client, usr string) (string, error) {
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
