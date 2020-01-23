// SPDX-License-Identifier: AGPL-3.0-only

package core

import (
	"crypto/x509"
	"errors"
	"fmt"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	drlm "github.com/brainupdaters/drlm-common/pkg/proto"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// NewClient returns a new client connection to the DRLM Core
func NewClient(tls bool, certPath, host string, port int) (drlm.DRLMClient, *grpc.ClientConn) {
	var grpcDialOptions = []grpc.DialOption{}

	if tls {
		cp, err := readCert(certPath)
		if err != nil {
			log.WithFields(log.Fields{
				"cert_path": certPath,
			}).Fatalf("error loading the TLS certificate: %v", err)
		}

		cred := credentials.NewClientTLSFromCert(cp, "")

		grpcDialOptions = append(grpcDialOptions, grpc.WithTransportCredentials(cred))
	} else {
		grpcDialOptions = append(grpcDialOptions, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpcDialOptions...)
	if err != nil {
		log.WithFields(log.Fields{
			"host": host,
			"port": port,
		}).Fatalf("error creating the client for DRLM Core: %v", err)
	}

	return drlm.NewDRLMClient(conn), conn
}

func readCert(certPath string) (*x509.CertPool, error) {
	b, err := afero.ReadFile(fs.FS, certPath)
	if err != nil {
		return &x509.CertPool{}, fmt.Errorf("error reading the certificate file: %v", err)
	}

	p := x509.NewCertPool()
	if ok := p.AppendCertsFromPEM(b); !ok {
		return &x509.CertPool{}, errors.New("error parsing the certificate: invalid certificate")
	}

	return p, nil
}
