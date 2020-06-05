// SPDX-License-Identifier: AGPL-3.0-only

package os

import (
	"fmt"
	"strings"

	"github.com/brainupdaters/drlm-common/pkg/os/client"
)

// Arch is the architecture of a CPU
type Arch int

const (
	// ArchUnknown is a not known architecture
	ArchUnknown Arch = iota
	// ArchAmd64 is the x86_64 architecture
	ArchAmd64
)

// DetectArch returns the architecture of the OS
func DetectArch(c client.Client) (Arch, error) {
	out, err := c.Exec("uname", "-m")
	if err != nil {
		// TOOD: Windows support
		return ArchUnknown, fmt.Errorf("error getting the kernel architecture: %v", err)
	}

	return ArchFromString(strings.TrimSpace(string(out))), nil
}

// ArchFromString returns an Arch based on a string
func ArchFromString(s string) Arch {
	s = strings.ToLower(s)
	switch {
	case s == "x86_64" || s == "amd64":
		return ArchAmd64

	default:
		return ArchUnknown
	}
}
