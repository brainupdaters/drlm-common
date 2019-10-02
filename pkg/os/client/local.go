package client

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// Local is the OS where the server is getting executed
type Local struct{}

// Exec executes a command to the server OS
func (c *Local) Exec(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return stdout.Bytes(), fmt.Errorf("%v: %s", err, stderr.String())
	}

	return stdout.Bytes(), nil
}

// Chmod changes the mode of the named file to mode
func (c *Local) Chmod(path string, mode os.FileMode) error {
	if err := os.Chmod(path, mode); err != nil {
		return fmt.Errorf("error changing the mode of the file: %v", err)
	}

	return nil
}

// Chown changes the numeric uid and gid of the named file.
func (c *Local) Chown(path string, uid, gid int) error {
	if err := os.Chown(path, uid, gid); err != nil {
		return fmt.Errorf("error changing the owners of the file: %v", err)
	}

	return nil
}

// Exists checks if a file exists or not
func (c *Local) Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return true, fmt.Errorf("error checking the file existence: %v", err)
	}

	return true, nil
}

// Mkdir creates a new directory with the specified name and permission bits (before umask)
func (c *Local) Mkdir(path string, perm os.FileMode) error {
	if err := os.Mkdir(path, perm); err != nil {
		return fmt.Errorf("error creating the directory: %v", err)
	}

	return nil
}

// Write writes content to a file
func (c *Local) Write(path string, b []byte) error {
	f, err := os.Create(path)
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
func (c *Local) Append(path string, b []byte) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
func (c *Local) ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
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
func (c *Local) Copy(src, dst string) error {
	sF, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("error checking the file properties: %v", err)
	}

	if _, err := os.Stat(dst); err != nil && !os.IsNotExist(err) {
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
func (c *Local) copyDir(src, dst string) error {
	sF, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("error checking the file properties: %v", err)
	}

	if err := os.Mkdir(dst, sF.Mode()); err != nil {
		return fmt.Errorf("error creating the directory: %v", err)
	}

	entries, err := ioutil.ReadDir(src)
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
func (c *Local) copyFile(src, dst string) error {
	sF, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("error opening the file: %v", err)
	}
	defer sF.Close()

	dF, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("error creating the file: %v", err)
	}

	if _, err := io.Copy(dF, sF); err != nil {
		return fmt.Errorf("error copying the file contents: %v", err)
	}

	if err := dF.Sync(); err != nil {
		return fmt.Errorf("error finishing the file content sync: %v", err)
	}

	sFStat, err := os.Stat(src)
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
func (c *Local) Move(src, dst string) error {
	if err := os.Rename(src, dst); err != nil {
		return fmt.Errorf("error moving the file: %v", err)
	}

	return nil
}
