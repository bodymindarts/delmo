FROM concourse/concourse-ci
RUN apt-get update && \
    apt-get install curl python-pip -y && \
    apt-get clean

# Install Go
RUN \
  mkdir -p /goroot && \
  curl https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1

# Set environment variables.
ENV GOROOT /goroot
ENV GOPATH /gopath
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

# Install docker-machine
RUN curl -L https://github.com/docker/machine/releases/download/v0.7.0/docker-machine-`uname -s`-`uname -m` > /usr/local/bin/docker-machine && \
    chmod +x /usr/local/bin/docker-machine \
    && mkdir -p /root/.docker/machine/machines
    && mkdir -p /root/.docker/machine/certs

# Install machine-share
RUN curl -sL https://deb.nodesource.com/setup_6.x | bash - \
     && apt-get install -y nodejs \
     && npm install -g machine-share

# Install aws cli
RUN curl "https://s3.amazonaws.com/aws-cli/awscli-bundle.zip" -o "awscli-bundle.zip" \
    && unzip awscli-bundle.zip \
    && ./awscli-bundle/install -i /usr/local/aws -b /usr/local/bin/aws
