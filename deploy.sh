#!/bin/bash

username=$(wc -l < $1)
password=$(wc -l < $2)

echo "$password" | docker login registry.digitalocean.com --username $username --password-stdin

docker run -d -p 80:8080 --name helloworld do-kevinrobayna-docker-registry/helloworld:latest
