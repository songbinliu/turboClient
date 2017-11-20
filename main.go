package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
)

var (
	fname    string
	host     string
	user     string
	password string
)

func genMyClient() (*TurboRestClient, error) {
	host := "https://localhost:9400"
	uname := "administrator"
	pass := "a"

	return NewRestClient(host, uname, pass)
}

func addLicense() {
	client, _ := NewRestClient(host, user, password)
	result, _ := client.AddLicense(fname)
	glog.V(2).Infof("result=%+v", result)
}

func parseFlags() {
	flag.Set("logtostderr", "true")
	flag.StringVar(&fname, "fname", "./data/license.xml", "the xml license file.")
	flag.StringVar(&host, "host", "https://localhost:9400", "the address of turbo.server")
	flag.StringVar(&user, "user", "administrator", "username to login to turbo.server")
	flag.StringVar(&password, "pass", "a", "password to login to turbo.server")
	flag.Parse()

	fmt.Printf("turbo.server: %v\n", host)
}

func addAWSTarget() {
	aws := &awsAccount{
		address:   "my.aws.amazon.com",
		accesskey: "abdcedef",
		secret:    "V+abcedef",
	}

	client, _ := NewRestClient(host, user, password)
	result, _ := client.AddTargetAWS(aws)
	glog.V(2).Infof("result=%+v", result)
}

func main() {
	parseFlags()
	//addLicense()
	addAWSTarget()
}
