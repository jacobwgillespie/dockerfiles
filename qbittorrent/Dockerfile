FROM jacobwgillespie/gui

RUN \
  apt-get update -qq && \

  # Install qbittorrent
  apt-get install -qy software-properties-common python-software-properties && \
  add-apt-repository -y ppa:qbittorrent-team/qbittorrent-stable && \
  apt-get update && \
  apt-get install -qy qbittorrent && \

  # Cleanup
  apt-get autoremove -y && \
  apt-get clean -y && \
  rm -rf /tmp/* /var/tmp/* /var/lib/apt/lists/* /var/cache/*

ADD /files/ /

ENV APP_NAME=qBittorrent
