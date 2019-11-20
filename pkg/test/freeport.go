// SPDX-License-Identifier: AGPL-3.0-only

package test

import "net"

// FreePort returns a free port to listen to
// https://github.com/phayes/freeport
func (t *Test) FreePort() int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	t.Require().Nil(err)

	l, err := net.ListenTCP("tcp", addr)
	defer l.Close()
	t.Require().Nil(err)

	return l.Addr().(*net.TCPAddr).Port
}
