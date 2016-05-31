FROM ubuntu
RUN apt-get install -y ruby rubygems git
RUN gem install bundler
ADD . /opt/sinatra/
EXPOSE 5000
RUN cd /opt/sinatra && bundle install
CMD ["/usr/local/bin/foreman","start","-d","/opt/sinatra"]
