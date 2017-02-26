# jacobwgillespie/makemkv

This provides the desktop MakeMKV app in a Docker container via VNC.

### Usage

```bash
$ docker create \
  --name makemkv \
  -e PUID=<UID> -e PGID=<GID> \
  -e TZ=<timezone> \
  -v </path/to/makemkv/config>:/config \
  jacobwgillespie/makemkv

```

### Environment Variables

Variable | Default | Description
-------- | ------- | -----------
`TZ` | none | Set this to control the timezone in the container
`PUID` | none | Set the UID of the container user
`PGID` | none | Set the GID of the container user
`VNC_PASSWD` | none | If set, the VNC server will require a password
`VNC_PORT` | 5900 | Change the VNC server port
`VNC_RESOLUTION` | `1920x1080` | The display resolution for MakeMKV

### Ports

The VNC server is exposed on port `5900`.

### Volumes

To persist the qBittorrent configuration across container restarts, mount a volume to `/config`.
