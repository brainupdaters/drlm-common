package client

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/brainupdaters/drlm-common/pkg/ssh"
)

// SSH is an OS client using SSH
type SSH struct {
	Session *ssh.Session
}

// Exec executes a command through SSH
func (c *SSH) Exec(name string, arg ...string) ([]byte, error) {
	cmd := append([]string{name}, arg...)
	return c.Session.Exec(strings.Join(cmd, " "))
}

// Chmod changes the mode of the named file to mode
func (c *SSH) Chmod(path string, mode os.FileMode) error {
	if err := c.Session.SFTP.Chmod(path, mode); err != nil {
		return fmt.Errorf("error changing the mode of the file: %v", err)
	}

	return nil
}

// Chown changes the numeric uid and gid of the named file.
func (c *SSH) Chown(path string, uid, gid int) error {
	if err := c.Session.SFTP.Chown(path, uid, gid); err != nil {
		return fmt.Errorf("error changing the owners of the file: %v", err)
	}

	return nil
}

// Exists checks if a file exists or not
// TODO: Check that the error returned is os.IsNotExist()
func (c *SSH) Exists(path string) (bool, error) {
	_, err := c.Session.SFTP.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return true, fmt.Errorf("error checking the file existence: %v", err)
	}

	return true, nil
}

// Mkdir creates a new directory with the specified name and permission bits (before umask)
func (c *SSH) Mkdir(path string, perm os.FileMode) error {
	if err := c.Session.SFTP.Mkdir(path); err != nil {
		return fmt.Errorf("error creating the directory: %v", err)
	}

	if err := c.Chmod(path, perm); err != nil {
		return fmt.Errorf("error changing the directory permissions: %v", err)
	}

	return nil
}

// Write writes content to a file
func (c *SSH) Write(path string, b []byte) error {
	f, err := c.Session.SFTP.Create(path)
	if err != nil {
		return fmt.Errorf("error creating the file: %v", err)
	}
	defer f.Close()

	if _, err := f.Write(b); err != nil {
		return fmt.Errorf("error writting the file: %v", err)
	}

	return nil
}

// Append appends content to a file
func (c *SSH) Append(path string, b []byte) error {
	f, err := c.Session.SFTP.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY)
	if err != nil {
		return fmt.Errorf("error opening the file: %v", err)
	}
	defer f.Close()

	if _, err := f.Write(b); err != nil {
		return fmt.Errorf("error writting the file: %v", err)
	}

	return nil
}

// ReadFile returns the content of a file
func (c *SSH) ReadFile(path string) ([]byte, error) {
	f, err := c.Session.SFTP.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening the file: %v", err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error reading the file: %v", err)
	}

	return b, nil
}

// Copy copies from a source to a destination. It's recursive, tries to preserve permissions and skips symlinks
// Copied from: https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
func (c *SSH) Copy(src, dst string) error {
	sF, err := c.Session.SFTP.Stat(src)
	if err != nil {
		return fmt.Errorf("error checking the file properties: %v", err)
	}

	if _, err := c.Session.SFTP.Stat(dst); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error checking the destination properties: %v", err)
	} else if err == nil {
		return errors.New("error copying the file: the destination already exists")
	}

	switch sF.IsDir() {
	case true:
		if err := c.copyDir(src, dst); err != nil {
			return fmt.Errorf("error copying the directory: %v", err)
		}

		return nil

	default:
		if err := c.copyFile(src, dst); err != nil {
			return fmt.Errorf("error copying the file: %v", err)
		}

		return nil
	}
}

// Copied from: https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
func (c *SSH) copyDir(src, dst string) error {
	sF, err := c.Session.SFTP.Stat(src)
	if err != nil {
		return fmt.Errorf("error checking the file properties: %v", err)
	}

	if err := c.Session.SFTP.Mkdir(dst); err != nil {
		return fmt.Errorf("error creating the directory: %v", err)
	}

	if err := c.Session.SFTP.Chmod(dst, sF.Mode()); err != nil {
		return fmt.Errorf("error changing the directory permissions: %v", err)
	}

	entries, err := c.Session.SFTP.ReadDir(src)
	if err != nil {
		return fmt.Errorf("error checking the directory items: %v", err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := c.copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			if err := c.copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

// Copied from: https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
func (c *SSH) copyFile(src, dst string) error {
	sF, err := c.Session.SFTP.Open(src)
	if err != nil {
		return fmt.Errorf("error opening the file: %v", err)
	}
	defer sF.Close()

	dF, err := c.Session.SFTP.Create(dst)
	if err != nil {
		return fmt.Errorf("error creating the file: %v", err)
	}

	if _, err := io.Copy(dF, sF); err != nil {
		return fmt.Errorf("error copying the file contents: %v", err)
	}

	sFStat, err := c.Session.SFTP.Stat(src)
	if err != nil {
		return fmt.Errorf("error checking the file properties: %v", err)
	}

	if err := c.Chmod(dst, sFStat.Mode()); err != nil {
		return fmt.Errorf("error changing the file permissions: %v", err)
	}

	if err := dF.Close(); err != nil {
		return fmt.Errorf("error closing the new file: %v", err)
	}

	return nil
}

// Move moves a file from a source to a destination
func (c *SSH) Move(src, dst string) error {
	if err := c.Session.SFTP.Rename(src, dst); err != nil {
		return fmt.Errorf("error moving the file: %v", err)
	}

	return nil
}
