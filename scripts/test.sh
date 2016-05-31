#!/usr/bin/env bash

# Run the tests
echo "--> Running tests"
go list ./... | grep -v '/vendor/' | xargs -n1 go test -cover -timeout=360s
