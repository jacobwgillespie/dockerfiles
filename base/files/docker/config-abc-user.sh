#!/bin/bash

# Set up permissions

PUID=${PUID:-911}
PGID=${PGID:-911}

groupmod -o -g "$PGID" abc
usermod -o -u "$PUID" abc

chown -R abc:abc /app
chown -R abc:abc /config
chown -R abc:abc /log

echo "abc user: UID ${PUID} GID ${PGID}"
