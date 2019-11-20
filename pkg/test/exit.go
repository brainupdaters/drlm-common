// SPDX-License-Identifier: AGPL-3.0-only

package test

import (
	"os"
	"os/exec"
)

// Exits asserts that the function passed as parameter exits with an unsuccessful code
// to test functions with commands or that returns values, you can do this:
// tests.AssertExits(t func() { funcWithArgs(a, b) })
// TODO: Coverage
func (t *Test) Exits(f func()) {
	if os.Getenv("ASSERT_EXISTS_"+t.T().Name()) == "1" {
		f()
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run="+t.T().Name())
	cmd.Env = append(os.Environ(), "ASSERT_EXISTS_"+t.T().Name()+"=1")
	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}

	t.Fail("expecting unsuccessful exit")
}
