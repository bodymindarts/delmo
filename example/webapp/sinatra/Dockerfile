FROM alpine:3.4
RUN apk add --update ruby ruby-rdoc ruby-irb  ruby-io-console \
        && rm -rf /var/cache/apk/*
RUN gem install bundler
ADD . /opt/sinatra/
EXPOSE 5000
RUN cd /opt/sinatra && bundle install
CMD ["foreman","start","-d","/opt/sinatra"]
