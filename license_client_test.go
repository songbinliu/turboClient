package main

import (
	"fmt"
	"github.com/golang/glog"
	"testing"
)

func Test_NewTurboRestClient(t *testing.T) {
	host := "https://localhost:9400"
	user := "administrator"
	pass := "a"

	client, err := NewRestClient(host, user, pass)
	if err != nil {
		t.Error(err)
	}

	client.Print()
}

func genExampleClient() *TurboRestClient {
	host := "https://localhost:9400"
	user := "administrator"
	pass := "a"

	client, err := NewRestClient(host, user, pass)
	if err != nil {
		glog.Errorf("failed to create client: %v", err)
		return nil
	}

	return client
}

func TestTurboRestClient_GenAddLicenseRequest(t *testing.T) {
	client := genExampleClient()
	data := []byte("<?xml version=\"1.0\"?>")

	result,err := client.genAddLicenseRequest(data)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%++v\n\n", result)
	fmt.Println(result)
}
