# playlist/redis - redis container

[![image](https://d207aa93qlcgug.cloudfront.net/img/icons/framed-icon-checked-repository.svg)](https://registry.hub.docker.com/u/playlist/redis/)

[**Trusted Build**](https://registry.hub.docker.com/u/playlist/redis/)

This Docker image is based on the official [ubuntu:14.04](https://registry.hub.docker.com/_/ubuntu/) base image.

## Building

```shell
$ make build
```

## Running

### Standalone

```shell
$ docker run -d -p 6379:6379 playlist/redis
```

### Cluster
It is required that you set the publish host IP in `HOST`, and if the etcd host is different than `HOST`, set it in `ETCD_HOST`.

```shell
$ docker run --name database --rm -p 6379:6379 -e PUBLISH=6379 -e HOST=$COREOS_PRIVATE_IPV4 -e CLUSTER=redis -v /path/to/local:/data playlist/redis:latest"
```
