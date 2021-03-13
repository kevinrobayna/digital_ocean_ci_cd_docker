#!/bin/bash

echo DO_REGISTRY_PASSWORD | docker login registry.digitalocean.com -u DO_REGISTRY_USERNAME --password-stdin
docker run -d -p 8080:8080 --name helloworld registry.digitalocean.com/do-kevinrobayna-docker-registry/helloworld:latest
curl localhost:8080/api/hello