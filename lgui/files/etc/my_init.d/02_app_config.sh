#!/bin/bash

APPNAME=${APP_NAME:-"GUI_APPLICATION"}

sed -i -e "s#GUI_APPLICATION#$APPNAME#" /etc/xrdp/xrdp.ini
sed -i -e "s#GUI_APPLICATION#$APPNAME#" /etc/guacamole/noauth-config.xml

if [[ -e /startapp.sh ]]; then
    chown nobody:users /startapp.sh
    chmod +x /startapp.sh
fi
