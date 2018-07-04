FROM jacobwgillespie/gui

RUN \
  apt-get update -qq && \
  # Install build dependencies
  apt-get install -qy \
    libboost-dev libboost-system-dev build-essential \
    qtbase5-dev qttools5-dev-tools \
    libboost-system-dev libboost-chrono-dev libboost-random-dev libssl-dev libgeoip-dev \
    git pkg-config automake libtool \
    python geoip-database tzdata && \
  # Prepare for compilation
  mkdir -p /root/sources && \
  # Compile libtorrent
  git clone https://github.com/arvidn/libtorrent.git /root/sources/libtorrent && \
  cd /root/sources/libtorrent && \
  git checkout RC_1_0 && \
  ./autotool.sh && \
  ./configure --disable-debug --enable-encryption --prefix=/usr --with-libgeoip=system && \
  make clean && \
  make && \
  # Install libtorrent
  make install && \
  # Compile qbittorrent
  git clone https://github.com/qbittorrent/qBittorrent /root/sources/qBittorrent && \
  cd /root/sources/qBittorrent && \
  git checkout release-3.3.12 && \
  ./configure --prefix=/usr && \
  make && \
  # Install qbittorrent
  make install && \
  # Cleanup
  apt-get remove -y build-essential && \
  apt-get autoremove -y && \
  apt-get clean -y && \
  rm -rf /tmp/* /var/tmp/* /var/lib/apt/lists/* /var/cache/* /root/sources/*

ADD /files/ /

ENV APP_NAME=qBittorrent
