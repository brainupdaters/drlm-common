// SPDX-License-Identifier: AGPL-3.0-only

package minio

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"

	sdk "github.com/minio/minio-go/v6"
	"github.com/minio/minio/pkg/madmin"
	"github.com/spf13/afero"
)

// conn returns the parameters for the minio connections
func conn(host string, port int, accessKey, secretKey string, ssl bool) (string, string, string, bool) {
	return fmt.Sprintf("%s:%d", host, port),
		accessKey,
		secretKey,
		ssl
}

// transport returns the http transport for the minio connections
func transport(fs afero.Fs, tr *http.Transport, certPath string) error {
	b, err := afero.ReadFile(fs, certPath)
	if err != nil {
		return fmt.Errorf("error creating the minio http transport: error reading the certificate: %v", err)
	}

	if tr.TLSClientConfig == nil {
		// Taken from minio/minio-go
		// Keep TLS config.
		tlsConfig := &tls.Config{
			// Can't use SSLv3 because of POODLE and BEAST
			// Can't use TLSv1.0 because of POODLE and BEAST using CBC cipher
			// Can't use TLSv1.1 because of RC4 cipher usage
			MinVersion: tls.VersionTLS12,
		}
		tr.TLSClientConfig = tlsConfig
	}

	if tr.TLSClientConfig.RootCAs == nil {
		// Taken from minio/minio-go
		rootCAs, _ := x509.SystemCertPool()
		if rootCAs == nil {
			// In some systems (like Windows) system cert pool is
			// not supported or no certificates are present on the
			// system - so we create a new cert pool.
			rootCAs = x509.NewCertPool()
		}
		tr.TLSClientConfig.RootCAs = rootCAs
	}

	if ok := tr.TLSClientConfig.RootCAs.AppendCertsFromPEM(b); !ok {
		return fmt.Errorf("error creating the minio http transport: error parsing the certificate")
	}

	return nil
}

// NewSDK returns a Minio SDK
func NewSDK(fs afero.Fs, host string, port int, accessKey, secretKey string, ssl bool, certPath string) (*sdk.Client, error) {
	minio, err := sdk.New(conn(host, port, accessKey, secretKey, ssl))
	if err != nil {
		return minio, fmt.Errorf("error creating the connection to minio: %v", err)
	}

	// If the certificate is self signed, add it to the transport certificates pool
	if ssl && certPath != "" {
		defaultTransport, err := sdk.DefaultTransport(true)
		if err != nil {
			return minio, fmt.Errorf("error creating the minio connection: error creating the default transport layer: %v", err)
		}

		tr := defaultTransport.(*http.Transport)
		if err = transport(fs, tr, certPath); err != nil {
			return minio, fmt.Errorf("error creating the minio connection: %v", err)
		}

		minio.SetCustomTransport(tr)
	}

	return minio, nil
}

// NewAdminClient returns a Minio Admin client
func NewAdminClient(fs afero.Fs, host string, port int, accessKey, secretKey string, ssl bool, certPath string) (*madmin.AdminClient, error) {
	cli, err := madmin.New(conn(host, port, accessKey, secretKey, ssl))
	if err != nil {
		return nil, fmt.Errorf("error creating the minio admin connection: %v", err)
	}
	// If the certificate is self signed, add it to the transport certificates pool
	if ssl && certPath != "" {
		tr := http.DefaultTransport.(*http.Transport)
		if err := transport(fs, tr, certPath); err != nil {
			return nil, fmt.Errorf("error creating the minio admin connection: %v", err)
		}

		cli.SetCustomTransport(tr)
	}

	return cli, nil
}
