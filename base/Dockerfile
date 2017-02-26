FROM ubuntu:16.04

ENV \
  DEBIAN_FRONTEND="noninteractive" \
  HOME="/root" \
  LANG="en_US.UTF-8" \
  LANGUAGE="en_US.UTF-8" \
  LC_ALL="C.UTF-8" \
  TERM="xterm"

RUN \
  # Install dependencies
  apt-get update -qq && \
  apt-get install -qy --no-install-recommends \
    unzip python3-pip python3-setuptools curl wget less git && \

  # Install process manager
  pip3 install wheel && \
  pip3 install chaperone && \
  mkdir -p /etc/chaperone.d && \

  # Add default directories
  mkdir -p /app /config /log && \

  # Add abc user
  groupmod -g 1000 users && \
  useradd -u 911 -U -d /config -s /bin/bash abc && \
  usermod -G users abc && \

  # Cleanup
  apt-get autoremove -y && \
  apt-get clean -y && \
  rm -rf /tmp/* /var/tmp/* /var/lib/apt/lists/* /var/cache/* /root/.cache/*

ADD /files/ /

ENTRYPOINT ["/usr/local/bin/chaperone"]
