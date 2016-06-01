FROM ubuntu:14.04

RUN apt-get update && apt-get install -y curl && apt-get clean

RUN curl -L  https://github.com/coreos/etcd/releases/download/v2.3.6/etcd-v2.3.6-linux-amd64.tar.gz -o /tmp/etcd-v2.3.6-linux-amd64.tar.gz
RUN cd /tmp && tar xzvf etcd-v2.3.6-linux-amd64.tar.gz
RUN cp -f /tmp/etcd-v2.3.6-linux-amd64/etcdctl /usr/local/bin
ENTRYPOINT ["/usr/local/bin/etcdctl"]
