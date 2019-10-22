package os

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	"github.com/brainupdaters/drlm-common/pkg/os/client"

	"github.com/spf13/afero"
	"golang.org/x/crypto/ssh"
)

// CmdSSHGetHostKeys returns the public SSH keys of a host
func (os OS) CmdSSHGetHostKeys(c client.Client, host string, port int) ([]string, error) {
	switch {
	case os.IsUnix():
		out, err := c.Exec("ssh-keyscan", "-p", strconv.Itoa(port), host)
		if err != nil {
			return []string{}, fmt.Errorf("error getting the host SSH keys: %v", err)
		}

		var keys []string
		for _, l := range strings.Split(strings.TrimSpace(string(out)), "\n") {
			if !strings.HasPrefix(l, "# ") {
				keys = append(keys, l)
			}
		}

		return keys, nil

	default:
		return []string{}, ErrUnsupportedOS
	}
}

// CmdSSHCopyID copies the key to the OS
func (os OS) CmdSSHCopyID(c client.Client, usr string, key []byte) error {
	switch {
	case os.IsUnix():
		home, err := os.CmdFSHome(c, usr)
		if err != nil {
			return fmt.Errorf("error copying the SSH key: %v", err)
		}

		uid, err := os.CmdUserUID(c, usr)
		if err != nil {
			return fmt.Errorf("error copying the SSH key: %v", err)
		}

		gid, err := os.CmdUserGID(c, usr)
		if err != nil {
			return fmt.Errorf("error copying the SSH key: %v", err)
		}

		sshDir := filepath.Join(home, ".ssh")
		exists, err := c.Exists(sshDir)
		if err != nil {
			return fmt.Errorf("error checking the SSH directory: %v", err)
		}

		// If the SSH directory doesn't exist, create it
		if !exists {
			if err = c.Mkdir(sshDir, 0700); err != nil {
				return fmt.Errorf("error creating the SSH directory: %v", err)
			}

			if err = c.Chown(sshDir, uid, gid); err != nil {
				return fmt.Errorf("error changing the SSH directory owner: %v", err)
			}
		}

		authKeys := filepath.Join(sshDir, "authorized_keys")
		tmpAuthKeys := filepath.Join(os.CmdFSTempDir(), "drlm_core_authorized_keys")
		exists, err = c.Exists(authKeys)
		if err != nil {
			return fmt.Errorf("error checking for the authorized_keys file: %v", err)
		}

		if exists {
			if err = c.Copy(authKeys, tmpAuthKeys); err != nil {
				return fmt.Errorf("error copying the authorized_keys file: %v", err)
			}
		}

		if err = c.Append(tmpAuthKeys, key); err != nil {
			return fmt.Errorf("error adding the key to the authorized_keys file: %v", err)
		}

		if err = c.Chown(tmpAuthKeys, uid, gid); err != nil {
			return fmt.Errorf("error changing the authorized_keys owner: %v", err)
		}

		if err = c.Chmod(tmpAuthKeys, 0600); err != nil {
			return fmt.Errorf("error changing the authorized_keys permissions: %v", err)
		}

		if err = c.Move(tmpAuthKeys, authKeys); err != nil {
			return fmt.Errorf("error replacing the authorized_keys file: %v", err)
		}

		return nil

	default:
		return ErrUnsupportedOS
	}
}

// CmdSSHGetKeysPath returns the SSH keys directory
func (os OS) CmdSSHGetKeysPath(c client.Client, usr string) (string, error) {
	switch {
	case os.IsUnix():
		home, err := os.CmdFSHome(c, usr)
		if err != nil {
			return "", fmt.Errorf("error getting the public key: %v", err)
		}

		return filepath.Join(home, ".ssh"), nil

	default:
		return "", ErrUnsupportedOS
	}
}

// CmdSSHGenerateKeyPair generates an RSA 4096 key pair and saves them to the path as `id_rsa` and `id_rsa.pub`
func (os OS) CmdSSHGenerateKeyPair(c client.Client, path string) error {
	pK, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return fmt.Errorf("error generating the private key: %v", err)
	}

	if err := pK.Validate(); err != nil {
		return fmt.Errorf("invalid private key generated: %v", err)
	}

	k, err := ssh.NewPublicKey(&pK.PublicKey)
	if err != nil {
		return fmt.Errorf("error generating the public key: %v", err)
	}

	kB := ssh.MarshalAuthorizedKey(k)

	pDER := x509.MarshalPKCS1PrivateKey(pK)
	pBlock := &pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   pDER,
	}
	pKB := pem.EncodeToMemory(pBlock)

	if err := afero.WriteFile(fs.FS, filepath.Join(path, "id_rsa"), pKB, 0400); err != nil {
		return fmt.Errorf("error writting the private key: %v", err)
	}

	if err := afero.WriteFile(fs.FS, filepath.Join(path, "id_rsa.pub"), kB, 0400); err != nil {
		return fmt.Errorf("error writting the public key: %v", err)
	}

	return nil
}
