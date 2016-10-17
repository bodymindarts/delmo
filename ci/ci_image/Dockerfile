FROM debian:wheezy

ENV DOCKER_MACHNE_VERSION=0.8.2 \
  DOCKER_COMPOSE_VERSION=1.8.1 \
  GO_VERSION=1.7.1

RUN apt-get update && \
    apt-get install -y git unzip curl python-pip && \
    apt-get clean

# Install Go
RUN \
  mkdir -p /goroot && \
  curl https://storage.googleapis.com/golang/go${GO_VERSION}.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1

# Set environment variables.
ENV GOROOT /goroot
ENV GOPATH /gopath
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH
RUN go get github.com/mitchellh/gox \
        && go get github.com/maxbrunsfeld/counterfeiter

# Install docker-machine
RUN curl -L https://github.com/docker/machine/releases/download/v${DOCKER_MACHNE_VERSION}/docker-machine-`uname -s`-`uname -m` > /usr/local/bin/docker-machine && \
    chmod +x /usr/local/bin/docker-machine \
    && mkdir -p /root/.docker/machine/machines \
    && mkdir -p /root/.docker/machine/certs

# Install docker-compose
RUN curl -L https://github.com/docker/compose/releases/download/${DOCKER_COMPOSE_VERSION}/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose \
     && chmod +x /usr/local/bin/docker-compose

# Install machine-share
RUN curl -sL https://deb.nodesource.com/setup_6.x | bash - \
     && apt-get install -y nodejs \
     && npm install -g machine-share

# Install aws cli
RUN curl "https://s3.amazonaws.com/aws-cli/awscli-bundle.zip" -o "awscli-bundle.zip" \
    && unzip awscli-bundle.zip \
    && ./awscli-bundle/install -i /usr/local/aws -b /usr/local/bin/aws
