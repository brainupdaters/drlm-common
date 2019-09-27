package os

import (
	"fmt"
	"path/filepath"
	"strings"
)

// CmdUserCreate creates a new user in the OS
func (os OS) CmdUserCreate(c Client, name, pwd string) error {
	switch {
	case os.IsUnix():
		out, err := c.Exec("echo", fmt.Sprintf(`"%q"`, pwd), "|", "openssl", "passwd", "-1", "-stdin")
		if err != nil {
			return fmt.Errorf("error encrypting the password: %v", err)
		}

		cryptPwd := strings.TrimSpace(string(out))
		_, err = c.ExecAdmin(os, "useradd", "-m", "-c", `"DRLM Agent user"`, "-p", fmt.Sprintf(`"%s"`, cryptPwd), fmt.Sprintf(`"%s"`, name))
		if err != nil {
			return fmt.Errorf("error creating the user: %v", err)
		}

		return nil

	default:
		return ErrUnsupportedOS
	}
}

// CmdUserGroup returns the group of an user
func (os OS) CmdUserGroup(c Client, usr string) (string, error) {
	switch {
	case os.IsUnix():
		out, err := c.Exec("id", "-gn", usr)
		if err != nil {
			return "", fmt.Errorf("error getting the user group: %v", err)
		}

		return strings.TrimSpace(string(out)), nil

		// TODO: Windows support
	default:
		return "", ErrUnsupportedOS
	}
}

// CmdUserDisable disables an user
func (os OS) CmdUserDisable(c Client, usr string) error {
	switch {
	case os.IsUnix():
		_, err := c.ExecAdmin(os, "passwd", "-l", usr)
		if err != nil {
			return fmt.Errorf("error disabling the user: %v", err)
		}

		return nil

		// TODO: Windows support
	default:
		return ErrUnsupportedOS
	}
}

// CmdUserMakeAdmin makes an user administrator
func (os OS) CmdUserMakeAdmin(c Client, usr string) error {
	switch os {
	case Linux:
		exists, err := os.CmdFSCheckDir(c, "/etc/sudoers.d")
		if err != nil {
			return fmt.Errorf("error checking for the /etc/sudoers.d directory: %v", err)
		}

		if !exists {
			err = os.CmdFSMkdir(c, "/etc/sudoers.d")
			if err != nil {
				return fmt.Errorf("error creating the /etc/sudoers.d directory: %v", err)
			}

			err = os.CmdFSChown(c, "/etc/sudoers.d", "root", "root")
			if err != nil {
				return fmt.Errorf("error changing the owner of /etc/sudoers.d: %v", err)
			}

			err = os.CmdFSChmod(c, "/etc/sudoers.d", "0755")
			if err != nil {
				return fmt.Errorf("error changing the permissions of /etc/sudoers.d: %v", err)
			}

			err = os.CmdFSAppendToFileAdmin(c, "/etc/sudoers", []byte("#includedir /etc/sudoers.d"))
			if err != nil {
				return fmt.Errorf("error adding the /etc/sudoers.d directory to the /etc/sudoers: %v", err)
			}
		}

		sudoFile := filepath.Join(os.CmdFSTempDir(), "sudoers_"+usr)
		err = os.CmdFSAppendToFile(c, sudoFile, []byte(fmt.Sprintf(`Cmnd_Alias DRLM= TODO: CANVIAR AIXÒ :D
%s	ALL=(root)	NOPASSWD: DRLM`, usr)))
		if err != nil {
			return fmt.Errorf("error creating the user sudoer file: %v", err)
		}

		err = os.CmdFSChown(c, sudoFile, "root", "root")
		if err != nil {
			return fmt.Errorf("error changing the owner of the user sudoer file: %v", err)
		}

		err = os.CmdFSChmod(c, sudoFile, "0440")
		if err != nil {
			return fmt.Errorf("error changing the permissions of the user sudoer file: %v", err)
		}

		err = os.CmdFSMove(c, sudoFile, fmt.Sprintf("/etc/sudoers.d/%s", usr))
		if err != nil {
			return fmt.Errorf("error moving the user sudoer file to /etc/sudoers.d: %v", err)
		}

		return nil

		// TODO: Windows support (and all other unixes)
	default:
		return ErrUnsupportedOS
	}
}
