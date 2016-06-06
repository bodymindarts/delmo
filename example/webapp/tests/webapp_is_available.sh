#!/bin/sh

echo "Running test: ${TEST_NAME}, host: ${HOST_IP}"
curl ${HOST_IP}:5000
