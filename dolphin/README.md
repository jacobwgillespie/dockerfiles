# jacobwgillespie/dolphin

This provides the desktop Dolphin app in a Docker container via VNC.

### Usage

```bash
$ docker create \
  --name dolphin \
  -e PUID=<UID> -e PGID=<GID> \
  -e TZ=<timezone> \
  -v </path/to/files>:/data \
  jacobwgillespie/dolphin
```

### Environment Variables

Variable | Default | Description
-------- | ------- | -----------
`TZ` | none | Set this to control the timezone in the container
`PUID` | none | Set the UID of the container user
`PGID` | none | Set the GID of the container user
`VNC_PASSWD` | none | If set, the VNC server will require a password
`VNC_PORT` | 5900 | Change the VNC server port
`VNC_RESOLUTION` | `1920x1080` | The display resolution for Dolphin

### Ports

The VNC server is exposed on port `5900`.

### Volumes

To access files from the host inside the container, mount them in a volume.
