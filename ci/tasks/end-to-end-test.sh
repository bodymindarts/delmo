#!/bin/bash

set -x

echo "Building binary"
export GOPATH=$PWD/delmo:$GOPATH
cd delmo/src/github.com/bodymindarts/delmo
go build -o bin/delmo

echo "Downloading machine info"
aws --region ${AWS_REGION} s3 cp s3://${AWS_BUCKET}/${machine_name}.zip ./

echo "Importing ${machine_name}"
machine-import ${machine_name}.zip
# The permission isn't set properly on import
chmod 0600 /root/.docker/machine/machines/${machine_name}/id_rsa

echo "Testing example/webapp"
bin/delmo -f example/webapp/delmo.yml -m ${machine_name}
