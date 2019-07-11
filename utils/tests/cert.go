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

package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

// GenerateCert generates a new TLS certificate and stores it in the specified FS: `path/certname.key` and `path/certname.crt`
func GenerateCert(t *testing.T, fs afero.Fs, certname, path string) {
	assert := assert.New(t)

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
	assert.Nil(err)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rsp, err := http.DefaultClient.Do(req)
	assert.Nil(err)
	defer rsp.Body.Close()

	assert.Equal(http.StatusOK, rsp.StatusCode)

	// Decode and validate the certificate
	b, err := ioutil.ReadAll(rsp.Body)
	assert.Nil(err)

	var certs certRsp
	assert.Nil(json.Unmarshal(b, &certs))

	assert.True(certs.Success)

	// Store the certificate
	assert.Nil(fs.MkdirAll(path, 0755))
	assert.Nil(afero.WriteFile(fs, filepath.Join(path, certname+".key"), []byte(certs.Result.PrivateKey), 0755))
	assert.Nil(afero.WriteFile(fs, filepath.Join(path, certname+".crt"), []byte(certs.Result.Certificate), 0755))
}

type certRsp struct {
	Success bool `json:"success"`
	Result  struct {
		Certificate string `json:"certificate"`
		PrivateKey  string `json:"private_key"`
	} `json:"result"`
	Errors []string `json:"errors"`
}
