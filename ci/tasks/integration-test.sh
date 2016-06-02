#!/bin/bash

export GOPATH=$PWD/delmo:$GOPATH
cd delmo/src/github.com/bodymindarts/delmo

echo "Building binary"
go build -o bin/delmo

echo "Downloading machine info"
aws --region ${AWS_REGION} s3 cp s3://${AWS_BUCKET}/${machine_name}.zip ${machine_name}.zip

echo "Importing ${machine_name}"
machine-import ${machine_name}.zip

echo "Evaluating environment"
eval $(docker-machine env ${machine_name})

bin/delmo -f example/webapp/delmo.yml
