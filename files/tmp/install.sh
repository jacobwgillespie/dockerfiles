#!/bin/bash

export DEBIAN_FRONTEND="noninteractive"

# Create abc user
groupmod -g 1000 users
useradd -u 911 -U -d /config -s /bin/false abc
usermod -G users abc

# Make our folders
mkdir -p /app /config

# Disable SSH
rm -rf /etc/service/sshd /etc/service/cron /etc/my_init.d/00_regen_ssh_host_keys.sh

cd /files && find . -type f -exec cp -f --parents '{}' / \;

# Repositories

cat <<'EOT' > /etc/apt/sources.list
deb http://us.archive.ubuntu.com/ubuntu/ xenial main restricted universe multiverse
deb http://us.archive.ubuntu.com/ubuntu/ xenial-security main restricted universe multiverse
deb http://us.archive.ubuntu.com/ubuntu/ xenial-updates main restricted universe multiverse
deb http://us.archive.ubuntu.com/ubuntu/ xenial-proposed main restricted universe multiverse
deb http://us.archive.ubuntu.com/ubuntu/ xenial-backports main restricted universe multiverse
EOT

# Dependencies

apt-get update -qq

# Install window manager and x-server
apt-get install -qy --no-install-recommends \
                x11-xserver-utils \
                openbox \
                xfonts-base \
                xfonts-100dpi \
                xfonts-75dpi \
                libfuse2 \
                xbase-clients \
                xkb-data

# Installation

/bin/bash /tmp/qbittorrent-install.sh
/bin/bash /tmp/tigervnc-install.sh

# Services

chmod -R +x /etc/service/ /etc/my_init.d/ /opt/startapp.sh /opt/stopapp.sh
chmod 777 /tmp

# Cleanup

apt-get autoremove -y
apt-get clean -y
rm -rf /var/lib/apt/lists/* /var/cache/* /var/tmp/* /files/
