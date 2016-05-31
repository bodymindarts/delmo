#!/bin/bash

set -e

export GOPATH=$PWD/delmo:$GOPATH
cd delmo/src/github.com/bodymindarts/delmo
scripts/test.sh
