#!/bin/sh

echo "Running test: ${TEST_NAME}, host: ${HOST_IP}"
curl -s ${HOST_IP}:5000
