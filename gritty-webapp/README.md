# What can this be used for ?

- For testing a backend api call when load balancing or reverse proxy is setup.
- For testing a REST API call written in JS.

# Command for running a reverse proxy for this web app
- `docker run --name nginx -v $(pwd):/usr/share/nginx/html:ro -v $(pwd)/default.conf:/etc/nginx/conf.d/default.conf:ro -p 80:80 --net=host -id nginx`
