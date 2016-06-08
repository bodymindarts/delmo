FROM alpine:3.4

RUN apk add --update curl nodejs py-pip \
      && rm -rf /var/cache/apk

RUN pip install awscli docker-compose

# Install machine-share
RUN npm install -g machine-share

# Install docker-machine
RUN curl -L https://github.com/docker/machine/releases/download/v0.7.0/docker-machine-`uname -s`-`uname -m` > /usr/local/bin/docker-machine \
    && chmod +x /usr/local/bin/docker-machine \
    && mkdir -p /root/.docker/machine/machines \
    && mkdir -p /root/.docker/machine/certs

# Install delmo
RUN curl -L https://github.com/bodymindarts/delmo/releases/download/v0.1.1/delmo-linux-amd64 > /usr/local/bin/delmo \
      && chmod +x /usr/local/bin/delmo

ADD entrypoint.sh /
ENTRYPOINT ["/entrypoint.sh"]
CMD ["delmo", "-v"]
