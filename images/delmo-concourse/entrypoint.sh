#!/bin/sh

if [ ! -z ${MACHINE_NAME} ] && [ ! -z ${AWS_ACCESS_KEY_ID} ] \
    && [ ! -z ${AWS_SECRET_ACCESS_KEY} ] &&  [ ! -z ${AWS_DEFAULT_REGION} ] \
    && [ ! -z ${AWS_BUCKET} ]; then
    echo "Downloading pre existing docker-machine configuration..."
    aws --region ${AWS_DEFAULT_REGION} s3 cp s3://${AWS_BUCKET}/${MACHINE_NAME}.zip ./ || exit 1

    echo "Importing configuration..."
    machine-import ${MACHINE_NAME}.zip
    # The permission isn't set properly on import
    chmod 0600 /root/.docker/machine/machines/${MACHINE_NAME}/id_rsa

    echo "Machine ${MACHINE_NAME} imported!"
fi

echo "Running '$@'"
exec $@
