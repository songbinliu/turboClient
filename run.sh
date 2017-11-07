#!/bin/bash

go build

#host="https://localhost:9400"
host="https://10.10.174.134"
fname="./data/trial.license.xml"
password="a"

##1. test
./turboClient -v=3 --host=$host --fname=$fname --pass=$password

