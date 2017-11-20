package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/golang/glog"
)

type Target struct {
	// Category of a probe, i.e. Hypervisor, Storage and so on.
	Category    string `json:"categroty,omitempty"`
	ClassName   string `json:"className,omitempty"`
	DisplayName string `json:"displayName,omitempty"`

	// List of field names, identifying the target of this type.
	IdentifyingFields []string `json:"identifyingFields,omitempty"`

	// List of all the account values of the target or probe.
	InputFields []*InputField `json:"inputFields,omitempty"`

	// Date of the last validation.
	LastValidated string  `json:"lastValidated,omitempty"`
	Links         []*Link `json:"links,omitempty"`

	// Description of the status.
	Status string `json:"status,omitempty"`

	// Probe type, i.ee vCenter, Hyper-V and so on.
	Type string `json:"type"`
	UUID string `json:"uuid,omitempty"`
}

func NewTarget() *Target{
	return &Target{
		Category: "Cloud Management",
		Type: "AWS",
	}
}

type awsAccount struct {
	address   string
	accesskey string
	secret    string
}

func NewAWSTarget(aws *awsAccount) *Target {
	target := NewTarget()
	inputs := []*InputField{}

	input := &InputField{
		Name: "address",
		Value: aws.address,
	}
	inputs = append(inputs, input)

	input = &InputField{
		Name: "username",
		Value: aws.accesskey,
	}
	inputs = append(inputs, input)

	input = &InputField{
		Name: "password",
		Value: aws.secret,
	}
	inputs = append(inputs, input)

	target.InputFields = inputs

	return target
}


func (c *TurboRestClient) genAddTargetRequest(target *Target) (*http.Request, error) {
	//0. data
	data, err := json.Marshal(target)
	if err != nil {
		glog.Errorf("failed to generate json: %v", err)
		return nil, err
	}

	glog.V(2).Infof("%++v", string(data))

	//1. a new http request
	urlStr := fmt.Sprintf("%s%s", c.host, API_PATH_TARGET)
	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(data))
	if err != nil {
		glog.Errorf("Failed to generate a http.request: %v", err)
		return nil, err
	}

	//2. set queries
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	//3. set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	//4. add user/password
	req.SetBasicAuth(c.username, c.password)

	return req, nil
}

func (c *TurboRestClient) AddTargetAWS(aws *awsAccount) (string, error) {

	//1. gen target
	target := NewAWSTarget(aws)

	//2. gen request
	req, err := c.genAddTargetRequest(target)
	if err != nil {
		glog.Errorf("Failed to generate AddTargetRquest: %v", err)
		return "", err
	}

	//3. send request and receive result
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

	glog.V(4).Infof("resp: %++v", resp)
	return string(result), nil
}