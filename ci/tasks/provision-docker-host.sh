#!/bin/bash

docker-machine create \
    -d amazonec2 \
    --amazonec2-access-key ${AWS_ACCESS_KEY_ID} \
    --amazonec2-secret-key ${AWS_SERCRET_ACCESS_KEY} \
    --amazonec2-region ${AWS_REGION} \
    delmo-pipeline-machine
