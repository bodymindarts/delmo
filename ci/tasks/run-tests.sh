#!/bin/bash

set -e

export GOPATH=$PWD/broker:$GOPATH
cd delmo/src/github.com/dingotiles/dingo-postgresql-broker
scripts/test.sh
