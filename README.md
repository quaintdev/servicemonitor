# Service Monitor

An API monitoring service

## Installation using Docker
1. Pull docker image of urlshortener on your local system using: 

    `docker pull quaintdev/servicemonitor`

2. Run urlshortener using

    `docker run servicemonitor`


The parameters for service are configurable within main.go. Ideally these should be stored in a JSON based configuration file but that is not currently implemented.
