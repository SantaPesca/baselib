#!/bin/bash
set -e
moduleName="$(echo $(sed -n 1p go.mod | sed -e 's/\<module\>//g'))"
rm go.mod go.sum
go mod init $moduleName
