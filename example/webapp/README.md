# Webapp Example

This is a complete example of how `delmo` could be used.
When in this directory you can run it via `delmo` or `delmo --parallel`.

[docker-compose.yml](./docker-compose.yml) defines the services that make up this system.

The `proxy` service consits of nginx [configured](./nginx/nginx.conf) to load balance between 2 webapp instances .

The `webapp` services run a [minimal sinatra server](./sinatra/server.rb).

The `tests` service is used by delmo to run tasks. In this example the tasks are scripts copied into the [image](./tests/Dockerfile).

During the [tests](./delmo.yml) individual services are started and stopped while [webapp_is_available.sh](./tests/webapp_is_available.sh) is used to assert that valid responses are still being transmitted.
