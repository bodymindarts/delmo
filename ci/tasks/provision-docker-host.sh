#!/bin/bash

set -e
set -x

cat >> ~/.aws/credentials <<EOF
[default]
aws_access_key_id = ${AWS_SERCRET_ACCESS_KEY}
aws_secret_access_key = ${AWS_SERCRET_ACCESS_KEY}
EOF

docker-machine create \
    -d amazonec2 \
    --amazonec2-region ${AWS_REGION} \
    ${machine_name}

machine-export ${machine_name}
aws s3 cp ${machine_name}.zip s3://${AWS_BUCKET}
