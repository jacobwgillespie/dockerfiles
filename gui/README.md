# jacobwgillespie/gui

A base image for containers that run GUI applications, exposing a VNC server for access.

### Usage

This image is not designed to be run directly, but is designed to be used as a base image for other GUI apps.

```
FROM jacobwgillespie/gui

ENV APP_NAME=MyApp
```

Add a file named `/opt/startapp.sh` with the commands necessary to start the GUI app - it will be started automatically.

### Environment Variables

Variable | Default | Description
-------- | ------- | -----------
`TZ` | none | Set this to control the timezone in the container
`PUID` | none | Set the UID of the container user
`PGID` | none | Set the GID of the container user
`VNC_PASSWD` | none | If set, the VNC server will require a password
`VNC_PORT` | 5900 | Change the VNC server port
`VNC_RESOLUTION` | `1920x1080` | The display resolution for the GUI app

### Ports

The VNC server is exposed on port 5900.
