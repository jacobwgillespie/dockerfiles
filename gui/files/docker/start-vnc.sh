#!/bin/bash

# Load default values if empty
VNC_RESOLUTION="${VNC_RESOLUTION:-1920x1080}"
VNC_CREDENTIALS=/tmp/vnc_passwd
VNC_PORT="${VNC_PORT:-5900}"

if [[ -n $VNC_PASSWD ]]; then
  VNC_SECURITY="SecurityTypes TLSVnc,VncAuth -PasswordFile ${VNC_CREDENTIALS}"
else
  VNC_SECURITY="SecurityTypes None"
fi

exec /usr/bin/Xvnc :1 \
  -geometry ${VNC_RESOLUTION} \
  -depth 24 \
  -rfbwait 30000 \
  -${VNC_SECURITY} \
  -rfbport ${VNC_PORT} \
  -bs \
  -ac \
  -pn \
  -fp /usr/share/fonts/X11/misc/,/usr/share/fonts/X11/75dpi/,/usr/share/fonts/X11/100dpi/ \
  -dpi 100 \
  -desktop "${APP_NAME:-App}"
