#!/bin/bash

set -e
set -x

docker-machine create \
    -d amazonec2 \
    --amazonec2-access-key ${AWS_ACCESS_KEY_ID} \
    --amazonec2-secret-key ${AWS_SECRET_ACCESS_KEY} \
    --amazonec2-region ${AWS_REGION} \
    ${machine_name}

machine-export ${machine_name}
aws s3 cp ${machine_name}.zip s3://${AWS_BUCKET}
