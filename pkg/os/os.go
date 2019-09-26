package os

import (
	"errors"
	"fmt"
	"strings"
)

// OS is an Operative System
type OS int

const (
	// Unknown is a not known OS
	Unknown OS = iota
	// Linux is a Linux OS
	Linux
	// Windows is a Microsoft Windows OS
	Windows
	// Darwin is a macOS OS
	Darwin
	// AIX is an IBM AIX OS
	AIX
	// Dragonfly is a Dragonfly BSD OS
	Dragonfly
	// FreeBSD is a FreeBSD OS
	FreeBSD
	// NetBSD is a NetBSD OS
	NetBSD
	// OpenBSD is an OpenBSD OS
	OpenBSD
	// Plan9 is a Plan9 OS
	Plan9
	// Solaris is a Solaris OS
	Solaris
)

// ErrUnsupportedOS is an error that gets returned when the command is not supported in the OS
var ErrUnsupportedOS = errors.New("os not supported yet")

// IsUnix returns whether the OS is an Unix-like OS
func (os OS) IsUnix() bool {
	switch os {
	case Linux:
		return true
	case Darwin:
		return true
	case Dragonfly:
		return true
	case FreeBSD:
		return true
	case NetBSD:
		return true
	case OpenBSD:
		return true
	default:
		return false
	}
}

// DetectOS detects the OS the Client is running
func DetectOS(c Client) (OS, error) {
	out, err := c.Exec("uname", "-s")
	if err != nil {
		// NOT UNIX
		return Unknown, fmt.Errorf("error getting the OS: %v", err)
	}

	switch strings.TrimSpace(string(out)) {
	case "Linux":
		return Linux, nil

	case "Darwin":
		return Darwin, nil

	case "OpenBSD":
		return OpenBSD, nil

	default:
		return Unknown, nil
	}
}

// DetectVersion returns the OS version
func (os OS) DetectVersion(c Client) (string, error) {
	if os.IsUnix() {
		out, err := c.Exec("uname", "-r")
		if err != nil {
			return "", fmt.Errorf("error detecting the OS version: %v", err)
		}

		return strings.TrimSpace(string(out)), nil
	}

	switch os {
	default:
		return "unknown", nil
	}
}

// DetectDistro returns the OS distro and distro version (or the OS equivalent)
func (os OS) DetectDistro(c Client) (string, string, error) {
	distro := "unknown"
	version := "unknown"

	switch os {
	case Linux:
		out, err := c.Exec("cat /etc/os-release")
		if err != nil {
			return distro, version, fmt.Errorf("error detecting the linux distro: %v", err)
		}

		for _, l := range strings.Split(string(out), "\n") {
			if strings.HasPrefix(l, "ID=") {
				distro = strings.TrimSpace(strings.Split(l, "=")[1])
			}

			if strings.HasPrefix(l, "VERSION_ID=") {
				version = strings.TrimSpace(strings.Split(l, "=")[1])
			}
		}
	}

	return distro, version, nil
}

// cmdAsAdmin modifies a command to execute it with administrator permissions (e.g. sudo)
func (os OS) cmdAsAdmin(name string, arg ...string) (string, []string) {
	switch os {
	case Linux:
		newArg := append([]string{}, name)
		newArg = append(newArg, arg...)
		return "sudo", newArg

		// TODO: Add support for all the other things
	default:
		return name, arg
	}
}
