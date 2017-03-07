#!/bin/bash
HOSTNAME="${HOSTNAME:-_}"
PUID=${PUID:-911}
PGID=${PGID:-911}

groupmod -o -g "$PGID" danbooru
usermod -o -u "$PUID" danbooru

sed -i -e "s/__hostname__/$HOSTNAME/g" /etc/nginx/conf.d/danbooru.conf

# Generate secret token and secret key
mkdir /config/.danbooru/
openssl rand -hex 32 > /config/.danbooru/secret_token
openssl rand -hex 32 > /config/.danbooru/session_secret_key
chmod 600 /config/.danbooru/*

chown -R danbooru:danbooru /app
chown -R danbooru:danbooru /config
chown -R danbooru:danbooru /log

echo "danbooru user: UID ${PUID} GID ${PGID}"

exec /usr/bin/supervisord
