package os

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/brainupdaters/drlm-common/pkg/ssh"
)

// Client is the responsible for executing commands in a OS. It can be the local OS or an OS connected through SSH
type Client interface {
	Exec(name string, arg ...string) ([]byte, error)
	ExecAdmin(os OS, name string, arg ...string) ([]byte, error)
}

// ClientLocal is the OS where the server is getting executed
type ClientLocal struct {
	IsAdmin bool
}

// Exec executes a command to the server OS
func (c *ClientLocal) Exec(name string, arg ...string) ([]byte, error) {
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

// ExecAdmin executes a command that requires Admin privilages
func (c *ClientLocal) ExecAdmin(os OS, name string, arg ...string) ([]byte, error) {
	if c.IsAdmin {
		return c.Exec(name, arg...)
	}

	name, arg = os.cmdAsAdmin(name, arg...)
	return c.Exec(name, arg...)
}

// ClientSSH is an OS client using SSH
type ClientSSH struct {
	Session *ssh.Session
	IsAdmin bool
}

// Exec executes a command through SSH
func (c *ClientSSH) Exec(name string, arg ...string) ([]byte, error) {
	cmd := append([]string{name}, arg...)
	return c.Session.Exec(strings.Join(cmd, " "))
}

// ExecAdmin executes a command that requires Admin privilages
func (c *ClientSSH) ExecAdmin(os OS, name string, arg ...string) ([]byte, error) {
	if c.IsAdmin {
		return c.Exec(name, arg...)
	}

	name, arg = os.cmdAsAdmin(name, arg...)
	return c.Exec(name, arg...)
}
