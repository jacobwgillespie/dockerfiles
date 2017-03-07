# jacobwgillespie/i2p

Provides the i2p router in a minimal Debian container.

### Usage

```bash
$ docker create \
  --name i2p \
  --net=host \
  -e PUID=<UID> -e PGID=<GID> \
  -e TZ=<timezone> \
  -v </path/to/i2p/config>:/var/lib/i2p \
  jacobwgillespie/i2p
```

### Environment Variables

Variable | Default | Description
-------- | ------- | -----------
`TZ` | none | Set this to control the timezone in the container
`PUID` | none | Set the UID of the i2p user
`PGID` | none | Set the GID of the i2p user

### Ports

i2p exposes a variety of ports, however you may wish to use `host` networking to simplify port forwarding.  Be aware that i2p has been modified to listen on all devices rather than only `127.0.0.1` so that control ports are able to be exposed.

### Volumes

To persist the i2p configuration across container restarts, mount a volume to `/var/lib/i2p`.
