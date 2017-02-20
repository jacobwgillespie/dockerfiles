# jacobwgillespie/flexget

Flexget in a simple Docker container.

### Usage

```
$ docker run -d \
  --net=host \
  --name flexget \
  -v <path/to/config>:/config \
  jacobwgillespie/flexget
```

### Environment Variables

Variable | Default | Description
-------- | ------- | -----------
`TZ` | none | Set this to control the timezone in the container
`PUID` | none | Set the UID of the container user
`PGID` | none | Set the GID of the container user
