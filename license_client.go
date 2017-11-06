package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
)

type TurboLicense struct {
	License string `json:"license"`
}

// read license from xml file, and generate encoded json content
func getLicenseData(fname string) ([]byte, error) {
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		glog.Errorf("Failed to read file %v: %v", fname, err)
		return []byte{}, err
	}
	glog.V(3).Infof("data: %++v", string(content))

	license := &TurboLicense{
		License: string(content),
	}
	data, err := json.Marshal(license)
	if err != nil {
		glog.Errorf("failed to generate json: %v", err)
		return []byte{}, err
	}
	glog.V(3).Infof("json: %++v", string(data))

	return data, nil
}

func (c *TurboRestClient) genAddLicenseRequest(data []byte) (*http.Request, error) {
	//1. a new http request
	urlStr := fmt.Sprintf("%s%s", c.host, API_PATH_LICENSE)
	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(data))
	if err != nil {
		glog.Errorf("Failed to generate a http.request: %v", err)
		return nil, err
	}

	//2. set queries
	q := req.URL.Query()
	q.Add("validate", "false")
	req.URL.RawQuery = q.Encode()

	//3. set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	//4. add user/password
	req.SetBasicAuth(c.username, c.password)

	return req, nil
}

func (c *TurboRestClient) AddLicense(fname string) (string, error) {
	//1. get json content
	data, err := getLicenseData(fname)
	if err != nil {
		glog.Errorf("failed to generate json: %v", err)
		return "", err
	}

	//2. httpRequest
	req, err := c.genAddLicenseRequest(data)
	if err != nil {
		glog.Errorf("Failed to generate request: %v", err)
		return "", err
	}

	//3. send request and get result
	resp, err := c.client.Do(req)
	if err != nil {
		glog.Errorf("Failed to send request: %v", err)
		return "", err
	}

	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Errorf("Failed to read response: %v", err)
		return "", err
	}

	glog.V(3).Infof("resp: %++v", resp)
	return string(result), nil
}
