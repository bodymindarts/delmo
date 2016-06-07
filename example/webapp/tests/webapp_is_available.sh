#!/bin/sh

echo "Running test: ${DELMO_TEST_NAME}, host: ${DOCKER_HOST_IP}"
curl -s ${DOCKER_HOST_IP}:5000
