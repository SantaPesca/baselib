#!/bin/bash
set -e
moduleName="$(sed -n 1p go.mod | awk '{print $2}')"
rm go.mod go.sum
go mod init "$moduleName"
go mod vendor
