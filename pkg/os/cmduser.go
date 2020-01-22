// SPDX-License-Identifier: AGPL-3.0-only

package os

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/brainupdaters/drlm-common/pkg/os/client"
	"github.com/rs/xid"
)

// CmdUserCreate creates a new user in the OS
func (os OS) CmdUserCreate(c client.Client, name, pwd string) error {
	switch {
	case os.IsUnix():
		_, err := c.Exec("useradd", "-m", "-c", `"DRLM user"`, fmt.Sprintf(`"%s"`, name))
		if err != nil {
			return fmt.Errorf("error creating the user: %v", err)
		}

		if err := os.CmdUserChangePasswd(c, name, pwd); err != nil {
			return err
		}

		return nil

	default:
		return ErrUnsupportedOS
	}
}

// CmdUserUID returns the uid of an user
func (os OS) CmdUserUID(c client.Client, usr string) (int, error) {
	switch {
	case os.IsUnix():
		out, err := c.Exec("id", "-u", usr)
		if err != nil {
			return 0, fmt.Errorf("error getting the user UID: %v", err)
		}

		uid, err := strconv.Atoi(strings.TrimSpace(string(out)))
		if err != nil {
			return 0, fmt.Errorf("error parsing the UID: %v", err)
		}
		return uid, nil

		// TODO: Windows support
	default:
		return 0, ErrUnsupportedOS
	}
}

// CmdUserGID returns the gid of an user
func (os OS) CmdUserGID(c client.Client, usr string) (int, error) {
	switch {
	case os.IsUnix():
		out, err := c.Exec("id", "-g", usr)
		if err != nil {
			return 0, fmt.Errorf("error getting the user GID: %v", err)
		}

		gid, err := strconv.Atoi(strings.TrimSpace(string(out)))
		if err != nil {
			return 0, fmt.Errorf("error parsing the GID: %v", err)
		}
		return gid, nil

		// TODO: Windows support
	default:
		return 0, ErrUnsupportedOS
	}
}

// CmdUserDisable disables an user
func (os OS) CmdUserDisable(c client.Client, usr string) error {
	switch {
	case os.IsUnix():
		distro, _, err := os.DetectDistro(c)
		if err != nil {
			return fmt.Errorf("error disabling the user: %v", err)
		}

		switch distro {
		case "alpine":
			// TODO: Maybe this should be an actual random password and not an "uuid"?
			return os.CmdUserChangePasswd(c, usr, xid.New().String())

		default:
			_, err := c.Exec("passwd", "-l", usr)
			if err != nil {
				return fmt.Errorf("error disabling the user: %v", err)
			}

			return nil
		}

		// TODO: Windows support
	default:
		return ErrUnsupportedOS
	}
}

// CmdUserMakeAdmin makes an user administrator
func (os OS) CmdUserMakeAdmin(c client.Client, usr string) error {
	switch os {
	case Linux:
		sudoersDir := "/etc/sudoers.d"

		exists, err := c.Exists(sudoersDir)
		if err != nil {
			return fmt.Errorf("error checking for the /etc/sudoers.d directory: %v", err)
		}

		if !exists {
			if err = c.MkdirAll(sudoersDir, 0755); err != nil {
				return fmt.Errorf("error creating the /etc/sudoers.d directory: %v", err)
			}

			if err = c.Chown(sudoersDir, 0, 0); err != nil { // uid 0 and gid 0 are root
				return fmt.Errorf("error changing the owner of /etc/sudoers.d: %v", err)
			}

			if err = c.Append("/etc/sudoers", []byte("#includedir /etc/sudoers.d")); err != nil {
				return fmt.Errorf("error adding the /etc/sudoers.d directory to the /etc/sudoers: %v", err)
			}
		}

		sudoFile := filepath.Join(sudoersDir, usr)
		if err = c.Append(sudoFile, []byte(fmt.Sprintf(`
Cmnd_Alias DRLM = TODO: CANVIAR AIXÒ :D
%s ALL(root) NOPASSWD: DRLM`, usr))); err != nil {
			return fmt.Errorf("error creating the user sudoer file: %v", err)
		}

		if err = c.Chown(sudoFile, 0, 0); err != nil { // uid 0 and gid 0 are root
			return fmt.Errorf("error changing the owner of the user sudoer file: %v", err)
		}

		if err = c.Chmod(sudoFile, 0440); err != nil {
			return fmt.Errorf("error changing the permissions of the user sudoer file: %v", err)
		}

		return nil

		// TODO: Windows support (and all other unixes)
	default:
		return ErrUnsupportedOS
	}
}

// CmdUserChangePasswd changes the password of an user
func (os OS) CmdUserChangePasswd(c client.Client, usr, pwd string) error {
	switch {
	case os.IsUnix():
		_, err := c.Exec(fmt.Sprintf(`echo -e "%s\n%s" | passwd %s`, pwd, pwd, usr))
		if err != nil {
			return fmt.Errorf("error changing the user password: %v", err)
		}

		return nil

		// TODO: Windows support
	default:
		return ErrUnsupportedOS
	}
}
