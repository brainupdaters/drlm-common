/*
 * Copyright (C) 2019 DRLM Project
 * Authors: DRLM Common authors
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

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
