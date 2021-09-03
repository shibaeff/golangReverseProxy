 # Simple reverse proxy
 A simple reverse proxy implementation in golang. Proxy server listening on port 9000 is able to redirect 
 incoming requests to the backend servers in the following way: localhost:9000/host1 is proxied to
localhost:8081 (these redirects are specified in the configs/config.yml file). Each of the server apps reports its relative cpu and mem usage when /stats endpoint is hit.
The entire setup is deployed via docker utilities.
 ## Requirements 
docker, docker-compose
## Run
Clone this repo.
In the project directory:

`docker-compose build`

`docker-compose up -d`

To stop the reverse proxy:

`docker-compose down`
