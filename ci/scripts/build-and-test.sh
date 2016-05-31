#!/bin/bash

set -e -x

export GOPATH=$PWD/broker:$PWD/broker/src/github.com/dingotiles/dingo-postgresql-broker/Godeps/_workspace
export PATH=$GOPATH/bin:$PATH

cd $PWD/broker/src/github.com/dingotiles/dingo-postgresql-broker
go install github.com/onsi/ginkgo/ginkgo
ginkgo -r "$@"
