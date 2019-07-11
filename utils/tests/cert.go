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

	"github.com/spf13/afero"
)

// GenerateCert generates a new TLS certificate and stores it in the specified FS: `path/certname.key` and `path/certname.crt`
func GenerateCert(certname string, fs afero.Fs, path string) error {
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
	if err != nil {
		return fmt.Errorf("error creating the certs API request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error calling the certs API: %v", err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid HTTP response code: %d", rsp.StatusCode)
	}

	// Decode and validate the certificate
	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return fmt.Errorf("error decoding the HTTP response body: %v", err)
	}

	var certs certRsp
	if err = json.Unmarshal(b, &certs); err != nil {
		return err
	}

	if !certs.Success {
		return fmt.Errorf("unsuccessful certs generation: %v", certs.Errors)
	}

	// Store the certificate
	if err = fs.MkdirAll(path, 0755); err != nil {
		return err
	}

	if err = afero.WriteFile(fs, filepath.Join(path, certname+".key"), []byte(certs.Result.PrivateKey), 0755); err != nil {
		return err
	}

	if err = afero.WriteFile(fs, filepath.Join(path, certname+".crt"), []byte(certs.Result.Certificate), 0755); err != nil {
		return err
	}

	return nil
}

type certRsp struct {
	Success bool `json:"success"`
	Result  struct {
		Certificate string `json:"certificate"`
		PrivateKey  string `json:"private_key"`
	} `json:"result"`
	Errors []string `json:"errors"`
}
