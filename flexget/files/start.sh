#!/bin/bash

if [[ $(cat /etc/timezone) != $TZ ]] ; then
  ln -sf "/usr/share/zoneinfo/$TZ" /etc/localtime
  dpkg-reconfigure -f noninteractive tzdata
fi

PUID=${PUID:-911}
PGID=${PGID:-911}

groupmod -o -g "$PGID" abc
usermod -o -u "$PUID" abc

chown -R abc:abc /config

cd /config

rm -f /config/.config-lock

su - abc -s/bin/bash -c "flexget daemon start"
