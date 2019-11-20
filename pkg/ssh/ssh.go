// SPDX-License-Identifier: AGPL-3.0-only

package ssh

import (
	"net"

	stdSSH "golang.org/x/crypto/ssh"
)

// MultipleFixedHostKeys is a HostKeyCallback that checks multiple host keys
func MultipleFixedHostKeys(keys []stdSSH.PublicKey) stdSSH.HostKeyCallback {
	hk := &multipleFixedHostKeys{keys}
	return hk.check
}

type multipleFixedHostKeys struct {
	keys []stdSSH.PublicKey
}

func (m *multipleFixedHostKeys) check(hostname string, remote net.Addr, key stdSSH.PublicKey) error {
	var err error
	for _, k := range m.keys {
		err = stdSSH.FixedHostKey(k)(hostname, remote, key)
		if err == nil {
			return nil
		}
	}

	return err
}
