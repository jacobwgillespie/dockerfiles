# jacobwgillespie/qbittorrent-vnc

This provides the desktop qBittorrent app in a Docker container via VNC.

### Usage

```bash
$ docker create \
  --name qbittorrent \
  --net=host \
  -e PUID=<UID> -e PGID=<GID> \
  -e TZ=<timezone> \
  -v </path/to/qbittorrent/config>:/config \
  jacobwgillespie/qbittorrent-vnc

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

The VNC server is exposed on port 5900, however you may wish to use `host` networking to allow the wide range of ports that qBittorrent can be configured to use.

### Volumes

To persist the qBittorrent configuration across container restarts, mount a volume to `/config`.
