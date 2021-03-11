# CI CD Pipeline with GH Actions and Digital Ocean
Example of a simple pipeline that would build a docker image and deploy the container into a newly provision droplet

## How to Run locally

You should be able to build the docker image like this

```
docker build -t helloworld -f Dockerfile app
```

Then you should run the image itself with the following command

```
docker run -p 8080:8080 --rm -ti helloworld
```

Finally you can query the api with a good old curl

```
curl localhost:8080/api/hello
```

