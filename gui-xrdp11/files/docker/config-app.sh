#!/bin/bash

USERID=${USER_ID:-99}
GROUPID=${GROUP_ID:-100}
groupmod -g $GROUPID users
usermod -u $USERID nobody
usermod -g $GROUPID nobody
usermod -d /nobody nobody
usermod -a -G adm,sudo,fuse nobody
chown -R nobody:users /nobody/

chmod 1777 /tmp/

APP_NAME=${APP_NAME:-"GUI_APPLICATION"}
sed -i -e "s#GUI_APPLICATION#$APP_NAME#" /etc/xrdp/xrdp.ini
sed -i -e "s#GUI_APPLICATION#$APP_NAME#" /etc/guacamole/noauth-config.xml

if [[ -e /opt/startapp.sh ]]; then
  chown nobody:users /opt/startapp.sh
  chmod +x /opt/startapp.sh
fi

ldconfig
