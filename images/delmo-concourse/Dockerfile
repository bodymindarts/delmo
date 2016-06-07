FROM alpine:3.4

RUN apk add --update curl nodejs py-pip \
      && rm -rf /var/cache/apk

RUN curl -L https://github.com/docker/machine/releases/download/v0.7.0/docker-machine-`uname -s`-`uname -m` > /usr/local/bin/docker-machine \
    && chmod +x /usr/local/bin/docker-machine \
    && mkdir -p /root/.docker/machine/machines \
    && mkdir -p /root/.docker/machine/certs

# Install machine-share
RUN npm install -g machine-share

# Install aws cli
RUN curl "https://s3.amazonaws.com/aws-cli/awscli-bundle.zip" -o "awscli-bundle.zip" \
    && unzip awscli-bundle.zip \
    && ./awscli-bundle/install -i /usr/local/aws -b /usr/local/bin/aws

RUN curl -L https://github.com/bodymindarts/delmo/releases/download/v0.0.3/delmo-linux-amd64 > /usr/local/bin/delmo \
      && chmod +x /usr/local/bin/delmo

ADD entrypoint.sh /
ENTRYPOINT ["/entrypoint.sh"]
CMD ["delmo", "-v"]
