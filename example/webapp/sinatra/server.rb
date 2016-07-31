require 'sinatra'

set :bind, '0.0.0.0'

set :logging, true

get '/' do
  'hello world from docker!'
end
