FROM ubuntu
RUN apt-get update && apt-get install -y ruby rubygems && apt-get clean
RUN gem install bundler
ADD . /opt/sinatra/
EXPOSE 5000
RUN cd /opt/sinatra && bundle install
CMD ["/usr/local/bin/foreman","start","-d","/opt/sinatra"]
