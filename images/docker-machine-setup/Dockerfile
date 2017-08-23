FROM alpine:3.4

ENV DOCKER_MACHNE_VERSION=0.12.2

RUN apk add --update curl nodejs py-pip \
      && rm -rf /var/cache/apk/*

RUN curl -L https://github.com/docker/machine/releases/download/v${DOCKER_MACHNE_VERSION}/docker-machine-`uname -s`-`uname -m` > /usr/local/bin/docker-machine \
    && chmod +x /usr/local/bin/docker-machine \
    && mkdir -p /root/.docker/machine/machines \
    && mkdir -p /root/.docker/machine/certs
#
# Install machine-share
RUN npm install -g machine-share

# Install aws cli
RUN curl "https://s3.amazonaws.com/aws-cli/awscli-bundle.zip" -o "awscli-bundle.zip" \
    && unzip awscli-bundle.zip \
    && ./awscli-bundle/install -i /usr/local/aws -b /usr/local/bin/aws

ADD ./delete-machine /delete-machine
ADD ./setup-aws /setup-aws
ADD ./setup-digitalocean /setup-digitalocean

CMD /setup-aws

