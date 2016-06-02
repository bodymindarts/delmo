#!/bin/bash

echo "Creating docker-machine ${machine_name}"
docker-machine create \
    -d amazonec2 \
    --amazonec2-access-key ${AWS_ACCESS_KEY_ID} \
    --amazonec2-secret-key ${AWS_SECRET_ACCESS_KEY} \
    --amazonec2-region ${AWS_REGION} \
    ${machine_name}

echo "Exporting connection info to ${machine_name}"
machine-export ${machine_name}

echo "Uploading info to bucket"
aws --region ${AWS_REGION} s3 cp ${machine_name}.zip s3://${AWS_BUCKET}
