FROM jacobwgillespie/base

RUN \
  apt-get update -qq && \

  # Install window manager and x-server
  apt-get install -qy --no-install-recommends \
    libfuse2 \
    openbox \
    python-xdg \
    x11-xserver-utils \
    xbase-clients \
    xfonts-100dpi \
    xfonts-75dpi \
    xfonts-base \
    xkb-data && \

  # Install TigerVNC
  curl -L https://dl.bintray.com/tigervnc/stable/tigervnc-1.7.1.x86_64.tar.gz | tar -xz --strip=1 -C / && \

  # Install vncpasswd.py
  mkdir -p /opt/vncpasswd && \
  curl -L https://github.com/trinitronx/vncpasswd.py/archive/master.tar.gz | tar -zx --strip=1 -C /opt/vncpasswd && \

  # Cleanup
  apt-get autoremove -y && \
  apt-get clean -y && \
  rm -rf /tmp/* /var/tmp/* /var/lib/apt/lists/* /var/cache/*

ADD /files/ /

EXPOSE 5900 6080
