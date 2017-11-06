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
