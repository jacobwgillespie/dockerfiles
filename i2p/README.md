# jacobwgillespie/i2p

Provides the i2p router in a minimal Debian container.

### Usage

```bash
$ docker create \
  --name i2p \
  --net=host \
  -v </path/to/i2p/config>:/var/lib/i2p \
  jacobwgillespie/i2p
```

### Environment Variables

Variable | Default | Description
-------- | ------- | -----------
`TZ` | none | Set this to control the timezone in the container
`PUID` | none | Set the UID of the container user
`PGID` | none | Set the GID of the container user
`VNC_PASSWD` | none | If set, the VNC server will require a password
`VNC_PORT` | 5900 | Change the VNC server port
`VNC_RESOLUTION` | `1920x1080` | The display resolution for qBittorrent

### Ports

i2p exposes a variety of ports, however you may wish to use `host` networking to simplify port forwarding.  Be aware that i2p has been modified to listen on all devices rather than only `127.0.0.1` so that control ports are able to be exposed.

### Volumes

To persist the i2p configuration across container restarts, mount a volume to `/var/lib/i2p`.
