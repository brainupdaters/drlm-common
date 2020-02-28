// SPDX-License-Identifier: AGPL-3.0-only

package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

// GenerateCert generates a new TLS certificate and stores it into fs.FS: `path/certname.key` and `path/certname.crt`
func (t *Test) GenerateCert(fs afero.Fs, certname, path string) {
	// Request the certificate to the cfssl certs API
	body := strings.NewReader(fmt.Sprintf(`{
		"request": {
			"CN": "%s",
			"hosts": [ "%s" ],
			"key": {
				"algo": "rsa",
				"size": 2048
			},
			"names": [{
				"O": "Brain Updaters"
			}]
		}
	}`, certname, certname))
	req, err := http.NewRequest("POST", "http://tls:8888/api/v1/cfssl/newcert", body)
	t.Require().Nil(err)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rsp, err := http.DefaultClient.Do(req)
	t.Require().Nil(err)
	t.Require().Equal(http.StatusOK, rsp.StatusCode)
	defer rsp.Body.Close()

	// Decode and validate the certificate
	b, err := ioutil.ReadAll(rsp.Body)
	t.Require().Nil(err)

	var certs certRsp
	t.Require().Nil(json.Unmarshal(b, &certs))
	t.Require().True(certs.Success)

	// Store the certificate
	t.Require().Nil(fs.MkdirAll(path, 0755))
	t.Require().Nil(afero.WriteFile(fs, filepath.Join(path, certname+".key"), []byte(certs.Result.PrivateKey), 0755))
	t.Require().Nil(afero.WriteFile(fs, filepath.Join(path, certname+".crt"), []byte(certs.Result.Certificate), 0755))
}

type certRsp struct {
	Success bool `json:"success"`
	Result  struct {
		Certificate string `json:"certificate"`
		PrivateKey  string `json:"private_key"`
	} `json:"result"`
	Errors []string `json:"errors"`
}
