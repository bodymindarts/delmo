FROM alpine:3.4

RUN apk add --update curl \
        && rm -rf /var/cache/apk/*

ADD ./webapp_is_available.sh /tasks/

CMD /bin/true
