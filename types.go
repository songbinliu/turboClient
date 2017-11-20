package main

import (
	"net/url"
	"net/http"
	"crypto/tls"
	"fmt"
	"time"
	"github.com/golang/glog"
)

const (
	API_PATH_LICENSE = "/vmturbo/rest/license"
	API_PATH_TARGET = "/vmturbo/rest/targets"

	defaultTimeOut = time.Duration(60 * time.Second)
)

type TurboRestClient struct {
	client *http.Client

	host     string
	username string
	password string
}

func NewRestClient(host, uname, pass string) (*TurboRestClient, error) {

	//1. get http client
	client := &http.Client{
		Timeout: defaultTimeOut,
	}

	//2. check whether it is using ssl
	addr, err := url.Parse(host)
	if err != nil {
		glog.Errorf("Invalid url:%v, %v", host, err)
		return nil, err
	}
	if addr.Scheme == "https" {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.Transport = tr
	}

	return &TurboRestClient{
		client:   client,
		host:     host,
		username: uname,
		password: pass,
	}, nil
}

func (c *TurboRestClient) Print() {
	fmt.Printf("host: %v\n", c.host)
	fmt.Printf("client: %+++v\n", c.client)
}


type InputField struct {
	ClassName string `json:"className,omitempty"`

	// Default value of the field
	DefaultValue string `json:"defaultName,omitempty"`

	// Additional information about what the input to the field should be
	Description string `json:"description,omitempty"`
	DisplayName string `json:"displayName,omitempty"`

	// Group scope structure, filled if this field represents group scope value
	GroupProperties []*List `json:"groupProperties,omitempty"`

	// Whether the field is mandatory. Valid targets must have all the mandatory fields set.
	IsMandatory bool `json:"isMandatory,omitempty"`

	// Whether the field is secret. This means, that field value is stored in an encrypted value and not shown in any logs.
	IsSecret bool    `json:"isSecret,omitempty"`
	Links    []*Link `json:"links,omitempty"`

	// Name of the field, used for field identification.
	Name string `json:"name"`
	UUID string `json:"uuid,omitempty"`

	// Field value. Used if field holds primitive value (String, number or boolean.
	Value string `json:"value,omitempty"`

	// Type of the value this field holds = ['STRING', 'BOOLEAN', 'NUMERIC', 'GROUP_SCOPE']
	// ValueType string `json:"valueType:omitempty"`

	// The regex pattern that needs to be satisfied for the input field text
	VerificationRegex string `json:"verificationRegex,omitempty"`
}

type Link struct {
	HRef      string `json:"href,omitempty"`
	Rel       string `json:"rel,omitempty"`
	Templated bool   `json:"templated,omitempty"`
}

type List struct{}

type APIErrorDTO struct {
	ResponseType int    `json:"type"`
	Exception    string `json:"exception,omitempty"`
	Message      string `json:"message"`
}