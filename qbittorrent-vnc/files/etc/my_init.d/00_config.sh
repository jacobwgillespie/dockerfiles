#!/bin/bash

# Configure the timezone

if [[ $(cat /etc/timezone) != $TZ ]] ; then
  ln -sf "/usr/share/zoneinfo/$TZ" /etc/localtime
  dpkg-reconfigure -f noninteractive tzdata
fi

# Set up permissions

PUID=${PUID:-911}
PGID=${PGID:-911}

groupmod -o -g "$PGID" abc
usermod -o -u "$PUID" abc

chown -R abc:abc /app
chown -R abc:abc /config

echo "abc user: UID ${PUID} GID ${PGID}"
