#!/bin/bash

#go build

#host="https://localhost:9400"
#host="http://35.196.47.12:8080"
#host="https://10.10.174.134"
host="https://52.26.2.223"
fname="./data/trial.license.xml"
fname="./data/new.license.xml"
password="11zSdpNsLajv"

##1. test
./turboClient --v=4 --host=$host --fname=$fname --pass=$password

