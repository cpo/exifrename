#!/usr/bin/env bash

GOOS=windows GOARCH=386 go build
GOOS=linux GOARCH=amd64 go build

mv exifrename exifrename_linux_amd64