# CI CD Pipeline with GH Actions and Digital Ocean
Example of a simple pipeline that would build a docker image and deploy the container into a newly provision droplet

## How to Run locally

```
docker build -t helloworld -f Dockerfile app
docker run -p 8080:8080 --rm -ti helloworld
```
